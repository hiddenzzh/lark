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

// CheckDriveMemberPermission 该接口用于根据 filetoken 判断当前登录用户是否具有某权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/auth
func (r *DriveService) CheckDriveMemberPermission(ctx context.Context, request *CheckDriveMemberPermissionReq, options ...MethodOptionFunc) (*CheckDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveCheckDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CheckDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveCheckDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CheckDriveMemberPermission",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/members/auth",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(checkDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCheckDriveMemberPermission mock DriveCheckDriveMemberPermission method
func (r *Mock) MockDriveCheckDriveMemberPermission(f func(ctx context.Context, request *CheckDriveMemberPermissionReq, options ...MethodOptionFunc) (*CheckDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveCheckDriveMemberPermission = f
}

// UnMockDriveCheckDriveMemberPermission un-mock DriveCheckDriveMemberPermission method
func (r *Mock) UnMockDriveCheckDriveMemberPermission() {
	r.mockDriveCheckDriveMemberPermission = nil
}

// CheckDriveMemberPermissionReq ...
type CheckDriveMemberPermissionReq struct {
	Token  string `path:"token" json:"-"`   // 文件的 token, 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type   string `query:"type" json:"-"`   // 文件类型, 需要与文件的 token 相匹配, 示例值: "doc", 可选值有: doc: 文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙记
	Action string `query:"action" json:"-"` // 需要判断的权限, 示例值: "view", 可选值有: view: 阅读, edit: 编辑, share: 分享, comment: 评论, export: 导出, copy: 拷贝, print: 打印
}

// CheckDriveMemberPermissionResp ...
type CheckDriveMemberPermissionResp struct {
	AuthResult bool `json:"auth_result,omitempty"` // 是否有权限
}

// checkDriveMemberPermissionResp ...
type checkDriveMemberPermissionResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *CheckDriveMemberPermissionResp `json:"data,omitempty"`
}
