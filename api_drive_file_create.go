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

// CreateDriveFile 在用户云空间指定文件夹中创建旧版文档、电子表格或者多维表格。如果目标文件夹是「我的空间」, 则新建的文档会在「我的空间」的「归我所有」列表里。
//
// - 云空间中文件夹单层节点上限是 1500 个, 超过此限制接口将会返回失败。如有创建大量节点的需求, 可以考虑将文档新建到不同文件夹下；
// - 我们对创建类接口进行了更细粒度的拆分和升级:
// --本接口不支持创建[新版文档](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-overview), 如需创建新版文档, 请调用[创建新版文档](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/create)接口；
// --如需创建电子表格, 也可以调用[创建表格](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create)接口。
// 该接口不支持并发创建, 且调用频率上限为 5QPS 且 10000次/天
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN
func (r *DriveService) CreateDriveFile(ctx context.Context, request *CreateDriveFileReq, options ...MethodOptionFunc) (*CreateDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveFile mock enable")
		return r.cli.mock.mockDriveCreateDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/explorer/v2/file/:folderToken",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveFile mock DriveCreateDriveFile method
func (r *Mock) MockDriveCreateDriveFile(f func(ctx context.Context, request *CreateDriveFileReq, options ...MethodOptionFunc) (*CreateDriveFileResp, *Response, error)) {
	r.mockDriveCreateDriveFile = f
}

// UnMockDriveCreateDriveFile un-mock DriveCreateDriveFile method
func (r *Mock) UnMockDriveCreateDriveFile() {
	r.mockDriveCreateDriveFile = nil
}

// CreateDriveFileReq ...
type CreateDriveFileReq struct {
	FolderToken string `path:"folderToken" json:"-"` // 文件夹 token, 用于在此文件夹下新建文档, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)
	Title       string `json:"title,omitempty"`      // 创建文档的标题。注: type 为 "doc" 时不可用（非必填, 请求会被过滤）, 有创建带标题doc文档需求可用 [创建文档](https://open.feishu.cn/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN) 接口
	Type        string `json:"type,omitempty"`       // 需要创建文档的类型  "doc" 、 "sheet"  or  "bitable"
}

// CreateDriveFileResp ...
type CreateDriveFileResp struct {
	URL      string `json:"url,omitempty"`      // 新创建文档的 url
	Token    string `json:"token,omitempty"`    // 新创建文档的 token
	Revision int64  `json:"revision,omitempty"` // 新创建文档的版本号
}

// createDriveFileResp ...
type createDriveFileResp struct {
	Code int64                `json:"code,omitempty"`
	Msg  string               `json:"msg,omitempty"`
	Data *CreateDriveFileResp `json:"data,omitempty"`
}
