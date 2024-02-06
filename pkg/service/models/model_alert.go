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

package models

import (
	"strconv"
	"time"

	"gorm.io/datatypes"
	"kubegems.io/kubegems/pkg/utils"
	"kubegems.io/kubegems/pkg/utils/msgbus"
)

type AlertInfo struct {
	Fingerprint     string `gorm:"type:varchar(50);primaryKey"` // 指纹作为主键
	Name            string `gorm:"type:varchar(50);"`
	Namespace       string `gorm:"type:varchar(50);"`
	ClusterName     string `gorm:"type:varchar(50);"`
	TenantName      string `gorm:"type:varchar(50);index"`
	ProjectName     string `gorm:"type:varchar(50);index"`
	EnvironmentName string `gorm:"type:varchar(50);index"`
	Labels          datatypes.JSON
	LabelMap        map[string]string `gorm:"-" json:"-"`

	SilenceStartsAt  *time.Time
	SilenceUpdatedAt *time.Time
	SilenceEndsAt    *time.Time
	SilenceCreator   string `gorm:"type:varchar(50);"`
	Summary          string `gorm:"-"` // 黑名单概要
}

type AlertMessage struct {
	ID uint

	// 级联删除
	InfoFingerprint string     `gorm:"type:varchar(50);column:fingerprint;"`
	AlertInfo       *AlertInfo `gorm:"foreignKey:InfoFingerprint;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Value     string
	Message   string
	StartsAt  *time.Time `gorm:"index"` // 告警开始时间
	EndsAt    *time.Time // 告警结束时间
	CreatedAt *time.Time `gorm:"index"` // 本次告警产生时间
	Status    string     // firing or resolved
}

func (a *AlertMessage) ToNormalMessage() Message {
	return Message{
		ID:          a.ID,
		MessageType: string(msgbus.Alert),
		Title:       a.Message,
		CreatedAt:   *a.CreatedAt,
	}
}

func (a *AlertMessage) ColumnSlice() []string {
	return []string{"id", "fingerprint", "value", "message", "starts_at", "ends_at", "created_at", "status", "labels"}
}

func (a *AlertMessage) ValueSlice() []string {
	return []string{
		strconv.Itoa(int(a.ID)),
		a.InfoFingerprint,
		a.Value,
		a.Message,
		utils.FormatMysqlDumpTime(a.StartsAt),
		utils.FormatMysqlDumpTime(a.EndsAt),
		utils.FormatMysqlDumpTime(a.CreatedAt),
		a.Status,
		func() string {
			if a.AlertInfo != nil {
				return a.AlertInfo.Labels.String()
			}
			return ""
		}(),
	}
}
