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

// GetDriveFolderMeta 该接口用于根据 folderToken 获取该文件夹的元信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN
func (r *DriveService) GetDriveFolderMeta(ctx context.Context, request *GetDriveFolderMetaReq, options ...MethodOptionFunc) (*GetDriveFolderMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFolderMeta != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFolderMeta mock enable")
		return r.cli.mock.mockDriveGetDriveFolderMeta(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveFolderMeta",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/explorer/v2/folder/:folderToken/meta",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveFolderMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveFolderMeta mock DriveGetDriveFolderMeta method
func (r *Mock) MockDriveGetDriveFolderMeta(f func(ctx context.Context, request *GetDriveFolderMetaReq, options ...MethodOptionFunc) (*GetDriveFolderMetaResp, *Response, error)) {
	r.mockDriveGetDriveFolderMeta = f
}

// UnMockDriveGetDriveFolderMeta un-mock DriveGetDriveFolderMeta method
func (r *Mock) UnMockDriveGetDriveFolderMeta() {
	r.mockDriveGetDriveFolderMeta = nil
}

// GetDriveFolderMetaReq ...
type GetDriveFolderMetaReq struct {
	FolderToken string `path:"folderToken" json:"-"` // 文件夹 token, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)
}

// GetDriveFolderMetaResp ...
type GetDriveFolderMetaResp struct {
	ID        string `json:"id,omitempty"`        // 文件夹的 id
	Name      string `json:"name,omitempty"`      // 文件夹的标题
	Token     string `json:"token,omitempty"`     // 文件夹的 token
	CreateUid string `json:"createUid,omitempty"` // 文件夹的创建者 id
	EditUid   string `json:"editUid,omitempty"`   // 文件夹的最后编辑者 id
	ParentID  string `json:"parentId,omitempty"`  // 文件夹的上级目录 id
	OwnUid    string `json:"ownUid,omitempty"`    // 文件夹为个人文件夹时, 为文件夹的所有者 id；文件夹为共享文件夹时, 为文件夹树id
}

// getDriveFolderMetaResp ...
type getDriveFolderMetaResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *GetDriveFolderMetaResp `json:"data,omitempty"`
}
