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

package api

import (
	"net/http"
	"sort"

	"kubegems.io/kubegems/pkg/installer/pluginmanager"
	"kubegems.io/library/rest/request"
	"kubegems.io/library/rest/response"
)

type PluginStatus struct {
	Name               string         `json:"name"`
	Namespace          string         `json:"namespace"`
	Description        string         `json:"description"`
	InstalledVersion   string         `json:"installedVersion"`
	UpgradeableVersion string         `json:"upgradeableVersion"`
	AvailableVersions  []string       `json:"availableVersions"`
	Required           bool           `json:"required"`
	Enabled            bool           `json:"enabled"`
	Healthy            bool           `json:"healthy"`
	Message            string         `json:"message"`
	Values             map[string]any `json:"values"` // current installed version values
	maincate           string
	cate               string
}

func (o *PluginsAPI) ListPlugins(resp http.ResponseWriter, req *http.Request) {
	if request.Query(req, "check-update", false) {
		upgradeable, err := o.PM.CheckUpdate(req.Context())
		if err != nil {
			response.Error(resp, err)
			return
		}
		response.OK(resp, upgradeable)
	} else {
		plugins, err := o.PM.ListPlugins(req.Context())
		if err != nil {
			response.Error(resp, err)
			return
		}
		response.OK(resp, plugins)
	}
}

func SortPluginStatusByName(list []PluginStatus) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
}

func ToViewPlugin(plugin pluginmanager.Plugin) PluginStatus {
	ps := PluginStatus{
		Name:        plugin.Name,
		Namespace:   plugin.Namespace,
		Description: plugin.Description,
		Required:    plugin.Required,
		maincate:    plugin.MainCategory,
		cate:        plugin.Category,
	}
	if installed := plugin.Installed; installed != nil {
		ps.Enabled = plugin.Installed.Enabled
		ps.InstalledVersion = installed.Version
		ps.Healthy = installed.Healthy
		ps.Namespace = installed.InstallNamespace
		ps.Message = installed.Message
		ps.Values = installed.Values.Object
	}
	if upgradble := plugin.Upgradeable; upgradble != nil {
		ps.UpgradeableVersion = upgradble.Version
	}

	availableVersion := []string{}
	for _, item := range plugin.Available {
		availableVersion = append(availableVersion, item.Version)
	}
	ps.AvailableVersions = availableVersion
	return ps
}

func CategoriedPlugins(plugins map[string]pluginmanager.Plugin) map[string]map[string][]PluginStatus {
	pluginstatus := []PluginStatus{}

	for _, plugin := range plugins {
		pluginstatus = append(pluginstatus, ToViewPlugin(plugin))
	}

	mainCategoryFunc := func(t PluginStatus) string {
		return t.maincate
	}
	categoryfunc := func(t PluginStatus) string {
		return t.cate
	}

	categoryPlugins := map[string]map[string][]PluginStatus{}
	for maincategory, list := range withCategory(pluginstatus, mainCategoryFunc) {
		categorized := withCategory(list, categoryfunc)
		// sort
		for _, list := range categorized {
			SortPluginStatusByName(list)
		}
		categoryPlugins[maincategory] = categorized
	}
	return categoryPlugins
}

func withCategory[T any](list []T, getCate func(T) string) map[string][]T {
	ret := map[string][]T{}
	for _, v := range list {
		cate := getCate(v)
		if cate == "" {
			cate = "others"
		}
		if _, ok := ret[cate]; !ok {
			ret[cate] = []T{}
		}
		ret[cate] = append(ret[cate], v)
	}
	return ret
}

func (o *PluginsAPI) GetPlugin(resp http.ResponseWriter, req *http.Request) {
	name := request.Path(req, "name", "")
	version := request.Query(req, "version", "")
	withSchema := request.Query(req, "schema", false)
	healthCheck := request.Query(req, "check", false)

	pv, err := o.PM.GetPluginVersion(req.Context(), name, version, withSchema, healthCheck)
	if err != nil {
		response.Error(resp, err)
		return
	}
	response.OK(resp, pv)
}

func (o *PluginsAPI) EnablePlugin(resp http.ResponseWriter, req *http.Request) {
	name := request.Path(req, "name", "")
	version := request.Query(req, "version", "")

	pv := &pluginmanager.PluginVersion{}
	if err := request.Body(req, pv); err != nil {
		response.Error(resp, err)
		return
	}
	if err := o.PM.Install(req.Context(), name, version, pv.Values.Object); err != nil {
		response.Error(resp, err)
		return
	}
	response.OK(resp, pv)
}

func (o *PluginsAPI) RemovePlugin(resp http.ResponseWriter, req *http.Request) {
	name := request.Path(req, "name", "")
	if err := o.PM.UnInstall(req.Context(), name); err != nil {
		response.Error(resp, err)
		return
	}
	response.OK(resp, "ok")
}
