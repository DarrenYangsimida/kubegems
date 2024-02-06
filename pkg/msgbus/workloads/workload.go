// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workloads

import (
	"context"
	"io/ioutil"
	"sync"
	"time"

	"kubegems.io/kubegems/pkg/log"
	"kubegems.io/kubegems/pkg/msgbus/switcher"
	"kubegems.io/kubegems/pkg/utils/agents"
	"kubegems.io/kubegems/pkg/utils/msgbus"
)

type AgentMessageCollector struct {
	context       context.Context
	clientsSet    *agents.ClientSet
	ms            *switcher.MessageSwitcher
	clusterStatus sync.Map
	messageCh     chan *msgbus.NotifyMessage
}

func RunWorkloadCollector(ctx context.Context, cs *agents.ClientSet, ms *switcher.MessageSwitcher) error {
	ctrl := NewCtrl(ctx, cs, ms)
	return ctrl.Run()
}

func NewCtrl(ctx context.Context, agent *agents.ClientSet, ms *switcher.MessageSwitcher) *AgentMessageCollector {
	return &AgentMessageCollector{
		context:       ctx,
		clientsSet:    agent,
		ms:            ms,
		clusterStatus: sync.Map{},
		messageCh:     make(chan *msgbus.NotifyMessage, 1000),
	}
}

func (c *AgentMessageCollector) Refresh() {
	clusters := c.clientsSet.Clusters()
	clusterMap := map[string]interface{}{}
	for _, cluster := range clusters {
		clusterMap[cluster] = cluster
		if _, ok := c.clusterStatus.Load(cluster); !ok {
			stopch := make(chan struct{})
			go c.MessageChan(c.context, cluster, stopch)
			c.clusterStatus.Store(cluster, stopch)
		}
	}
	var toStopCluster []string
	c.clusterStatus.Range(func(k, v interface{}) bool {
		if _, exist := clusterMap[k.(string)]; !exist {
			toStopCluster = append(toStopCluster, k.(string))
		}
		return true
	})
	for _, clusterName := range toStopCluster {
		ch, exist := c.clusterStatus.Load(clusterName)
		if !exist {
			continue
		}
		stopch := ch.(chan struct{})
		stopch <- struct{}{}
		c.clusterStatus.Delete(clusterName)
	}
}

func (c *AgentMessageCollector) Run() error {
	clusters := c.clientsSet.Clusters()
	msgswitch := c.ms
	for _, cluster := range clusters {
		stopch := make(chan struct{})
		go c.MessageChan(c.context, cluster, stopch)
		c.clusterStatus.Store(cluster, stopch)
	}

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-c.context.Done():
			ticker.Stop()
			return c.context.Err()
		case msg := <-c.messageCh:
			msgswitch.DispatchMessage(msg)
		case <-ticker.C:
			c.Refresh()
		}
	}
}

func (c *AgentMessageCollector) MessageChan(ctx context.Context, clustername string, stopch chan struct{}) {
	log := log.FromContextOrDiscard(ctx).WithValues("cluster", clustername)
	uri := "notify"
	for {
		func() error {
			clusterProxy, err := c.clientsSet.ClientOf(ctx, clustername)
			if err != nil {
				log.Error(err, "get client")
				return err
			}
			conn, resp, err := clusterProxy.Websocket().DialPath(ctx, uri, nil)
			if err != nil {
				content := ""
				if resp != nil {
					t, _ := ioutil.ReadAll(resp.Body)
					content = string(t)
				}
				log.Error(err, "dial websocket", "content", content)
				return err
			}
			defer resp.Body.Close()
			defer conn.Close()

			for {
				tmp := msgbus.NotifyMessage{}
				if err := conn.ReadJSON(&tmp); err != nil {
					log.Error(err, "decode json")
					return err
				}
				switch tmp.MessageType {
				case msgbus.Changed:
					tmp.InvolvedObject.Cluster = clustername
					c.messageCh <- &tmp
				case msgbus.Alert:
					c.messageCh <- &tmp
				}
			}
		}()
		// 删除逻辑，等到出错后再处理
		select {
		case <-stopch:
			return
		default:
			time.Sleep(time.Second * 5)
		}
	}
}
