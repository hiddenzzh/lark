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

// GetHelpdeskFAQList 该接口用于获取服务台知识库详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/list
func (r *HelpdeskService) GetHelpdeskFAQList(ctx context.Context, request *GetHelpdeskFAQListReq, options ...MethodOptionFunc) (*GetHelpdeskFAQListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskFAQList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskFAQList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskFAQList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskFAQList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskFAQListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskFAQList mock HelpdeskGetHelpdeskFAQList method
func (r *Mock) MockHelpdeskGetHelpdeskFAQList(f func(ctx context.Context, request *GetHelpdeskFAQListReq, options ...MethodOptionFunc) (*GetHelpdeskFAQListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskFAQList = f
}

// UnMockHelpdeskGetHelpdeskFAQList un-mock HelpdeskGetHelpdeskFAQList method
func (r *Mock) UnMockHelpdeskGetHelpdeskFAQList() {
	r.mockHelpdeskGetHelpdeskFAQList = nil
}

// GetHelpdeskFAQListReq ...
type GetHelpdeskFAQListReq struct {
	CategoryID *string `query:"category_id" json:"-"` // 知识库分类ID, 示例值: "6856395522433908739"
	Status     *string `query:"status" json:"-"`      // 搜索条件: 知识库状态 1:在线 0:删除, 可恢复 2: 删除, 不可恢复, 示例值: "1"
	Search     *string `query:"search" json:"-"`      // 搜索条件: 关键词, 匹配问题标题, 问题关键字, 用户姓名, 示例值: "点餐"
	PageToken  *string `query:"page_token" json:"-"`  // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "6856395634652479491"
	PageSize   *int64  `query:"page_size" json:"-"`   // 分页大小, 示例值: 10, 最大值: `100`
}

// GetHelpdeskFAQListResp ...
type GetHelpdeskFAQListResp struct {
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                        `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	PageSize  int64                         `json:"page_size,omitempty"`  // 实际返回的FAQ数量
	Total     int64                         `json:"total,omitempty"`      // 总数
	Items     []*GetHelpdeskFAQListRespItem `json:"items,omitempty"`      // 知识库列表
}

// GetHelpdeskFAQListRespItem ...
type GetHelpdeskFAQListRespItem struct {
	FAQID          string                                `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                                `json:"id,omitempty"`              // 知识库旧版ID, 请使用faq_id
	HelpdeskID     string                                `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                                `json:"question,omitempty"`        // 问题
	Answer         string                                `json:"answer,omitempty"`          // 答案
	AnswerRichtext string                                `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64                                 `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64                                 `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory                   `json:"categories,omitempty"`      // 分类
	Tags           []string                              `json:"tags,omitempty"`            // 相似问题列表
	ExpireTime     int64                                 `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *GetHelpdeskFAQListRespItemUpdateUser `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *GetHelpdeskFAQListRespItemCreateUser `json:"create_user,omitempty"`     // 创建用户
}

// GetHelpdeskFAQListRespItemCreateUser ...
type GetHelpdeskFAQListRespItemCreateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

// GetHelpdeskFAQListRespItemUpdateUser ...
type GetHelpdeskFAQListRespItemUpdateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

// getHelpdeskFAQListResp ...
type getHelpdeskFAQListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskFAQListResp `json:"data,omitempty"`
}
