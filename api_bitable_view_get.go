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

// GetBitableView 该接口根据 view_id 检索现有视图
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/get
func (r *BitableService) GetBitableView(ctx context.Context, request *GetBitableViewReq, options ...MethodOptionFunc) (*GetBitableViewResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableView mock enable")
		return r.cli.mock.mockBitableGetBitableView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableView",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views/:view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableView mock BitableGetBitableView method
func (r *Mock) MockBitableGetBitableView(f func(ctx context.Context, request *GetBitableViewReq, options ...MethodOptionFunc) (*GetBitableViewResp, *Response, error)) {
	r.mockBitableGetBitableView = f
}

// UnMockBitableGetBitableView un-mock BitableGetBitableView method
func (r *Mock) UnMockBitableGetBitableView() {
	r.mockBitableGetBitableView = nil
}

// GetBitableViewReq ...
type GetBitableViewReq struct {
	AppToken string `path:"app_token" json:"-"` // base app token, 示例值: "bascnCMII2ORej2RItqpZZUNMIe", 最小长度: `1` 字符
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值: "tblsRc9GRRXKqhvW"
	ViewID   string `path:"view_id" json:"-"`   // 视图 ID, 示例值: "vewTpR1urY"
}

// GetBitableViewResp ...
type GetBitableViewResp struct {
	View *GetBitableViewRespView `json:"view,omitempty"` // 视图信息
}

// GetBitableViewRespView ...
type GetBitableViewRespView struct {
	ViewID   string                          `json:"view_id,omitempty"`   // 视图Id
	ViewName string                          `json:"view_name,omitempty"` // 视图名字
	ViewType string                          `json:"view_type,omitempty"` // 视图类型
	Property *GetBitableViewRespViewProperty `json:"property,omitempty"`  // 视图属性
}

// GetBitableViewRespViewProperty ...
type GetBitableViewRespViewProperty struct {
	FilterInfo   *GetBitableViewRespViewPropertyFilterInfo `json:"filter_info,omitempty"`   // 过滤条件
	HiddenFields []string                                  `json:"hidden_fields,omitempty"` // 隐藏字段ID列表
}

// GetBitableViewRespViewPropertyFilterInfo ...
type GetBitableViewRespViewPropertyFilterInfo struct {
	Conjunction      string                                               `json:"conjunction,omitempty"`       // 多个筛选条件的关系, 可选值有: and: 与, or: 或
	Conditions       []*GetBitableViewRespViewPropertyFilterInfoCondition `json:"conditions,omitempty"`        // 筛选条件
	ConditionOmitted bool                                                 `json:"condition_omitted,omitempty"` // 筛选条件是否缺省
}

// GetBitableViewRespViewPropertyFilterInfoCondition ...
type GetBitableViewRespViewPropertyFilterInfoCondition struct {
	FieldID     string `json:"field_id,omitempty"`     // 用于过滤的字段唯一ID
	Operator    string `json:"operator,omitempty"`     // 过滤操作的类型, 可选值有: is: 等于, isNot: 不等于, contains: 包含, doesNotContain: 不包含, isEmpty: 为空, isNotEmpty: 不为空, isGreater: 大于, isGreaterEqual: 大于等于, isLess: 小于, isLessEqual: 小于等于
	Value       string `json:"value,omitempty"`        // 筛选值
	ConditionID string `json:"condition_id,omitempty"` // 过滤条件的唯一ID
	FieldType   string `json:"field_type,omitempty"`   // 用于过滤的字段类型
}

// getBitableViewResp ...
type getBitableViewResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableViewResp `json:"data,omitempty"`
}
