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

// UpdateAttendanceUserStatsView 更新开发者定制的日度统计或月度统计的统计报表表头设置信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/update
func (r *AttendanceService) UpdateAttendanceUserStatsView(ctx context.Context, request *UpdateAttendanceUserStatsViewReq, options ...MethodOptionFunc) (*UpdateAttendanceUserStatsViewResp, *Response, error) {
	if r.cli.mock.mockAttendanceUpdateAttendanceUserStatsView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#UpdateAttendanceUserStatsView mock enable")
		return r.cli.mock.mockAttendanceUpdateAttendanceUserStatsView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "UpdateAttendanceUserStatsView",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_stats_views/:user_stats_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateAttendanceUserStatsViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceUpdateAttendanceUserStatsView mock AttendanceUpdateAttendanceUserStatsView method
func (r *Mock) MockAttendanceUpdateAttendanceUserStatsView(f func(ctx context.Context, request *UpdateAttendanceUserStatsViewReq, options ...MethodOptionFunc) (*UpdateAttendanceUserStatsViewResp, *Response, error)) {
	r.mockAttendanceUpdateAttendanceUserStatsView = f
}

// UnMockAttendanceUpdateAttendanceUserStatsView un-mock AttendanceUpdateAttendanceUserStatsView method
func (r *Mock) UnMockAttendanceUpdateAttendanceUserStatsView() {
	r.mockAttendanceUpdateAttendanceUserStatsView = nil
}

// UpdateAttendanceUserStatsViewReq ...
type UpdateAttendanceUserStatsViewReq struct {
	UserStatsViewID string                                `path:"user_stats_view_id" json:"-"` // 用户视图 ID, 获取方式: 1）[查询统计设置](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query), 示例值: "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09"
	EmployeeType    EmployeeType                          `query:"employee_type" json:"-"`     // 员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	View            *UpdateAttendanceUserStatsViewReqView `json:"view,omitempty"`              // 统计设置
}

// UpdateAttendanceUserStatsViewReqView ...
type UpdateAttendanceUserStatsViewReqView struct {
	ViewID    string                                      `json:"view_id,omitempty"`    // 视图 ID, 示例值: "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09"
	StatsType string                                      `json:"stats_type,omitempty"` // 视图类型, 示例值: "month", 可选值有: daily: 日度统计, month: 月度统计
	UserID    string                                      `json:"user_id,omitempty"`    // 操作者的用户id, 示例值: "ec8ddg56"
	Items     []*UpdateAttendanceUserStatsViewReqViewItem `json:"items,omitempty"`      // 用户设置字段
}

// UpdateAttendanceUserStatsViewReqViewItem ...
type UpdateAttendanceUserStatsViewReqViewItem struct {
	Code       string                                               `json:"code,omitempty"`        // 标题编号, 示例值: "522"
	ChildItems []*UpdateAttendanceUserStatsViewReqViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

// UpdateAttendanceUserStatsViewReqViewItemChildItem ...
type UpdateAttendanceUserStatsViewReqViewItemChildItem struct {
	Code  string `json:"code,omitempty"`  // 子标题编号, 示例值: "50101"
	Value string `json:"value,omitempty"` // 开关字段, 0: 关闭, 1: 开启（非开关字段场景: code = 51501 可选值为1-6）, 示例值: "0"
}

// UpdateAttendanceUserStatsViewResp ...
type UpdateAttendanceUserStatsViewResp struct {
	View *UpdateAttendanceUserStatsViewRespView `json:"view,omitempty"` // 视图
}

// UpdateAttendanceUserStatsViewRespView ...
type UpdateAttendanceUserStatsViewRespView struct {
	ViewID    string                                       `json:"view_id,omitempty"`    // 视图 ID
	StatsType string                                       `json:"stats_type,omitempty"` // 视图类型, 可选值有: daily: 日度统计, month: 月度统计
	UserID    string                                       `json:"user_id,omitempty"`    // 操作者的用户id
	Items     []*UpdateAttendanceUserStatsViewRespViewItem `json:"items,omitempty"`      // 用户设置字段
}

// UpdateAttendanceUserStatsViewRespViewItem ...
type UpdateAttendanceUserStatsViewRespViewItem struct {
	Code       string                                                `json:"code,omitempty"`        // 标题编号
	Title      string                                                `json:"title,omitempty"`       // 标题名称
	ChildItems []*UpdateAttendanceUserStatsViewRespViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

// UpdateAttendanceUserStatsViewRespViewItemChildItem ...
type UpdateAttendanceUserStatsViewRespViewItemChildItem struct {
	Code       string `json:"code,omitempty"`        // 子标题编号
	Value      string `json:"value,omitempty"`       // 开关字段, 0: 关闭, 1: 开启（非开关字段场景: code = 51501 可选值为1-6）
	Title      string `json:"title,omitempty"`       // 子标题名称
	ColumnType int64  `json:"column_type,omitempty"` // 列类型
	ReadOnly   bool   `json:"read_only,omitempty"`   // 是否只读
	MinValue   string `json:"min_value,omitempty"`   // 最小值
	MaxValue   string `json:"max_value,omitempty"`   // 最大值
}

// updateAttendanceUserStatsViewResp ...
type updateAttendanceUserStatsViewResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAttendanceUserStatsViewResp `json:"data,omitempty"`
}
