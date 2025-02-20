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

// GetHireEmployeeByApplication 通过投递 ID 获取入职信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get_by_application
func (r *HireService) GetHireEmployeeByApplication(ctx context.Context, request *GetHireEmployeeByApplicationReq, options ...MethodOptionFunc) (*GetHireEmployeeByApplicationResp, *Response, error) {
	if r.cli.mock.mockHireGetHireEmployeeByApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireEmployeeByApplication mock enable")
		return r.cli.mock.mockHireGetHireEmployeeByApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireEmployeeByApplication",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/employees/get_by_application",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireEmployeeByApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireEmployeeByApplication mock HireGetHireEmployeeByApplication method
func (r *Mock) MockHireGetHireEmployeeByApplication(f func(ctx context.Context, request *GetHireEmployeeByApplicationReq, options ...MethodOptionFunc) (*GetHireEmployeeByApplicationResp, *Response, error)) {
	r.mockHireGetHireEmployeeByApplication = f
}

// UnMockHireGetHireEmployeeByApplication un-mock HireGetHireEmployeeByApplication method
func (r *Mock) UnMockHireGetHireEmployeeByApplication() {
	r.mockHireGetHireEmployeeByApplication = nil
}

// GetHireEmployeeByApplicationReq ...
type GetHireEmployeeByApplicationReq struct {
	ApplicationID    string            `query:"application_id" json:"-"`     // 投递ID, 示例值: "123"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 的类型, 示例值: "department_id", 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_admin_department_id: 以 people_admin_department_id 来标识部门, 默认值: `people_admin_department_id`
}

// GetHireEmployeeByApplicationResp ...
type GetHireEmployeeByApplicationResp struct {
	Employee *GetHireEmployeeByApplicationRespEmployee `json:"employee,omitempty"` // 员工信息
}

// GetHireEmployeeByApplicationRespEmployee ...
type GetHireEmployeeByApplicationRespEmployee struct {
	ID                     string       `json:"id,omitempty"`                       // 员工ID
	ApplicationID          string       `json:"application_id,omitempty"`           // 投递ID
	OnboardStatus          int64        `json:"onboard_status,omitempty"`           // 入职状态, 可选值有: 1: 已入职, 2: 已离职
	ConversionStatus       int64        `json:"conversion_status,omitempty"`        // 转正状态, 可选值有: 1: 未转正, 2: 已转正
	OnboardTime            int64        `json:"onboard_time,omitempty"`             // 实际入职时间
	ExpectedConversionTime int64        `json:"expected_conversion_time,omitempty"` // 预期转正时间
	ActualConversionTime   int64        `json:"actual_conversion_time,omitempty"`   // 实际转正时间
	OverboardTime          int64        `json:"overboard_time,omitempty"`           // 离职时间
	OverboardNote          string       `json:"overboard_note,omitempty"`           // 离职原因
	OnboardCityCode        string       `json:"onboard_city_code,omitempty"`        // 办公地点
	Department             string       `json:"department,omitempty"`               // 入职部门
	Leader                 string       `json:"leader,omitempty"`                   // 直属上级
	Sequence               string       `json:"sequence,omitempty"`                 // 序列
	Level                  string       `json:"level,omitempty"`                    // 职级
	EmployeeType           EmployeeType `json:"employee_type,omitempty"`            // 员工类型
}

// getHireEmployeeByApplicationResp ...
type getHireEmployeeByApplicationResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetHireEmployeeByApplicationResp `json:"data,omitempty"`
}
