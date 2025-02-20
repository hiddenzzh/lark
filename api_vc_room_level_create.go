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

// CreateVCRoomLevel 该接口用于创建会议室层级。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/create
func (r *VCService) CreateVCRoomLevel(ctx context.Context, request *CreateVCRoomLevelReq, options ...MethodOptionFunc) (*CreateVCRoomLevelResp, *Response, error) {
	if r.cli.mock.mockVCCreateVCRoomLevel != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#CreateVCRoomLevel mock enable")
		return r.cli.mock.mockVCCreateVCRoomLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "CreateVCRoomLevel",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/room_levels",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createVCRoomLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCCreateVCRoomLevel mock VCCreateVCRoomLevel method
func (r *Mock) MockVCCreateVCRoomLevel(f func(ctx context.Context, request *CreateVCRoomLevelReq, options ...MethodOptionFunc) (*CreateVCRoomLevelResp, *Response, error)) {
	r.mockVCCreateVCRoomLevel = f
}

// UnMockVCCreateVCRoomLevel un-mock VCCreateVCRoomLevel method
func (r *Mock) UnMockVCCreateVCRoomLevel() {
	r.mockVCCreateVCRoomLevel = nil
}

// CreateVCRoomLevelReq ...
type CreateVCRoomLevelReq struct {
	Name          string  `json:"name,omitempty"`            // 层级名称, 示例值: "测试层级"
	ParentID      string  `json:"parent_id,omitempty"`       // 父层级ID, 示例值: "omb_4ad1a2c7a2fbc5fc9570f38456931293"
	CustomGroupID *string `json:"custom_group_id,omitempty"` // 自定义层级ID, 示例值: "10000"
}

// CreateVCRoomLevelResp ...
type CreateVCRoomLevelResp struct {
	RoomLevel *CreateVCRoomLevelRespRoomLevel `json:"room_level,omitempty"` // 层级详情
}

// CreateVCRoomLevelRespRoomLevel ...
type CreateVCRoomLevelRespRoomLevel struct {
	RoomLevelID   string   `json:"room_level_id,omitempty"`   // 层级ID
	Name          string   `json:"name,omitempty"`            // 层级名称
	ParentID      string   `json:"parent_id,omitempty"`       // 父层级ID
	Path          []string `json:"path,omitempty"`            // 层级路径
	HasChild      bool     `json:"has_child,omitempty"`       // 是否有子层级
	CustomGroupID string   `json:"custom_group_id,omitempty"` // 自定义层级ID
}

// createVCRoomLevelResp ...
type createVCRoomLevelResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateVCRoomLevelResp `json:"data,omitempty"`
}
