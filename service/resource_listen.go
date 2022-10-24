/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package service

import (
	"context"

	api "github.com/polarismesh/polaris/common/api/v1"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
)

// ResourceHook The listener is placed before and after the resource operation, only normal flow
type ResourceHook interface {

	// Before
	//  @param ctx
	//  @param resourceType
	Before(ctx context.Context, resourceType model.Resource)

	// After
	//  @param ctx
	//  @param resourceType
	//  @param res
	After(ctx context.Context, resourceType model.Resource, res *ResourceEvent) error
}

// ResourceEvent 资源事件
type ResourceEvent struct {
	ReqService *api.Service
	Service    *model.Service
	IsRemove   bool
}

// Before this function is called before the resource operation
func (svr *serverAuthAbility) Before(ctx context.Context, resourceType model.Resource) {
	// do nothing
}

// After this function is called after the resource operation
func (svr *serverAuthAbility) After(ctx context.Context, resourceType model.Resource, res *ResourceEvent) error {
	switch resourceType {
	case model.RService:
		return svr.onServiceResource(ctx, res)
	default:
		return nil
	}
}

// onServiceResource 服务资源的处理，只处理服务，namespace 只由 namespace 相关的进行处理，
func (svr *serverAuthAbility) onServiceResource(ctx context.Context, res *ResourceEvent) error {
	authCtx := ctx.Value(utils.ContextAuthContextKey).(*model.AcquireContext)
	ownerId := utils.ParseOwnerID(ctx)

	authCtx.SetAttachment(model.ResourceAttachmentKey, map[api.ResourceType][]model.ResourceEntry{
		api.ResourceType_Services: {
			{
				ID:    res.Service.ID,
				Owner: ownerId,
			},
		},
	})

	users := utils.ConvertStringValuesToSlice(res.ReqService.UserIds)
	removeUses := utils.ConvertStringValuesToSlice(res.ReqService.RemoveUserIds)

	groups := utils.ConvertStringValuesToSlice(res.ReqService.GroupIds)
	removeGroups := utils.ConvertStringValuesToSlice(res.ReqService.RemoveGroupIds)

	authCtx.SetAttachment(model.LinkUsersKey, utils.StringSliceDeDuplication(users))
	authCtx.SetAttachment(model.RemoveLinkUsersKey, utils.StringSliceDeDuplication(removeUses))

	authCtx.SetAttachment(model.LinkGroupsKey, utils.StringSliceDeDuplication(groups))
	authCtx.SetAttachment(model.RemoveLinkGroupsKey, utils.StringSliceDeDuplication(removeGroups))

	return svr.authSvr.AfterResourceOperation(authCtx)
}
