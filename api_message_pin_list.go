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

// GetMessagePinList 获取所在群内指定时间范围内的所有 Pin 消息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 获取Pin消息时, 机器人必须在群组中
// - 获取的Pin消息按Pin的创建时间降序排列
// - 接口默认限流为[50 QPS]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/pin/list
func (r *MessageService) GetMessagePinList(ctx context.Context, request *GetMessagePinListReq, options ...MethodOptionFunc) (*GetMessagePinListResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessagePinList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetMessagePinList mock enable")
		return r.cli.mock.mockMessageGetMessagePinList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessagePinList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/pins",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getMessagePinListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessagePinList mock MessageGetMessagePinList method
func (r *Mock) MockMessageGetMessagePinList(f func(ctx context.Context, request *GetMessagePinListReq, options ...MethodOptionFunc) (*GetMessagePinListResp, *Response, error)) {
	r.mockMessageGetMessagePinList = f
}

// UnMockMessageGetMessagePinList un-mock MessageGetMessagePinList method
func (r *Mock) UnMockMessageGetMessagePinList() {
	r.mockMessageGetMessagePinList = nil
}

// GetMessagePinListReq ...
type GetMessagePinListReq struct {
	ChatID    string  `query:"chat_id" json:"-"`    // 待获取Pin消息的Chat ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_234jsi43d3ssi993d43545f"
	StartTime *string `query:"start_time" json:"-"` // Pin信息的起始时间（毫秒级时间戳）。若未填写默认获取到群聊内最早的Pin信息, 示例值: "1658632251800"
	EndTime   *string `query:"end_time" json:"-"`   // Pin信息的结束时间（毫秒级时间戳）。若未填写默认从群聊内最新的Pin信息开始获取, 注意: `end_time`值应大于`start_time`值, 示例值: "1658731646425"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ["
}

// GetMessagePinListResp ...
type GetMessagePinListResp struct {
	Items     []*GetMessagePinListRespItem `json:"items,omitempty"`      // Pin的操作信息
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetMessagePinListRespItem ...
type GetMessagePinListRespItem struct {
	MessageID      string `json:"message_id,omitempty"`       // Pin的消息ID
	ChatID         string `json:"chat_id,omitempty"`          // Pin消息所在的群聊ID
	OperatorID     string `json:"operator_id,omitempty"`      // Pin的操作人ID
	OperatorIDType IDType `json:"operator_id_type,omitempty"` // Pin的操作人ID类型。当Pin的操作人为用户时, 为]open_id[；当Pin的操作人为机器人时, 为]app_id[
	CreateTime     string `json:"create_time,omitempty"`      // Pin的创建时间（毫秒级时间戳）
}

// getMessagePinListResp ...
type getMessagePinListResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetMessagePinListResp `json:"data,omitempty"`
}
