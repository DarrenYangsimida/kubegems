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

package task

import (
	"kubegems.io/kubegems/pkg/apps/application"
	"kubegems.io/kubegems/pkg/utils/agents"
	"kubegems.io/kubegems/pkg/utils/argo"
	"kubegems.io/kubegems/pkg/utils/database"
	"kubegems.io/kubegems/pkg/utils/git"
	"kubegems.io/kubegems/pkg/utils/redis"
)

type ApplicationTasker struct {
	*application.ApplicationProcessor
}

func MustNewApplicationTasker(db *database.Database, gitp git.Provider, argo *argo.Client, redis *redis.Client, agents *agents.ClientSet) *ApplicationTasker {
	app := application.NewApplicationProcessor(db, gitp, argo, redis, agents)
	return &ApplicationTasker{ApplicationProcessor: app}
}

func (t *ApplicationTasker) ProvideFuntions() map[string]interface{} {
	return t.ApplicationProcessor.ProvideFuntions()
}
