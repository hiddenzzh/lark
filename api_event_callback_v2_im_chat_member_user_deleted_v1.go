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

// EventV2IMChatMemberUserDeletedV1 用户主动退群或被移出群聊时推送事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=chat.member.user&event=deleted)
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)并且机器人所在群发生上述变化
// - 机器人需要订阅 [消息与群组] 分类下的 [用户主动退群或被移出群聊] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/deleted
func (r *EventCallbackService) HandlerEventV2IMChatMemberUserDeletedV1(f EventV2IMChatMemberUserDeletedV1Handler) {
	r.cli.eventHandler.eventV2IMChatMemberUserDeletedV1Handler = f
}

// EventV2IMChatMemberUserDeletedV1Handler event EventV2IMChatMemberUserDeletedV1 handler
type EventV2IMChatMemberUserDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatMemberUserDeletedV1) (string, error)

// EventV2IMChatMemberUserDeletedV1 ...
type EventV2IMChatMemberUserDeletedV1 struct {
	ChatID            string                                      `json:"chat_id,omitempty"`             // 群组 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	OperatorID        *EventV2IMChatMemberUserDeletedV1OperatorID `json:"operator_id,omitempty"`         // 用户 ID
	External          bool                                        `json:"external,omitempty"`            // 是否是外部群
	OperatorTenantKey string                                      `json:"operator_tenant_key,omitempty"` // 操作者租户 Key
	Users             []*EventV2IMChatMemberUserDeletedV1User     `json:"users,omitempty"`               // 被移除用户列表
}

// EventV2IMChatMemberUserDeletedV1OperatorID ...
type EventV2IMChatMemberUserDeletedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMChatMemberUserDeletedV1User ...
type EventV2IMChatMemberUserDeletedV1User struct {
	Name      string                                      `json:"name,omitempty"`       // 用户名字
	TenantKey string                                      `json:"tenant_key,omitempty"` // 租户 Key
	UserID    *EventV2IMChatMemberUserDeletedV1UserUserID `json:"user_id,omitempty"`    // 用户 ID
}

// EventV2IMChatMemberUserDeletedV1UserUserID ...
type EventV2IMChatMemberUserDeletedV1UserUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
