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

// GetVCParticipantList 查询参会人明细。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/participant_list/get
func (r *VCService) GetVCParticipantList(ctx context.Context, request *GetVCParticipantListReq, options ...MethodOptionFunc) (*GetVCParticipantListResp, *Response, error) {
	if r.cli.mock.mockVCGetVCParticipantList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCParticipantList mock enable")
		return r.cli.mock.mockVCGetVCParticipantList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCParticipantList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/participant_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCParticipantListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCParticipantList mock VCGetVCParticipantList method
func (r *Mock) MockVCGetVCParticipantList(f func(ctx context.Context, request *GetVCParticipantListReq, options ...MethodOptionFunc) (*GetVCParticipantListResp, *Response, error)) {
	r.mockVCGetVCParticipantList = f
}

// UnMockVCGetVCParticipantList un-mock VCGetVCParticipantList method
func (r *Mock) UnMockVCGetVCParticipantList() {
	r.mockVCGetVCParticipantList = nil
}

// GetVCParticipantListReq ...
type GetVCParticipantListReq struct {
	MeetingStartTime string  `query:"meeting_start_time" json:"-"` // 会议开始时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingEndTime   string  `query:"meeting_end_time" json:"-"`   // 会议结束时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingNo        string  `query:"meeting_no" json:"-"`         // 9位会议号, 示例值: "123456789"
	UserID           *string `query:"user_id" json:"-"`            // 按参会Lark用户筛选（最多一个筛选条件）, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	RoomID           *string `query:"room_id" json:"-"`            // 按参会Rooms筛选（最多一个筛选条件）, 示例值: "omm_eada1d61a550955240c28757e7dec3af"
	PageSize         *int64  `query:"page_size" json:"-"`          // 分页大小, 示例值: 20, 默认值: `20`, 最大值: `100`
	PageToken        *string `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "20"
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCParticipantListResp ...
type GetVCParticipantListResp struct {
	Participants []*GetVCParticipantListRespParticipant `json:"participants,omitempty"` // 参会人列表
	PageToken    string                                 `json:"page_token,omitempty"`   // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore      bool                                   `json:"has_more,omitempty"`     // 是否还有更多项
}

// GetVCParticipantListRespParticipant ...
type GetVCParticipantListRespParticipant struct {
	ParticipantName string `json:"participant_name,omitempty"` // 参会者
	Department      string `json:"department,omitempty"`       // 部门
	UserID          string `json:"user_id,omitempty"`          // 用户ID
	EmployeeID      string `json:"employee_id,omitempty"`      // 工号
	Phone           string `json:"phone,omitempty"`            // 电话
	Email           string `json:"email,omitempty"`            // 邮箱
	Device          string `json:"device,omitempty"`           // 设备
	AppVersion      string `json:"app_version,omitempty"`      // 客户端版本
	PublicIp        string `json:"public_ip,omitempty"`        // 公网IP
	InternalIp      string `json:"internal_ip,omitempty"`      // 内网IP
	UseRtcProxy     bool   `json:"use_rtc_proxy,omitempty"`    // 代理服务
	Location        string `json:"location,omitempty"`         // 位置
	NetworkType     string `json:"network_type,omitempty"`     // 网络类型
	Protocol        string `json:"protocol,omitempty"`         // 连接类型
	Microphone      string `json:"microphone,omitempty"`       // 麦克风
	Speaker         string `json:"speaker,omitempty"`          // 扬声器
	Camera          string `json:"camera,omitempty"`           // 摄像头
	Audio           bool   `json:"audio,omitempty"`            // 音频
	Video           bool   `json:"video,omitempty"`            // 视频
	Sharing         bool   `json:"sharing,omitempty"`          // 共享
	JoinTime        string `json:"join_time,omitempty"`        // 入会时间
	LeaveTime       string `json:"leave_time,omitempty"`       // 离会时间
	TimeInMeeting   string `json:"time_in_meeting,omitempty"`  // 参会时长
	LeaveReason     string `json:"leave_reason,omitempty"`     // 离会原因
}

// getVCParticipantListResp ...
type getVCParticipantListResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetVCParticipantListResp `json:"data,omitempty"`
}
