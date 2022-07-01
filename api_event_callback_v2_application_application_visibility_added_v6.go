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

// EventV2ApplicationApplicationVisibilityAddedV6 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 仅当企业的用户通过「普通成员安装」方式获得应用可用性时推送此事件。
// - 订阅前提: 需要是应用商店应用
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app-availability-scope-extended
func (r *EventCallbackService) HandlerEventV2ApplicationApplicationVisibilityAddedV6(f EventV2ApplicationApplicationVisibilityAddedV6Handler) {
	r.cli.eventHandler.eventV2ApplicationApplicationVisibilityAddedV6Handler = f
}

// EventV2ApplicationApplicationVisibilityAddedV6Handler event EventV2ApplicationApplicationVisibilityAddedV6 handler
type EventV2ApplicationApplicationVisibilityAddedV6Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ApplicationApplicationVisibilityAddedV6) (string, error)

// EventV2ApplicationApplicationVisibilityAddedV6 ...
type EventV2ApplicationApplicationVisibilityAddedV6 struct {
	Source int64                                                 `json:"source,omitempty"` // 事件来源, 为 1 时代表通过普通成员安装增加可见性. 如: 1
	Users  []*EventV2ApplicationApplicationVisibilityAddedV6User `json:"users,omitempty"`
}

// EventV2ApplicationApplicationVisibilityAddedV6User ...
type EventV2ApplicationApplicationVisibilityAddedV6User struct {
	UserID map[string]interface{} `json:"user_id,omitempty"` // 开通的用户 id
}
