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

// DeleteMessagePin 移除一条指定消息的 Pin。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 移除Pin消息时, 机器人必须在对应的群组中
// - 若消息未被Pin或已被撤回, 返回成功信息
// - 不能移除一条对操作者不可见的Pin消息
// - 对同一条消息移除Pin的操作不能超过[5 QPS]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/pin/delete
func (r *MessageService) DeleteMessagePin(ctx context.Context, request *DeleteMessagePinReq, options ...MethodOptionFunc) (*DeleteMessagePinResp, *Response, error) {
	if r.cli.mock.mockMessageDeleteMessagePin != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#DeleteMessagePin mock enable")
		return r.cli.mock.mockMessageDeleteMessagePin(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "DeleteMessagePin",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/pins/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteMessagePinResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageDeleteMessagePin mock MessageDeleteMessagePin method
func (r *Mock) MockMessageDeleteMessagePin(f func(ctx context.Context, request *DeleteMessagePinReq, options ...MethodOptionFunc) (*DeleteMessagePinResp, *Response, error)) {
	r.mockMessageDeleteMessagePin = f
}

// UnMockMessageDeleteMessagePin un-mock MessageDeleteMessagePin method
func (r *Mock) UnMockMessageDeleteMessagePin() {
	r.mockMessageDeleteMessagePin = nil
}

// DeleteMessagePinReq ...
type DeleteMessagePinReq struct {
	MessageID string `path:"message_id" json:"-"` // 待移除Pin的消息ID, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: "om_dc13264520392913993dd051dba21dcf"
}

// DeleteMessagePinResp ...
type DeleteMessagePinResp struct {
}

// deleteMessagePinResp ...
type deleteMessagePinResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMessagePinResp `json:"data,omitempty"`
}
