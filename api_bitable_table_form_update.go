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

// UpdateBitableTableForm 该接口用于更新表单中的元数据项
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-form/patch
func (r *BitableService) UpdateBitableTableForm(ctx context.Context, request *UpdateBitableTableFormReq, options ...MethodOptionFunc) (*UpdateBitableTableFormResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateBitableTableForm != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#UpdateBitableTableForm mock enable")
		return r.cli.mock.mockBitableUpdateBitableTableForm(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "UpdateBitableTableForm",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateBitableTableFormResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableUpdateBitableTableForm mock BitableUpdateBitableTableForm method
func (r *Mock) MockBitableUpdateBitableTableForm(f func(ctx context.Context, request *UpdateBitableTableFormReq, options ...MethodOptionFunc) (*UpdateBitableTableFormResp, *Response, error)) {
	r.mockBitableUpdateBitableTableForm = f
}

// UnMockBitableUpdateBitableTableForm un-mock BitableUpdateBitableTableForm method
func (r *Mock) UnMockBitableUpdateBitableTableForm() {
	r.mockBitableUpdateBitableTableForm = nil
}

// UpdateBitableTableFormReq ...
type UpdateBitableTableFormReq struct {
	AppToken        string  `path:"app_token" json:"-"`          // 多维表格文档 Token, 示例值: "bascnv1jIEppJdTCn3jOosabcef"
	TableID         string  `path:"table_id" json:"-"`           // 表格 ID, 示例值: "tblz8nadEUdxNMt5"
	FormID          string  `path:"form_id" json:"-"`            // 表单 ID, 示例值: "vew6oMbAa4"
	Name            *string `json:"name,omitempty"`              // 表单名称, 示例值: "表单"
	Description     *string `json:"description,omitempty"`       // 表单描述, 示例值: "表单描述"
	Shared          *bool   `json:"shared,omitempty"`            // 是否开启共享, 示例值: true
	SharedLimit     *string `json:"shared_limit,omitempty"`      // 分享范围限制, 示例值: "tenant_editable", 可选值有: off: 仅邀请的人可填写, tenant_editable: 组织内获得链接的人可填写, anyone_editable: 互联网上获得链接的人可填写
	SubmitLimitOnce *bool   `json:"submit_limit_once,omitempty"` // 填写次数限制一次, 示例值: true
}

// UpdateBitableTableFormResp ...
type UpdateBitableTableFormResp struct {
	Form *UpdateBitableTableFormRespForm `json:"form,omitempty"` // 表单元数据信息
}

// UpdateBitableTableFormRespForm ...
type UpdateBitableTableFormRespForm struct {
	Name            string `json:"name,omitempty"`              // 表单名称
	Description     string `json:"description,omitempty"`       // 表单描述
	Shared          bool   `json:"shared,omitempty"`            // 是否开启共享
	SharedURL       string `json:"shared_url,omitempty"`        // 分享 URL
	SharedLimit     string `json:"shared_limit,omitempty"`      // 分享范围限制, 可选值有: off: 仅邀请的人可填写, tenant_editable: 组织内获得链接的人可填写, anyone_editable: 互联网上获得链接的人可填写
	SubmitLimitOnce bool   `json:"submit_limit_once,omitempty"` // 填写次数限制一次
}

// updateBitableTableFormResp ...
type updateBitableTableFormResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBitableTableFormResp `json:"data,omitempty"`
}
