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

// GetChatModeration 获取群发言模式、可发言用户名单等
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 机器人 或 授权用户 必须在群里
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/get
func (r *ChatService) GetChatModeration(ctx context.Context, request *GetChatModerationReq, options ...MethodOptionFunc) (*GetChatModerationResp, *Response, error) {
	if r.cli.mock.mockChatGetChatModeration != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChatModeration mock enable")
		return r.cli.mock.mockChatGetChatModeration(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatModeration",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/moderation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatModerationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChatModeration mock ChatGetChatModeration method
func (r *Mock) MockChatGetChatModeration(f func(ctx context.Context, request *GetChatModerationReq, options ...MethodOptionFunc) (*GetChatModerationResp, *Response, error)) {
	r.mockChatGetChatModeration = f
}

// UnMockChatGetChatModeration un-mock ChatGetChatModeration method
func (r *Mock) UnMockChatGetChatModeration() {
	r.mockChatGetChatModeration = nil
}

// GetChatModerationReq ...
type GetChatModerationReq struct {
	ChatID     string  `path:"chat_id" json:"-"`       // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ["
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `100`
}

// GetChatModerationResp ...
type GetChatModerationResp struct {
	ModerationSetting string                       `json:"moderation_setting,omitempty"` // 群发言模式（all_members/only_owner/moderator_list, 其中 moderator_list 表示部分用户可发言的模式）
	PageToken         string                       `json:"page_token,omitempty"`         // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore           bool                         `json:"has_more,omitempty"`           // 是否还有更多项
	Items             []*GetChatModerationRespItem `json:"items,omitempty"`              // 可发言用户列表
}

// GetChatModerationRespItem ...
type GetChatModerationRespItem struct {
	UserIDType IDType `json:"user_id_type,omitempty"` // 可发言用户 ID 类型
	UserID     string `json:"user_id,omitempty"`      // 可发言用户 ID
}

// getChatModerationResp ...
type getChatModerationResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetChatModerationResp `json:"data,omitempty"`
}
