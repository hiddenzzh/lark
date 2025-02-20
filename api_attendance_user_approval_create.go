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

// CreateAttendanceUserApproval 由于部分企业使用的是自己的审批系统, 而不是飞书审批系统, 因此员工的请假、加班等数据无法流入到飞书考勤系统中, 导致员工在请假时间段内依然收到打卡提醒, 并且被记为缺卡。
//
// 对于这些只使用飞书考勤系统, 而未使用飞书审批系统的企业, 可以通过考勤开放接口的形式, 将三方审批结果数据回写到飞书考勤系统中。
// 1. 目前支持写入加班、请假、出差和外出这四种审批结果, 写入只会追加(insert), 不会覆盖(update)（开放接口导入的加班假期记录, 在管理后台的假期加班里查不到, 可以在考勤统计报表查看, 或者通过[获取审批通过数据](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/query)来查询）
// 2. 离职人员没有考勤组, 所以写入和返回的时间会有差异
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/create
func (r *AttendanceService) CreateAttendanceUserApproval(ctx context.Context, request *CreateAttendanceUserApprovalReq, options ...MethodOptionFunc) (*CreateAttendanceUserApprovalResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateAttendanceUserApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateAttendanceUserApproval mock enable")
		return r.cli.mock.mockAttendanceCreateAttendanceUserApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateAttendanceUserApproval",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_approvals",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createAttendanceUserApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceCreateAttendanceUserApproval mock AttendanceCreateAttendanceUserApproval method
func (r *Mock) MockAttendanceCreateAttendanceUserApproval(f func(ctx context.Context, request *CreateAttendanceUserApprovalReq, options ...MethodOptionFunc) (*CreateAttendanceUserApprovalResp, *Response, error)) {
	r.mockAttendanceCreateAttendanceUserApproval = f
}

// UnMockAttendanceCreateAttendanceUserApproval un-mock AttendanceCreateAttendanceUserApproval method
func (r *Mock) UnMockAttendanceCreateAttendanceUserApproval() {
	r.mockAttendanceCreateAttendanceUserApproval = nil
}

// CreateAttendanceUserApprovalReq ...
type CreateAttendanceUserApprovalReq struct {
	EmployeeType EmployeeType                                 `query:"employee_type" json:"-"` // 请求体和响应体中的 user_id 的员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的工号
	UserApproval *CreateAttendanceUserApprovalReqUserApproval `json:"user_approval,omitempty"` // 审批信息
}

// CreateAttendanceUserApprovalReqUserApproval ...
type CreateAttendanceUserApprovalReqUserApproval struct {
	UserID        string                                                     `json:"user_id,omitempty"`        // 审批用户 ID, 示例值: "abd754f7"
	Date          string                                                     `json:"date,omitempty"`           // 审批作用日期, 示例值: "20210104"
	Outs          []*CreateAttendanceUserApprovalReqUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateAttendanceUserApprovalReqUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateAttendanceUserApprovalReqUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateAttendanceUserApprovalReqUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

// CreateAttendanceUserApprovalReqUserApprovalLeave ...
type CreateAttendanceUserApprovalReqUserApprovalLeave struct {
	UniqID        *string    `json:"uniq_id,omitempty"`        // 假期类型唯一 ID, 代表一种假期类型, 长度小于 14, * 此ID对应假期类型(即: i18n_names), 因此需要保证唯一, 示例值: "6852582717813440527"
	Unit          int64      `json:"unit,omitempty"`           // 假期时长单位, 示例值: 1, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Interval      int64      `json:"interval,omitempty"`       // 假期时长（单位: 秒）, 暂未开放提供, 待后续提供, 示例值: 28800
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 09:00:00"
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 19:00:00"
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 假期多语言展示, 格式为 map, key 为 ["ch"、"en"、"ja"], 其中 ch 代表中文、en 代表英语、ja 代表日语
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型, 由于飞书客户端支持中、英、日三种语言, 当用户切换语言时, 如果假期名称没有所对应的语言, 会使用默认语言的名称, 示例值: "ch", 可选值有: ch: 中文, en: 英文, ja: 日文
	Reason        string     `json:"reason,omitempty"`         // 请假理由, 必选字段, 示例值: "家里有事"
}

// CreateAttendanceUserApprovalReqUserApprovalOut ...
type CreateAttendanceUserApprovalReqUserApprovalOut struct {
	UniqID        string     `json:"uniq_id,omitempty"`        // 外出类型唯一 ID, 代表一种假期类型, 长度小于 14, * 此ID对应假期类型(即: i18n_names), 因此需要保证唯一, 示例值: "9496E43696967658A512969523E89870"
	Unit          int64      `json:"unit,omitempty"`           // 外出时长单位, 示例值: 1, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Interval      int64      `json:"interval,omitempty"`       // 外出时长（单位: 秒）, 示例值: 28800
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 09:00:00"
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 19:00:00"
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 外出多语言展示, 格式为 map, key 为 ["ch"、"en"、"ja"], 其中 ch 代表中文、en 代表英语、ja 代表日语
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型, 由于飞书客户端支持中、英、日三种语言, 当用户切换语言时, 如果假期名称没有所对应的语言, 会使用默认语言的名称, 示例值: "ch"
	Reason        string     `json:"reason,omitempty"`         // 外出理由, 示例值: "外出办事"
}

// CreateAttendanceUserApprovalReqUserApprovalOvertimeWork ...
type CreateAttendanceUserApprovalReqUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长, 示例值: 1.5
	Unit      int64   `json:"unit,omitempty"`       // 加班时长单位, 示例值: 1, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Category  int64   `json:"category,omitempty"`   // 加班日期类型, 示例值: 2, 可选值有: 1: 工作日, 2: 休息日, 3: 节假日
	Type      int64   `json:"type,omitempty"`       // 加班规则类型, 示例值: 1, 可选值有: 0: 不关联加班规则, 1: 调休, 2: 加班费, 3: 关联加班规则, 没有调休或加班费
	StartTime string  `json:"start_time,omitempty"` // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-09 09:00:00"
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-10 13:00:00"
}

// CreateAttendanceUserApprovalReqUserApprovalTrip ...
type CreateAttendanceUserApprovalReqUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 09:00:00"
	EndTime          string `json:"end_time,omitempty"`           // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 19:00:00"
	Reason           string `json:"reason,omitempty"`             // 出差理由, 示例值: "培训"
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 12:00:00"
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间, 时间格式为 yyyy-MM-dd HH:mm:ss, 示例值: "2021-01-04 11:00:00"
}

