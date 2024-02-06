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

package sels

import (
	"github.com/gin-gonic/gin"
	"kubegems.io/kubegems/pkg/service/handlers"
	"kubegems.io/kubegems/pkg/service/models"
)

/*
	sels 模块处理全局的selsecor数据，只处理ID和名字
*/

//	@Tags			Sels
//	@Summary		user sels
//	@Description	user sels
//	@Accept			json
//	@Produce		json
//	@Param			tenant_id	query		string															false	"tenant_id"
//	@Param			all			query		string															false	"all"
//	@Success		200			{object}	handlers.ResponseStruct{Data=handlers.PageData{List=[]UserSel}}	"data"
//	@Router			/v1/sels/users [get]
//	@Security		JWT
func (h *SelsHandler) UserSels(c *gin.Context) {
	ret := []UserSel{}
	tenant_id := c.Query("tenant_id")
	all := c.Query("all")
	ctx := c.Request.Context()
	if tenant_id == "" {
		if all == "true" {
			h.GetDB().WithContext(ctx).Model(&models.User{}).Find(&ret)
		} else {
			h.GetDB().WithContext(ctx).Model(&models.User{}).Find(&ret, "is_active = true")
		}
	} else {
		h.GetDB().WithContext(ctx).Model(&models.User{}).Joins(
			"join tenant_user_rels on tenant_user_rels.user_id = users.id",
		).Find(
			&ret,
			"tenant_user_rels.tenant_id = ? and is_active = true",
			tenant_id,
		)
	}
	handlers.OK(c, handlers.Page(int64(len(ret)), ret, 1, 10000))
}

//	@Tags			Sels
//	@Summary		tenant sels
//	@Description	tenant sels
//	@Accept			json
//	@Produce		json
//	@Param			all	query		string																false	"是否全部，默认(false)，只有激活的"
//	@Success		200	{object}	handlers.ResponseStruct{Data=handlers.PageData{List=[]TenantSel}}	"data"
//	@Router			/v1/sels/tenants [get]
//	@Security		JWT
func (h *SelsHandler) TenantSels(c *gin.Context) {
	ret := []TenantSel{}
	all := c.Query("all")
	ctx := c.Request.Context()
	if all == "true" {
		h.GetDB().WithContext(ctx).Model(&models.Tenant{}).Find(&ret)
	} else {
		h.GetDB().WithContext(ctx).Model(&models.Tenant{}).Find(&ret, "is_active = true")
	}
	quotas := []models.TenantResourceQuota{}
	h.GetDB().WithContext(ctx).Model(&models.TenantResourceQuota{}).Preload("Cluster").Find(&quotas)
	tmap := map[uint][]string{}
	for _, quota := range quotas {
		if quota.Cluster != nil {
			tmap[quota.TenantID] = append(tmap[quota.TenantID], quota.Cluster.ClusterName)
		}
	}
	for idx := range ret {
		v, exist := tmap[ret[idx].ID]
		if exist {
			ret[idx].Clusters = v
		} else {
			ret[idx].Clusters = []string{}
		}
	}
	handlers.OK(c, handlers.Page(int64(len(ret)), ret, 1, 10000))
}

//	@Tags			Sels
//	@Summary		project sels
//	@Description	project sels
//	@Accept			json
//	@Produce		json
//	@Param			tenant_id	query		string																false	"tenant_id"
//	@Success		200			{object}	handlers.ResponseStruct{Data=handlers.PageData{List=[]ProjectSel}}	"data"
//	@Router			/v1/sels/projects [get]
//	@Security		JWT
func (h *SelsHandler) ProjectSels(c *gin.Context) {
	ret := []ProjectSel{}
	tenant_id := c.Query("tenant_id")
	ctx := c.Request.Context()
	if tenant_id == "" {
		h.GetDB().WithContext(ctx).Model(&models.Project{}).Find(&ret)
	} else {
		h.GetDB().WithContext(ctx).Model(&models.Project{}).Find(&ret, "tenant_id = ?", tenant_id)
	}
	handlers.OK(c, handlers.Page(int64(len(ret)), ret, 1, 10000))
}

//	@Tags			Sels
//	@Summary		environment sels
//	@Description	environment sels
//	@Accept			json
//	@Produce		json
//	@Param			tenant_id	query		string																	false	"tenant_id"
//	@Success		200			{object}	handlers.ResponseStruct{Data=handlers.PageData{List=[]EnvironmentSel}}	"data"
//	@Router			/v1/sels/environments [get]
//	@Security		JWT
func (h *SelsHandler) EnvironmentSels(c *gin.Context) {
	ret := []EnvironmentSel{}
	tenant_id := c.Query("tenant_id")
	ctx := c.Request.Context()
	if tenant_id == "" {
		h.GetDB().WithContext(ctx).Model(&models.Environment{}).Find(&ret)
	} else {
		h.GetDB().WithContext(ctx).Model(&models.Environment{}).Find(&ret, "tenant_id = ?", tenant_id)
	}
	handlers.OK(c, handlers.Page(int64(len(ret)), ret, 1, 10000))
}

//	@Tags			Sels
//	@Summary		application sels
//	@Description	application sels
//	@Accept			json
//	@Produce		json
//	@Param			project_id	query		string																	false	"project_id"
//	@Success		200			{object}	handlers.ResponseStruct{Data=handlers.PageData{List=[]ApplicationSel}}	"data"
//	@Router			/v1/sels/applications [get]
//	@Security		JWT
func (h *SelsHandler) ApplicationSels(c *gin.Context) {
	ret := []ApplicationSel{}
	project_id := c.Query("project_id")
	ctx := c.Request.Context()
	if project_id == "" {
		h.GetDB().WithContext(ctx).Model(&models.Application{}).Find(&ret)
	} else {
		h.GetDB().WithContext(ctx).Model(&models.Application{}).Find(&ret, "project_id = ?", project_id)
	}
	handlers.OK(c, handlers.Page(int64(len(ret)), ret, 1, 10000))
}

type UserSel struct {
	ID       uint
	Username string
	Email    string
	IsActive bool
}

type TenantSel struct {
	ID         uint
	TenantName string
	IsActive   bool
	Clusters   []string `gorm:"-"`
}

type ProjectSel struct {
	ID          uint
	ProjectName string
}

type EnvironmentSel struct {
	ID              uint
	EnvironmentName string
}

type ApplicationSel struct {
	ID              uint
	ApplicationName string
}
