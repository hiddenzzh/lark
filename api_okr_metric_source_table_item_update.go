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

// UpdateOKRMetricSourceTableItem - 该接口用于更新某项指标, 接口仅限 OKR 企业版使用。
//
// 更新成功后 OKR 系统会给以下人员发送消息通知:
// - 首次更新目标值的人员
// - 已经将指标添加为 KR、且本次目标值/起始值/支撑的上级有变更的人员, 不包含仅更新了进度值的人员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/metric_source-table-item/patch
func (r *OKRService) UpdateOKRMetricSourceTableItem(ctx context.Context, request *UpdateOKRMetricSourceTableItemReq, options ...MethodOptionFunc) (*UpdateOKRMetricSourceTableItemResp, *Response, error) {
	if r.cli.mock.mockOKRUpdateOKRMetricSourceTableItem != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#UpdateOKRMetricSourceTableItem mock enable")
		return r.cli.mock.mockOKRUpdateOKRMetricSourceTableItem(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "UpdateOKRMetricSourceTableItem",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/metric_sources/:metric_source_id/tables/:metric_table_id/items/:metric_item_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateOKRMetricSourceTableItemResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRUpdateOKRMetricSourceTableItem mock OKRUpdateOKRMetricSourceTableItem method
func (r *Mock) MockOKRUpdateOKRMetricSourceTableItem(f func(ctx context.Context, request *UpdateOKRMetricSourceTableItemReq, options ...MethodOptionFunc) (*UpdateOKRMetricSourceTableItemResp, *Response, error)) {
	r.mockOKRUpdateOKRMetricSourceTableItem = f
}

// UnMockOKRUpdateOKRMetricSourceTableItem un-mock OKRUpdateOKRMetricSourceTableItem method
func (r *Mock) UnMockOKRUpdateOKRMetricSourceTableItem() {
	r.mockOKRUpdateOKRMetricSourceTableItem = nil
}

// UpdateOKRMetricSourceTableItemReq ...
type UpdateOKRMetricSourceTableItemReq struct {
	MetricSourceID     string   `path:"metric_source_id" json:"-"`      // okr指标库id, 示例值: "7041857032248410131"
	MetricTableID      string   `path:"metric_table_id" json:"-"`       // okr指标表id, 示例值: "7041857032248410131"
	MetricItemID       string   `path:"metric_item_id" json:"-"`        // okr指标项id, 示例值: "7041857032248410131"
	UserIDType         *IDType  `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MetricInitialValue *float64 `json:"metric_initial_value,omitempty"` // 指标起始值, 示例值: 1.0
	MetricTargetValue  *float64 `json:"metric_target_value,omitempty"`  // 指标目标值, 示例值: 3.0
	MetricCurrentValue *float64 `json:"metric_current_value,omitempty"` // 指标进度值, 示例值: 2.0
	SupportedUserID    *string  `json:"supported_user_id,omitempty"`    // 指标支撑的上级人员 id, 示例值: "7041857032248410131"
}

// UpdateOKRMetricSourceTableItemResp ...
type UpdateOKRMetricSourceTableItemResp struct {
	MetricItemID       string                                        `json:"metric_item_id,omitempty"`       // 指标表id
	UserID             string                                        `json:"user_id,omitempty"`              // 指标承接人员id
	PeriodID           string                                        `json:"period_id,omitempty"`            // 指标的okr周期
	MetricUnit         *UpdateOKRMetricSourceTableItemRespMetricUnit `json:"metric_unit,omitempty"`          // 指标单位
	MetricInitialValue float64                                       `json:"metric_initial_value,omitempty"` // 指标起始值
	MetricTargetValue  float64                                       `json:"metric_target_value,omitempty"`  // 指标目标值
	MetricCurrentValue float64                                       `json:"metric_current_value,omitempty"` // 指标进度值
	SupportedUserID    string                                        `json:"supported_user_id,omitempty"`    // 指标支撑的上级人员id
	KrID               string                                        `json:"kr_id,omitempty"`                // 指标关联的kr
	UpdatedAt          string                                        `json:"updated_at,omitempty"`           // 更新时间
	UpdatedBy          string                                        `json:"updated_by,omitempty"`           // 更新人
}

// UpdateOKRMetricSourceTableItemRespMetricUnit ...
type UpdateOKRMetricSourceTableItemRespMetricUnit struct {
	ZhCn string `json:"zh_cn,omitempty"` // 指标单位中文
	EnUs string `json:"en_us,omitempty"` // 指标单位英文
	JaJp string `json:"ja_jp,omitempty"` // 指标单位日文
}

// updateOKRMetricSourceTableItemResp ...
type updateOKRMetricSourceTableItemResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *UpdateOKRMetricSourceTableItemResp `json:"data,omitempty"`
}
