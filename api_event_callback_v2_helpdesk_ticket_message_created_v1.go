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

// EventV2HelpdeskTicketMessageCreatedV1 该消息事件属于工单消息事件。需使用订阅接口订阅: [事件订阅](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/overview)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_message/events/created
func (r *EventCallbackService) HandlerEventV2HelpdeskTicketMessageCreatedV1(f EventV2HelpdeskTicketMessageCreatedV1Handler) {
	r.cli.eventHandler.eventV2HelpdeskTicketMessageCreatedV1Handler = f
}

// EventV2HelpdeskTicketMessageCreatedV1Handler event EventV2HelpdeskTicketMessageCreatedV1 handler
type EventV2HelpdeskTicketMessageCreatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2HelpdeskTicketMessageCreatedV1) (string, error)

// EventV2HelpdeskTicketMessageCreatedV1 ...
type EventV2HelpdeskTicketMessageCreatedV1 struct {
	TicketMessageID string                                         `json:"ticket_message_id,omitempty"` // 工单消息ID
	MessageID       string                                         `json:"message_id,omitempty"`        // chat消息open ID
	MsgType         MsgType                                        `json:"msg_type,omitempty"`          // 消息类型；text: 纯文本
	Position        string                                         `json:"position,omitempty"`          // 消息位置
	SenderID        *EventV2HelpdeskTicketMessageCreatedV1SenderID `json:"sender_id,omitempty"`         // 用户 ID
	SenderType      int64                                          `json:"sender_type,omitempty"`       // 发送者类型 1: 机器人；2: 用户；3: 客服
	Text            string                                         `json:"text,omitempty"`              // 内容
	Ticket          *EventV2HelpdeskTicketMessageCreatedV1Ticket   `json:"ticket,omitempty"`            // 工单信息
	EventID         string                                         `json:"event_id,omitempty"`          // 消息事件ID
	ChatID          string                                         `json:"chat_id,omitempty"`           // 会话ID
	Content         *EventV2HelpdeskTicketMessageCreatedV1Content  `json:"content,omitempty"`           // 内容详情
}

// EventV2HelpdeskTicketMessageCreatedV1Content ...
type EventV2HelpdeskTicketMessageCreatedV1Content struct {
	Content   string   `json:"content,omitempty"`    // 内容
	MsgType   MsgType  `json:"msg_type,omitempty"`   // 消息类型；text: 纯文本；post: 富文本；image: 图片
	ImageKeys []string `json:"image_keys,omitempty"` // 图片ID
	ImageKey  string   `json:"image_key,omitempty"`  // 图片ID
}

// EventV2HelpdeskTicketMessageCreatedV1SenderID ...
type EventV2HelpdeskTicketMessageCreatedV1SenderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2HelpdeskTicketMessageCreatedV1Ticket ...
type EventV2HelpdeskTicketMessageCreatedV1Ticket struct {
	TicketID string `json:"ticket_id,omitempty"` // 工单ID, [可以从工单列表里面取](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list), [也可以订阅工单创建事件获取](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/events/created)
	Stage    int64  `json:"stage,omitempty"`     // 工单阶段, 1: bot, 2: 人工
	Status   int64  `json:"status,omitempty"`    // 工单状态, 1: 已创建 2: 处理中 3: 排队中 4: 待定 5: 待用户响应 50: 被机器人关闭 51: 被客服关闭 52: 用户自己关闭
}