// CreateAttendanceUserApprovalResp ...
type CreateAttendanceUserApprovalResp struct {
	UserApproval *CreateAttendanceUserApprovalRespUserApproval `json:"user_approval,omitempty"` // 审批信息
}

// CreateAttendanceUserApprovalRespUserApproval ...
type CreateAttendanceUserApprovalRespUserApproval struct {
	UserID        string                                                      `json:"user_id,omitempty"`        // 审批用户 ID
	Date          string                                                      `json:"date,omitempty"`           // 审批作用日期
	Outs          []*CreateAttendanceUserApprovalRespUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateAttendanceUserApprovalRespUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateAttendanceUserApprovalRespUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateAttendanceUserApprovalRespUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

// CreateAttendanceUserApprovalRespUserApprovalLeave ...
type CreateAttendanceUserApprovalRespUserApprovalLeave struct {
	ApprovalID       string     `json:"approval_id,omitempty"`        // 审批实例 ID
	UniqID           string     `json:"uniq_id,omitempty"`            // 假期类型唯一 ID, 代表一种假期类型, 长度小于 14, * 此ID对应假期类型(即: i18n_names), 因此需要保证唯一
	Unit             int64      `json:"unit,omitempty"`               // 假期时长单位, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Interval         int64      `json:"interval,omitempty"`           // 假期时长（单位: 秒）, 暂未开放提供, 待后续提供
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 假期多语言展示, 格式为 map, key 为 ["ch"、"en"、"ja"], 其中 ch 代表中文、en 代表英语、ja 代表日语
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型, 由于飞书客户端支持中、英、日三种语言, 当用户切换语言时, 如果假期名称没有所对应的语言, 会使用默认语言的名称, 可选值有: ch: 中文, en: 英文, ja: 日文
	Reason           string     `json:"reason,omitempty"`             // 请假理由, 必选字段
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间, 时间格式为 yyyy-MM-dd HH:mm:ss
}

// CreateAttendanceUserApprovalRespUserApprovalOut ...
type CreateAttendanceUserApprovalRespUserApprovalOut struct {
	ApprovalID       string     `json:"approval_id,omitempty"`        // 审批实例 ID
	UniqID           string     `json:"uniq_id,omitempty"`            // 外出类型唯一 ID, 代表一种假期类型, 长度小于 14, * 此ID对应假期类型(即: i18n_names), 因此需要保证唯一
	Unit             int64      `json:"unit,omitempty"`               // 外出时长单位, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Interval         int64      `json:"interval,omitempty"`           // 外出时长（单位: 秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 外出多语言展示, 格式为 map, key 为 ["ch"、"en"、"ja"], 其中 ch 代表中文、en 代表英语、ja 代表日语
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型, 由于飞书客户端支持中、英、日三种语言, 当用户切换语言时, 如果假期名称没有所对应的语言, 会使用默认语言的名称
	Reason           string     `json:"reason,omitempty"`             // 外出理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间
}

// CreateAttendanceUserApprovalRespUserApprovalOvertimeWork ...
type CreateAttendanceUserApprovalRespUserApprovalOvertimeWork struct {
	ApprovalID string  `json:"approval_id,omitempty"` // 审批实例 ID
	Duration   float64 `json:"duration,omitempty"`    // 加班时长
	Unit       int64   `json:"unit,omitempty"`        // 加班时长单位, 可选值有: 1: 天, 2: 小时, 3: 半天, 4: 半小时
	Category   int64   `json:"category,omitempty"`    // 加班日期类型, 可选值有: 1: 工作日, 2: 休息日, 3: 节假日
	Type       int64   `json:"type,omitempty"`        // 加班规则类型, 可选值有: 0: 不关联加班规则, 1: 调休, 2: 加班费, 3: 关联加班规则, 没有调休或加班费
	StartTime  string  `json:"start_time,omitempty"`  // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime    string  `json:"end_time,omitempty"`    // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss
}

// CreateAttendanceUserApprovalRespUserApprovalTrip ...
type CreateAttendanceUserApprovalRespUserApprovalTrip struct {
	ApprovalID       string `json:"approval_id,omitempty"`        // 审批实例 ID
	StartTime        string `json:"start_time,omitempty"`         // 开始时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间, 时间格式为 yyyy-MM-dd HH:mm:ss
}

// createAttendanceUserApprovalResp ...
type createAttendanceUserApprovalResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *CreateAttendanceUserApprovalResp `json:"data,omitempty"`
}
