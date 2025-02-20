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

// GetContactJobLevel 该接口可以获取单个职级的信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_level/get
func (r *ContactService) GetContactJobLevel(ctx context.Context, request *GetContactJobLevelReq, options ...MethodOptionFunc) (*GetContactJobLevelResp, *Response, error) {
	if r.cli.mock.mockContactGetContactJobLevel != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactJobLevel mock enable")
		return r.cli.mock.mockContactGetContactJobLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactJobLevel",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_levels/:job_level_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactJobLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactJobLevel mock ContactGetContactJobLevel method
func (r *Mock) MockContactGetContactJobLevel(f func(ctx context.Context, request *GetContactJobLevelReq, options ...MethodOptionFunc) (*GetContactJobLevelResp, *Response, error)) {
	r.mockContactGetContactJobLevel = f
}

// UnMockContactGetContactJobLevel un-mock ContactGetContactJobLevel method
func (r *Mock) UnMockContactGetContactJobLevel() {
	r.mockContactGetContactJobLevel = nil
}

// GetContactJobLevelReq ...
type GetContactJobLevelReq struct {
	JobLevelID string `path:"job_level_id" json:"-"` // 职级ID, 示例值: "mga5oa8ayjlp9rb"
}

// GetContactJobLevelResp ...
type GetContactJobLevelResp struct {
	JobLevel *GetContactJobLevelRespJobLevel `json:"job_level,omitempty"` // 职级信息
}

// GetContactJobLevelRespJobLevel ...
type GetContactJobLevelRespJobLevel struct {
	Name            string                                           `json:"name,omitempty"`             // 职级名称
	Description     string                                           `json:"description,omitempty"`      // 职级描述
	Order           int64                                            `json:"order,omitempty"`            // 职级的排序, 可填入自然数100-100000的数值, 系统按照数值大小从小到大排序。不填写该字段时, 默认新增排序在当前职级列表中最后位（最大值）
	Status          bool                                             `json:"status,omitempty"`           // 是否启用
	JobLevelID      string                                           `json:"job_level_id,omitempty"`     // 职级ID
	I18nName        []*GetContactJobLevelRespJobLevelI18nName        `json:"i18n_name,omitempty"`        // 多语言名称
	I18nDescription []*GetContactJobLevelRespJobLevelI18nDescription `json:"i18n_description,omitempty"` // 多语言描述
}

// GetContactJobLevelRespJobLevelI18nDescription ...
type GetContactJobLevelRespJobLevelI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// GetContactJobLevelRespJobLevelI18nName ...
type GetContactJobLevelRespJobLevelI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// getContactJobLevelResp ...
type getContactJobLevelResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetContactJobLevelResp `json:"data,omitempty"`
}
