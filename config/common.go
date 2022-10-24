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

package config

import (
	"context"
	"errors"

	api "github.com/polarismesh/polaris/common/api/v1"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/store"
)

const (
	ContextTxKey = utils.StringContext("Config-Tx")
)

// StartTxAndSetToContext 开启一个事务，并放入到上下文里
func (s *Server) StartTxAndSetToContext(ctx context.Context) (store.Tx, context.Context, error) {
	tx, err := s.storage.StartTx()
	return tx, context.WithValue(ctx, ContextTxKey, tx), err
}

// getTx 从上下文里获取事务对象
func (s *Server) getTx(ctx context.Context) store.Tx {
	tx := ctx.Value(ContextTxKey)
	if tx == nil {
		return nil
	}
	return tx.(store.Tx)
}

func (s *Server) checkNamespaceExisted(namespaceName string) bool {
	namespace, _ := s.storage.GetNamespace(namespaceName)
	return namespace != nil
}

func convertToErrCode(err error) uint32 {
	if errors.Is(err, model.ErrorTokenNotExist) {
		return api.TokenNotExisted
	}
	if errors.Is(err, model.ErrorTokenDisabled) {
		return api.TokenDisabled
	}
	return api.NotAllowedAccess
}
