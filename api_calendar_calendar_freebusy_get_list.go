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

// GetCalendarFreeBusyList 查询用户主日历或会议室的忙闲信息。
//
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/freebusy/list
func (r *CalendarService) GetCalendarFreeBusyList(ctx context.Context, request *GetCalendarFreeBusyListReq, options ...MethodOptionFunc) (*GetCalendarFreeBusyListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarFreeBusyList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarFreeBusyList mock enable")
		return r.cli.mock.mockCalendarGetCalendarFreeBusyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarFreeBusyList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/freebusy/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCalendarFreeBusyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarFreeBusyList mock CalendarGetCalendarFreeBusyList method
func (r *Mock) MockCalendarGetCalendarFreeBusyList(f func(ctx context.Context, request *GetCalendarFreeBusyListReq, options ...MethodOptionFunc) (*GetCalendarFreeBusyListResp, *Response, error)) {
	r.mockCalendarGetCalendarFreeBusyList = f
}

// UnMockCalendarGetCalendarFreeBusyList un-mock CalendarGetCalendarFreeBusyList method
func (r *Mock) UnMockCalendarGetCalendarFreeBusyList() {
	r.mockCalendarGetCalendarFreeBusyList = nil
}

// GetCalendarFreeBusyListReq ...
type GetCalendarFreeBusyListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	TimeMin    string  `json:"time_min,omitempty"`     // 查询时段开始时间, 需要url编码, 示例值: "2020-10-28T12:00:00+08:00"
	TimeMax    string  `json:"time_max,omitempty"`     // 查询时段结束时间, 需要url编码, 示例值: "2020-12-28T12:00:00+08:00"
	UserID     *string `json:"user_id,omitempty"`      // 用户user_id, 输入时与 room_id 二选一。参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "ou_xxxxxxxxxx"
	RoomID     *string `json:"room_id,omitempty"`      // 会议室room_id, 输入时与 user_id 二选一, 示例值: "omm_xxxxxxxxxx"
}

// GetCalendarFreeBusyListResp ...
type GetCalendarFreeBusyListResp struct {
	FreebusyList []*GetCalendarFreeBusyListRespFreebusy `json:"freebusy_list,omitempty"` // 日历上请求时间区间内的忙碌时间段信息。
}

// GetCalendarFreeBusyListRespFreebusy ...
type GetCalendarFreeBusyListRespFreebusy struct {
	StartTime string `json:"start_time,omitempty"` // 忙闲信息开始时间, RFC3339 date_time 格式
	EndTime   string `json:"end_time,omitempty"`   // 忙闲信息结束时间, RFC3339 date_time 格式
}

// getCalendarFreeBusyListResp ...
type getCalendarFreeBusyListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarFreeBusyListResp `json:"data,omitempty"`
}
