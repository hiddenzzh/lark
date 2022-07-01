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

// ImportSheet >  为了更好地提升该接口的安全性, 我们对其进行了升级, 请尽快迁移至[新版本](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide)
//
// 该接口用于将本地表格导入到云空间上。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATO2YjLwkjN24CM5YjN
//
// Deprecated
func (r *DriveService) ImportSheet(ctx context.Context, request *ImportSheetReq, options ...MethodOptionFunc) (*ImportSheetResp, *Response, error) {
	if r.cli.mock.mockDriveImportSheet != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#ImportSheet mock enable")
		return r.cli.mock.mockDriveImportSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "ImportSheet",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/import",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(importSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveImportSheet mock DriveImportSheet method
func (r *Mock) MockDriveImportSheet(f func(ctx context.Context, request *ImportSheetReq, options ...MethodOptionFunc) (*ImportSheetResp, *Response, error)) {
	r.mockDriveImportSheet = f
}

// UnMockDriveImportSheet un-mock DriveImportSheet method
func (r *Mock) UnMockDriveImportSheet() {
	r.mockDriveImportSheet = nil
}

// ImportSheetReq ...
type ImportSheetReq struct {
	File        []byte  `json:"file,omitempty"`        // 需要导入的文件数据, 转换成字节数组的形式, 支持"xlsx", "csv"格式, 最大不超过20M
	Name        string  `json:"name,omitempty"`        // 文件名, 带上文件拓展名, 如"hello.csv"、"hello.xlsx"。导入后sheet的标题将去除文件拓展名, 如"hello.xlsx"导入后标题为"hello"。
	FolderToken *string `json:"folderToken,omitempty"` // 导入的文件夹token, 默认导入到根目录下
}

// ImportSheetResp ...
type ImportSheetResp struct {
	Ticket string `json:"ticket,omitempty"` // 与导入文件一一对应的凭证, 用于查询文件导入的进度, 详见[查询导入结果的接口](https://open.feishu.cn/document/ukTMukTMukTM/uETO2YjLxkjN24SM5YjN)
}

// importSheetResp ...
type importSheetResp struct {
	Code int64            `json:"code,omitempty"`
	Msg  string           `json:"msg,omitempty"`
	Data *ImportSheetResp `json:"data,omitempty"`
}
