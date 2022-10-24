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

	"go.uber.org/zap"

	api "github.com/polarismesh/polaris/common/api/v1"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
)

var (
	clientFilterAttributes = map[string]struct{}{
		"type":    {},
		"host":    {},
		"limit":   {},
		"offset":  {},
		"version": {},
	}
)

func (s *Server) checkAndStoreClient(ctx context.Context, req *api.Client) *api.Response {
	clientId := req.GetId().GetValue()
	var needStore bool
	client := s.caches.Client().GetClient(clientId)
	var resp *api.Response
	if nil == client {
		needStore = true
	} else {
		needStore = !clientEquals(client.Proto(), req)
	}
	if needStore {
		client, resp = s.createClient(ctx, req)
	}

	if resp != nil {
		if resp.GetCode().GetValue() != api.ExistedResource {
			return resp
		}
	}

	return s.HealthServer().ReportByClient(context.Background(), req)
}

func (s *Server) createClient(ctx context.Context, req *api.Client) (*model.Client, *api.Response) {

	if namingServer.bc == nil || !namingServer.bc.ClientRegisterOpen() {
		return nil, nil
	}
	return s.asyncCreateClient(ctx, req) // 批量异步
}

// 异步新建客户端
// 底层函数会合并create请求，增加并发创建的吞吐
// req 原始请求
// ins 包含了req数据与instanceID，serviceToken
func (s *Server) asyncCreateClient(ctx context.Context, req *api.Client) (*model.Client, *api.Response) {
	rid := utils.ParseRequestID(ctx)
	pid := utils.ParsePlatformID(ctx)
	future := s.bc.AsyncRegisterClient(req)
	if err := future.Wait(); err != nil {
		log.Error("[Server][ReportClient] async create client", zap.Error(err), utils.ZapRequestID(rid),
			utils.ZapPlatformID(pid))
		if future.Code() == api.ExistedResource {
			req.Id = utils.NewStringValue(req.GetId().GetValue())
		}
		return nil, api.NewClientResponse(future.Code(), req)
	}

	return future.Client(), nil
}

// GetReportClients create one instance
func (s *Server) GetReportClients(ctx context.Context, query map[string]string) *api.BatchQueryResponse {
	searchFilters := make(map[string]string)
	var (
		offset, limit uint32
		err           error
	)

	for key, value := range query {
		if _, ok := clientFilterAttributes[key]; !ok {
			log.Errorf("[Server][Client] attribute(%s) it not allowed", key)
			return api.NewBatchQueryResponseWithMsg(api.InvalidParameter, key+" is not allowed")
		}
		searchFilters[key] = value
	}

	var (
		total   uint32
		clients []*model.Client
	)

	offset, limit, err = utils.ParseOffsetAndLimit(searchFilters)
	if err != nil {
		return api.NewBatchQueryResponse(api.InvalidParameter)
	}

	total, services, err := s.caches.Client().GetClientsByFilter(searchFilters, offset, limit)
	if err != nil {
		log.Errorf("[Server][Client][Query] req(%+v) store err: %s", query, err.Error())
		return api.NewBatchQueryResponse(api.StoreLayerException)
	}

	resp := api.NewBatchQueryResponse(api.ExecuteSuccess)
	resp.Amount = utils.NewUInt32Value(total)
	resp.Size = utils.NewUInt32Value(uint32(len(services)))
	resp.Clients = enhancedClients2Api(clients, client2Api)
	return resp
}

type Client2Api func(client *model.Client) *api.Client

// client 数组转为[]*api.Client
func enhancedClients2Api(clients []*model.Client, handler Client2Api) []*api.Client {
	out := make([]*api.Client, 0, len(clients))
	for _, entry := range clients {
		outUser := handler(entry)
		out = append(out, outUser)
	}
	return out
}

// model.Client 转为 api.Client
func client2Api(client *model.Client) *api.Client {
	if client == nil {
		return nil
	}
	out := client.Proto()
	return out
}

func clientEquals(client1 *api.Client, client2 *api.Client) bool {
	if client1.GetId().GetValue() != client2.GetId().GetValue() {
		return false
	}
	if client1.GetHost().GetValue() != client2.GetHost().GetValue() {
		return false
	}
	if client1.GetVersion().GetValue() != client2.GetVersion().GetValue() {
		return false
	}
	if client1.GetType() != client2.GetType() {
		return false
	}
	if client1.GetLocation().GetRegion().GetValue() != client2.GetLocation().GetRegion().GetValue() {
		return false
	}
	if client1.GetLocation().GetZone().GetValue() != client2.GetLocation().GetZone().GetValue() {
		return false
	}
	if client1.GetLocation().GetCampus().GetValue() != client2.GetLocation().GetCampus().GetValue() {
		return false
	}
	if len(client1.Stat) != len(client2.Stat) {
		return false
	}
	for i := 0; i < len(client1.Stat); i++ {
		if client1.Stat[i].GetTarget().GetValue() != client2.Stat[i].GetTarget().GetValue() {
			return false
		}
		if client1.Stat[i].GetPort().GetValue() != client2.Stat[i].GetPort().GetValue() {
			return false
		}
		if client1.Stat[i].GetPath().GetValue() != client2.Stat[i].GetPath().GetValue() {
			return false
		}
		if client1.Stat[i].GetProtocol().GetValue() != client2.Stat[i].GetProtocol().GetValue() {
			return false
		}
	}
	return true
}
