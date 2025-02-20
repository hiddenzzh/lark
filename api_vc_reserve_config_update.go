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

// UpdateVCReserveConfig 更新会议室预定限制。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config/patch
func (r *VCService) UpdateVCReserveConfig(ctx context.Context, request *UpdateVCReserveConfigReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCReserveConfig != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCReserveConfig mock enable")
		return r.cli.mock.mockVCUpdateVCReserveConfig(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "UpdateVCReserveConfig",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reserve_configs/:reserve_config_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateVCReserveConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCReserveConfig mock VCUpdateVCReserveConfig method
func (r *Mock) MockVCUpdateVCReserveConfig(f func(ctx context.Context, request *UpdateVCReserveConfigReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigResp, *Response, error)) {
	r.mockVCUpdateVCReserveConfig = f
}

// UnMockVCUpdateVCReserveConfig un-mock VCUpdateVCReserveConfig method
func (r *Mock) UnMockVCUpdateVCReserveConfig() {
	r.mockVCUpdateVCReserveConfig = nil
}

// UpdateVCReserveConfigReq ...
type UpdateVCReserveConfigReq struct {
	ReserveConfigID    string                                      `path:"reserve_config_id" json:"-"`     // 会议室或层级id, 示例值: "omm_3c5dd7e09bac0c1758fcf9511bd1a771"
	UserIDType         *IDType                                     `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ScopeType          string                                      `json:"scope_type,omitempty"`           // 1 代表层级, 2 代表会议室, 示例值: "2"
	ApprovalConfig     *UpdateVCReserveConfigReqApprovalConfig     `json:"approval_config,omitempty"`      // 预定审批设置
	TimeConfig         *UpdateVCReserveConfigReqTimeConfig         `json:"time_config,omitempty"`          // 预定时间设置
	ReserveScopeConfig *UpdateVCReserveConfigReqReserveScopeConfig `json:"reserve_scope_config,omitempty"` // 预定范围设置
}

// UpdateVCReserveConfigReqApprovalConfig ...
type UpdateVCReserveConfigReqApprovalConfig struct {
	ApprovalSwitch    *int64                                            `json:"approval_switch,omitempty"`    // 预定审批开关: 0 代表关闭, 1 代表打开, <b>说明</b>: 1.  未设置值时不更新原开关的值, 但此时必填  approval_condition, 2.  设置值为 1 时, 必填  approval_condition, 3.  设置值为 0 时整个, approval_config 其他字段均可省略, 示例值: 1, 取值范围: `0` ～ `1`
	ApprovalCondition *int64                                            `json:"approval_condition,omitempty"` // 预定审批条件: 0 代表所有预定均需审批, 1 代表满足条件的需审批, <b>说明</b>: 为 1 时必填 meeting_duration, 示例值: 1, 取值范围: `0` ～ `1`
	MeetingDuration   *float64                                          `json:"meeting_duration,omitempty"`   // 超过 meeting_duration, 的预定需要审批（单位: 小时, 取值范围[0.1-99]）, <b>说明</b>: 1.  当 approval_condition, 为 0, 更新时如果未设置值, 默认更新为 99 ., 2.  传入的值小数点后超过 2 位, 自动四舍五入保留两位, 示例值: 3
	Approvers         []*UpdateVCReserveConfigReqApprovalConfigApprover `json:"approvers,omitempty"`          // 审批人列表, 当打开审批开关时, 至少需要设置一位审批人
}

// UpdateVCReserveConfigReqApprovalConfigApprover ...
type UpdateVCReserveConfigReqApprovalConfigApprover struct {
	UserID string `json:"user_id,omitempty"` // 预定管理员ID, 示例值: "ou_a27b07a9071d90577c0177bcec98f856"
}

// UpdateVCReserveConfigReqReserveScopeConfig ...
type UpdateVCReserveConfigReqReserveScopeConfig struct {
	AllowAllUsers *int64                                                 `json:"allow_all_users,omitempty"` // 可预定成员范围: 0 代表部分成员, 1 代表全部成员, <b>说明</b>: 1.  此值必填, 2.  当设置为 0 时, 至少需要 1 个预定部门或预定人, 示例值: 0, 取值范围: `0` ～ `1`
	AllowUsers    []*UpdateVCReserveConfigReqReserveScopeConfigAllowUser `json:"allow_users,omitempty"`     // 可预定成员列表
	AllowDepts    []*UpdateVCReserveConfigReqReserveScopeConfigAllowDept `json:"allow_depts,omitempty"`     // 可预定部门列表
}

// UpdateVCReserveConfigReqReserveScopeConfigAllowDept ...
type UpdateVCReserveConfigReqReserveScopeConfigAllowDept struct {
	DepartmentID string `json:"department_id,omitempty"` // 预定管理部门ID, 示例值: "od-47d8b570b0a011e9679a755efcc5f61a"
}

// UpdateVCReserveConfigReqReserveScopeConfigAllowUser ...
type UpdateVCReserveConfigReqReserveScopeConfigAllowUser struct {
	UserID string `json:"user_id,omitempty"` // 预定管理员ID, 示例值: "ou_a27b07a9071d90577c0177bcec98f856"
}

// UpdateVCReserveConfigReqTimeConfig ...
type UpdateVCReserveConfigReqTimeConfig struct {
	TimeSwitch    int64   `json:"time_switch,omitempty"`     // 预定时间开关: 0 代表关闭, 1 代表开启, 示例值: 1, 取值范围: `0` ～ `1`
	DaysInAdvance *int64  `json:"days_in_advance,omitempty"` // 最早可提前, days_in_advance 预定会议室（单位: 天, 取值范围[1-730]）, <b>说明</b>: 不填写时, 默认更新为 365, 示例值: 30, 取值范围: `1` ～ `99`
	OpeningHour   *string `json:"opening_hour,omitempty"`    // 开放当天可于, opening_hour 开始预定（单位: 秒, 取值范围[0, 86400]）, <b>说明</b>: 1.  不填写时默认更新为, 28800, 2.  如果填写的值不是 60, 的倍数, 则自动会更新为离其最近的 60 整数倍的值, 示例值: "27900"
	StartTime     *string `json:"start_time,omitempty"`      // 每日可预定时间范围的开始时间（单位: 秒, 取值范围[0, 86400]）, <b>说明</b>: 1.  不填写时, 默认更新为 0, 此时填写的  end_time 不得小于 30, 2.  当 start_time 与, end_time 均填写时, end_time 至少超过, start_time 30, 3.  如果填写的值不是 60 的倍数, 则自动会更新为离其最近的 60 整数倍的值, 示例值: "0"
	EndTime       *string `json:"end_time,omitempty"`        // 每日可预定时间范围结束时间（单位: 秒, 取值范围[0, 86400]）, <b>说明</b>: 1.  不填写时, 默认更新为 86400, 此时填写的, start_time 不得大于等于 86370, 2.  当 start_time 与, end_time 均填写时, end_time 至少要超过, start_time 30, 3.  如果填写的值不是  60 的倍数, 则自动会更新为离其最近的 60 整数倍的值, 示例值: "86400"
	MaxDuration   *int64  `json:"max_duration,omitempty"`    // 单次会议室可预定时长上限（单位: 小时, 取值范围[1, 99]）, <b>说明</b>: 不填写时默认更新为 2, 示例值: 24, 取值范围: `1` ～ `99`
}

// UpdateVCReserveConfigResp ...
type UpdateVCReserveConfigResp struct {
}

// updateVCReserveConfigResp ...
type updateVCReserveConfigResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCReserveConfigResp `json:"data,omitempty"`
}
