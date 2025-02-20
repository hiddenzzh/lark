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

// UpdateBitableTableFormField 该接口用于更新表单中的问题项
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-form-field/patch
func (r *BitableService) UpdateBitableTableFormField(ctx context.Context, request *UpdateBitableTableFormFieldReq, options ...MethodOptionFunc) (*UpdateBitableTableFormFieldResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateBitableTableFormField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#UpdateBitableTableFormField mock enable")
		return r.cli.mock.mockBitableUpdateBitableTableFormField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "UpdateBitableTableFormField",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id/fields/:field_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateBitableTableFormFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableUpdateBitableTableFormField mock BitableUpdateBitableTableFormField method
func (r *Mock) MockBitableUpdateBitableTableFormField(f func(ctx context.Context, request *UpdateBitableTableFormFieldReq, options ...MethodOptionFunc) (*UpdateBitableTableFormFieldResp, *Response, error)) {
	r.mockBitableUpdateBitableTableFormField = f
}

// UnMockBitableUpdateBitableTableFormField un-mock BitableUpdateBitableTableFormField method
func (r *Mock) UnMockBitableUpdateBitableTableFormField() {
	r.mockBitableUpdateBitableTableFormField = nil
}

// UpdateBitableTableFormFieldReq ...
type UpdateBitableTableFormFieldReq struct {
	AppToken    string  `path:"app_token" json:"-"`     // 多维表格文档 Token, 示例值: "bascnCMII2ORej2RItqpZZUNMIe"
	TableID     string  `path:"table_id" json:"-"`      // 表格 ID, 示例值: "tblsRc9GRRXKqhvW"
	FormID      string  `path:"form_id" json:"-"`       // 表单 ID, 示例值: "vewTpR1urY"
	FieldID     string  `path:"field_id" json:"-"`      // 表单问题 ID, 示例值: "fldjX7dUj5"
	PreFieldID  *string `json:"pre_field_id,omitempty"` // 上一个表单问题 ID, 用于支持调整表单问题的顺序, 通过前一个表单问题的 field_id 来确定位置；如果 pre_field_id 为空字符串, 则说明要排到首个表单问题, 示例值: "fldjX7dUj5"
	Title       *string `json:"title,omitempty"`        // 表单问题, 示例值: "多行文本"
	Description *string `json:"description,omitempty"`  // 问题描述, 示例值: "多行文本描述"
	Required    *bool   `json:"required,omitempty"`     // 是否必填, 示例值: true
	Visible     *bool   `json:"visible,omitempty"`      // 是否可见, 当值为 false 时, 不允许更新其他字段, 示例值: true
}

// UpdateBitableTableFormFieldResp ...
type UpdateBitableTableFormFieldResp struct {
	Field *UpdateBitableTableFormFieldRespField `json:"field,omitempty"` // 更新后的表单问题项
}

// UpdateBitableTableFormFieldRespField ...
type UpdateBitableTableFormFieldRespField struct {
	PreFieldID  string `json:"pre_field_id,omitempty"` // 上一个表单问题 ID, 用于支持调整表单问题的顺序, 通过前一个表单问题的 field_id 来确定位置；如果 pre_field_id 为空字符串, 则说明要排到首个表单问题
	Title       string `json:"title,omitempty"`        // 表单问题
	Description string `json:"description,omitempty"`  // 问题描述
	Required    bool   `json:"required,omitempty"`     // 是否必填
	Visible     bool   `json:"visible,omitempty"`      // 是否可见, 当值为 false 时, 不允许更新其他字段。
}

// updateBitableTableFormFieldResp ...
type updateBitableTableFormFieldResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBitableTableFormFieldResp `json:"data,omitempty"`
}
