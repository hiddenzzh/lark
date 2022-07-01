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

// BatchGetMeetingRoomFreebusy 该接口用于获取指定会议室的忙闲日程实例列表。非重复日程只有唯一实例；重复日程可能存在多个实例, 依据重复规则和时间范围扩展。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDOyUjLygjM14iM4ITN
func (r *MeetingRoomService) BatchGetMeetingRoomFreebusy(ctx context.Context, request *BatchGetMeetingRoomFreebusyReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomFreebusyResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomFreebusy != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomFreebusy mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomFreebusy(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomFreebusy",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/freebusy/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomFreebusyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMeetingRoomBatchGetMeetingRoomFreebusy mock MeetingRoomBatchGetMeetingRoomFreebusy method
func (r *Mock) MockMeetingRoomBatchGetMeetingRoomFreebusy(f func(ctx context.Context, request *BatchGetMeetingRoomFreebusyReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomFreebusyResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomFreebusy = f
}

// UnMockMeetingRoomBatchGetMeetingRoomFreebusy un-mock MeetingRoomBatchGetMeetingRoomFreebusy method
func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomFreebusy() {
	r.mockMeetingRoomBatchGetMeetingRoomFreebusy = nil
}

// BatchGetMeetingRoomFreebusyReq ...
type BatchGetMeetingRoomFreebusyReq struct {
	RoomIDs []string `query:"room_ids" json:"-"` // 用于查询指定会议室的 ID
	TimeMin string   `query:"time_min" json:"-"` // 查询会议室忙闲的起始时间, 需要遵循格式 [RFC3339](https://tools.ietf.org/html/rfc3339), 需要进行URL Encode
	TimeMax string   `query:"time_max" json:"-"` // 查询会议室忙闲的结束时间, 需要遵循格式 [RFC3339](https://tools.ietf.org/html/rfc3339), 需要进行URL Encode
}

// BatchGetMeetingRoomFreebusyResp ...
type BatchGetMeetingRoomFreebusyResp struct {
	TimeMin  string                                              `json:"time_min,omitempty"`  // 查询会议室忙闲的起始时间, 与请求参数完全相同
	TimeMax  string                                              `json:"time_max,omitempty"`  // 查询会议室忙闲的结束时间, 与请求参数完全相同
	FreeBusy map[string]*BatchGetMeetingRoomFreebusyRespFreeBusy `json:"free_busy,omitempty"` // 会议室忙闲列表
}

// BatchGetMeetingRoomFreebusyRespFreeBusy ...
type BatchGetMeetingRoomFreebusyRespFreeBusy struct {
	StartTime     string                                                `json:"start_time,omitempty"`     // 忙碌起始时间
	EndTime       string                                                `json:"end_time,omitempty"`       // 忙碌结束时间
	Uid           string                                                `json:"uid,omitempty"`            // 日程 ID
	OriginalTime  int64                                                 `json:"original_time,omitempty"`  // 日程实例的原始时间, 非重复日程以及重复性日程的原日程为0, 重复性日程的例外日程为非0
	OrganizerInfo *BatchGetMeetingRoomFreebusyRespFreeBusyOrganizerInfo `json:"organizer_info,omitempty"` // 组织者信息, 私密日程不返回该信息
}

// BatchGetMeetingRoomFreebusyRespFreeBusyOrganizerInfo ...
type BatchGetMeetingRoomFreebusyRespFreeBusyOrganizerInfo struct {
	Name   string `json:"name,omitempty"`    // 组织者姓名
	OpenID string `json:"open_id,omitempty"` // 组织者 open_id
}

// batchGetMeetingRoomFreebusyResp ...
type batchGetMeetingRoomFreebusyResp struct {
	Code int64                            `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 返回码的描述, "success" 表示成功, 其他为错误提示信息
	Data *BatchGetMeetingRoomFreebusyResp `json:"data,omitempty"` // 返回业务信息
}
