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

// GetMessage 通过 message_id 查询消息内容。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人必须在群组中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get
func (r *MessageService) GetMessage(ctx context.Context, request *GetMessageReq, options ...MethodOptionFunc) (*GetMessageResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetMessage mock enable")
		return r.cli.mock.mockMessageGetMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessage",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessage mock MessageGetMessage method
func (r *Mock) MockMessageGetMessage(f func(ctx context.Context, request *GetMessageReq, options ...MethodOptionFunc) (*GetMessageResp, *Response, error)) {
	r.mockMessageGetMessage = f
}

// UnMockMessageGetMessage un-mock MessageGetMessage method
func (r *Mock) UnMockMessageGetMessage() {
	r.mockMessageGetMessage = nil
}

// GetMessageReq ...
type GetMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待获取消息内容的消息的ID, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: "om_dc13264520392913993dd051dba21dcf"
}

// GetMessageResp ...
type GetMessageResp struct {
	Items []*GetMessageRespItem `json:"items,omitempty"` // --
}

// GetMessageRespItem ...
type GetMessageRespItem struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[接收消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳（毫秒）
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳（毫秒）
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回或删除
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被@的用户或机器人的id列表
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 合并转发消息中, 上一层级的消息id message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
}

// getMessageResp ...
type getMessageResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageResp `json:"data,omitempty"`
}
