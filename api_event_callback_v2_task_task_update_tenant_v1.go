// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// EventV2TaskTaskUpdateTenantV1 APP 订阅此事件后可接收到该 APP 所在租户的所有来源接口创建的任务的变更事件。事件体为发生变更任务的相关用户的 open_id, 可用此 open_id, 通过 获取任务列表接口获取与该用户相关的所有任务。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=task&version=v1&resource=task&event=update_tenant)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/events/update_tenant
func (r *EventCallbackService) HandlerEventV2TaskTaskUpdateTenantV1(f EventV2TaskTaskUpdateTenantV1Handler) {
	r.cli.eventHandler.eventV2TaskTaskUpdateTenantV1Handler = f
}

// EventV2TaskTaskUpdateTenantV1Handler event EventV2TaskTaskUpdateTenantV1 handler
type EventV2TaskTaskUpdateTenantV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2TaskTaskUpdateTenantV1) (string, error)

// EventV2TaskTaskUpdateTenantV1 ...
type EventV2TaskTaskUpdateTenantV1 struct {
	UserIDList *EventV2TaskTaskUpdateTenantV1UserIDList `json:"user_id_list,omitempty"` // 用户 ID 列表
	TaskID     string                                   `json:"task_id,omitempty"`      // 任务的id
	ObjectType string                                   `json:"object_type,omitempty"`  // 变更的数据类型
	EventType  string                                   `json:"event_type,omitempty"`   // 事件类型
}

// EventV2TaskTaskUpdateTenantV1UserIDList ...
type EventV2TaskTaskUpdateTenantV1UserIDList struct {
	UserIDList []*EventV2TaskTaskUpdateTenantV1UserIDListUserID `json:"user_id_list,omitempty"` // 用户 ID 列表
}

// EventV2TaskTaskUpdateTenantV1UserIDListUserID ...
type EventV2TaskTaskUpdateTenantV1UserIDListUserID struct {
	UnionID string `json:"union_id,omitempty"` // 忽略此字段
	UserID  string `json:"user_id,omitempty"`  // 忽略此字段
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open_id
}
