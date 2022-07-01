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
	"io"
)

// GetACSUserFace 对于已经录入人脸图片的用户, 可以使用该接口下载用户人脸图片
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/get
func (r *ACSService) GetACSUserFace(ctx context.Context, request *GetACSUserFaceReq, options ...MethodOptionFunc) (*GetACSUserFaceResp, *Response, error) {
	if r.cli.mock.mockACSGetACSUserFace != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] ACS#GetACSUserFace mock enable")
		return r.cli.mock.mockACSGetACSUserFace(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "GetACSUserFace",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/acs/v1/users/:user_id/face",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getACSUserFaceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSGetACSUserFace mock ACSGetACSUserFace method
func (r *Mock) MockACSGetACSUserFace(f func(ctx context.Context, request *GetACSUserFaceReq, options ...MethodOptionFunc) (*GetACSUserFaceResp, *Response, error)) {
	r.mockACSGetACSUserFace = f
}

// UnMockACSGetACSUserFace un-mock ACSGetACSUserFace method
func (r *Mock) UnMockACSGetACSUserFace() {
	r.mockACSGetACSUserFace = nil
}

// GetACSUserFaceReq ...
type GetACSUserFaceReq struct {
	UserID     string  `path:"user_id" json:"-"`       // 用户 ID, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	IsCropped  *bool   `query:"is_cropped" json:"-"`   // 裁剪图, 示例值: true
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// getACSUserFaceResp ...
type getACSUserFaceResp struct {
	Code int64               `json:"code,omitempty"`
	Msg  string              `json:"msg,omitempty"`
	Data *GetACSUserFaceResp `json:"data,omitempty"`
}

func (r *getACSUserFaceResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &GetACSUserFaceResp{}
	}
	r.Data.File = file
}

// GetACSUserFaceResp ...
type GetACSUserFaceResp struct {
	File io.Reader `json:"file,omitempty"`
}
