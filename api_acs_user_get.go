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

// GetACSUser 该接口用于获取智能门禁中单个用户的信息。
//
// 只能获取已加入智能门禁权限组的用户
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/get
func (r *ACSService) GetACSUser(ctx context.Context, request *GetACSUserReq, options ...MethodOptionFunc) (*GetACSUserResp, *Response, error) {
	if r.cli.mock.mockACSGetACSUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] ACS#GetACSUser mock enable")
		return r.cli.mock.mockACSGetACSUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "GetACSUser",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/acs/v1/users/:user_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getACSUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSGetACSUser mock ACSGetACSUser method
func (r *Mock) MockACSGetACSUser(f func(ctx context.Context, request *GetACSUserReq, options ...MethodOptionFunc) (*GetACSUserResp, *Response, error)) {
	r.mockACSGetACSUser = f
}

// UnMockACSGetACSUser un-mock ACSGetACSUser method
func (r *Mock) UnMockACSGetACSUser() {
	r.mockACSGetACSUser = nil
}

// GetACSUserReq ...
type GetACSUserReq struct {
	UserID     string  `path:"user_id" json:"-"`       // 用户 ID, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetACSUserResp ...
type GetACSUserResp struct {
	User *GetACSUserRespUser `json:"user,omitempty"` // 门禁用户信息
}

// GetACSUserRespUser ...
type GetACSUserRespUser struct {
	Feature *GetACSUserRespUserFeature `json:"feature,omitempty"` // 用户特征
	UserID  string                     `json:"user_id,omitempty"` // 用户 ID
}

// GetACSUserRespUserFeature ...
type GetACSUserRespUserFeature struct {
	Card         int64 `json:"card,omitempty"`          // 卡号
	FaceUploaded bool  `json:"face_uploaded,omitempty"` // 是否已上传人脸图片
}

// getACSUserResp ...
type getACSUserResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetACSUserResp `json:"data,omitempty"`
}
