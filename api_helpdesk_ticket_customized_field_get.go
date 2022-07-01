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

// GetHelpdeskTicketCustomizedField 该接口用于获取工单自定义字段详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/get-ticket-customized-field
func (r *HelpdeskService) GetHelpdeskTicketCustomizedField(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicketCustomizedField",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskTicketCustomizedField mock HelpdeskGetHelpdeskTicketCustomizedField method
func (r *Mock) MockHelpdeskGetHelpdeskTicketCustomizedField(f func(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicketCustomizedField = f
}

// UnMockHelpdeskGetHelpdeskTicketCustomizedField un-mock HelpdeskGetHelpdeskTicketCustomizedField method
func (r *Mock) UnMockHelpdeskGetHelpdeskTicketCustomizedField() {
	r.mockHelpdeskGetHelpdeskTicketCustomizedField = nil
}

// GetHelpdeskTicketCustomizedFieldReq ...
type GetHelpdeskTicketCustomizedFieldReq struct {
	TicketCustomizedFieldID string `path:"ticket_customized_field_id" json:"-"` // 工单自定义字段ID, 示例值: "6948728206392295444"
}

// GetHelpdeskTicketCustomizedFieldResp ...
type GetHelpdeskTicketCustomizedFieldResp struct {
	TicketCustomizedFieldID string                                         `json:"ticket_customized_field_id,omitempty"` // 工单自定义字段ID
	HelpdeskID              string                                         `json:"helpdesk_id,omitempty"`                // 服务台ID
	KeyName                 string                                         `json:"key_name,omitempty"`                   // 键名
	DisplayName             string                                         `json:"display_name,omitempty"`               // 名称
	Position                string                                         `json:"position,omitempty"`                   // 字段在列表后台管理列表中的位置
	FieldType               string                                         `json:"field_type,omitempty"`                 // 类型
	Description             string                                         `json:"description,omitempty"`                // 描述
	Visible                 bool                                           `json:"visible,omitempty"`                    // 是否可见
	Editable                bool                                           `json:"editable,omitempty"`                   // 是否可以修改
	Required                bool                                           `json:"required,omitempty"`                   // 是否必填
	CreatedAt               string                                         `json:"created_at,omitempty"`                 // 创建时间
	UpdatedAt               string                                         `json:"updated_at,omitempty"`                 // 更新时间
	CreatedBy               *GetHelpdeskTicketCustomizedFieldRespCreatedBy `json:"created_by,omitempty"`                 // 创建用户
	UpdatedBy               *GetHelpdeskTicketCustomizedFieldRespUpdatedBy `json:"updated_by,omitempty"`                 // 更新用户
	DropdownAllowMultiple   bool                                           `json:"dropdown_allow_multiple,omitempty"`    // 是否支持多选, 仅在字段类型是dropdown的时候有效
}

// GetHelpdeskTicketCustomizedFieldRespCreatedBy ...
type GetHelpdeskTicketCustomizedFieldRespCreatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

// GetHelpdeskTicketCustomizedFieldRespUpdatedBy ...
type GetHelpdeskTicketCustomizedFieldRespUpdatedBy struct {
	ID              string                  `json:"id,omitempty"`               // 用户ID
	AvatarURL       string                  `json:"avatar_url,omitempty"`       // 用户头像url
	Name            string                  `json:"name,omitempty"`             // 用户名
	Email           string                  `json:"email,omitempty"`            // 用户邮箱
	DropdownOptions *HelpdeskDropdownOption `json:"dropdown_options,omitempty"` // 下拉列表选项
}

// getHelpdeskTicketCustomizedFieldResp ...
type getHelpdeskTicketCustomizedFieldResp struct {
	Code int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskTicketCustomizedFieldResp `json:"data,omitempty"`
}
