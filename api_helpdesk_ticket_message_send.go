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

// SendHelpdeskTicketMessage 该接口用于工单发送消息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/create
func (r *HelpdeskService) SendHelpdeskTicketMessage(ctx context.Context, request *SendHelpdeskTicketMessageReq, options ...MethodOptionFunc) (*SendHelpdeskTicketMessageResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSendHelpdeskTicketMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SendHelpdeskTicketMessage mock enable")
		return r.cli.mock.mockHelpdeskSendHelpdeskTicketMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "SendHelpdeskTicketMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/tickets/:ticket_id/messages",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(sendHelpdeskTicketMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskSendHelpdeskTicketMessage mock HelpdeskSendHelpdeskTicketMessage method
func (r *Mock) MockHelpdeskSendHelpdeskTicketMessage(f func(ctx context.Context, request *SendHelpdeskTicketMessageReq, options ...MethodOptionFunc) (*SendHelpdeskTicketMessageResp, *Response, error)) {
	r.mockHelpdeskSendHelpdeskTicketMessage = f
}

// UnMockHelpdeskSendHelpdeskTicketMessage un-mock HelpdeskSendHelpdeskTicketMessage method
func (r *Mock) UnMockHelpdeskSendHelpdeskTicketMessage() {
	r.mockHelpdeskSendHelpdeskTicketMessage = nil
}

// SendHelpdeskTicketMessageReq ...
type SendHelpdeskTicketMessageReq struct {
	TicketID string  `path:"ticket_id" json:"-"` // 工单ID, 示例值: "6948728206392295444"
	MsgType  MsgType `json:"msg_type,omitempty"` // 消息类型；text: 纯文本；post: 富文本, 示例值: "post"
	Content  string  `json:"content,omitempty"`  // - 纯文本, 参考[发送文本消息](https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM)中的content；, 富文本, 参考[发送富文本消息](https://open.feishu.cn/document/ukTMukTMukTM/uMDMxEjLzATMx4yMwETM)中的content, 示例值: "{, "msg_type": "post", "content": {, "post": {, "zh_cn": {, "title": "this is title", "content": [, [, {, "tag": "text", "un_escape": true, "text": "第一行:", }, {, "tag": "a", "text": "超链接", "href": "http://www.feishu.cn", }, ], [, {, "tag": "text", "text": "第二行 :", }, {, "tag": "text", "text": "文本测试", }, ], ], }, }, }, }"
}

// SendHelpdeskTicketMessageResp ...
type SendHelpdeskTicketMessageResp struct {
	MessageID string `json:"message_id,omitempty"` // chat消息open ID
}

// sendHelpdeskTicketMessageResp ...
type sendHelpdeskTicketMessageResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *SendHelpdeskTicketMessageResp `json:"data,omitempty"`
}
