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

// GetDriveFileMeta 该接口用于根据 token 获取各类文件的元数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/meta/batch_query
func (r *DriveService) GetDriveFileMeta(ctx context.Context, request *GetDriveFileMetaReq, options ...MethodOptionFunc) (*GetDriveFileMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFileMeta != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFileMeta mock enable")
		return r.cli.mock.mockDriveGetDriveFileMeta(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveFileMeta",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/metas/batch_query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveFileMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveFileMeta mock DriveGetDriveFileMeta method
func (r *Mock) MockDriveGetDriveFileMeta(f func(ctx context.Context, request *GetDriveFileMetaReq, options ...MethodOptionFunc) (*GetDriveFileMetaResp, *Response, error)) {
	r.mockDriveGetDriveFileMeta = f
}

// UnMockDriveGetDriveFileMeta un-mock DriveGetDriveFileMeta method
func (r *Mock) UnMockDriveGetDriveFileMeta() {
	r.mockDriveGetDriveFileMeta = nil
}

// GetDriveFileMetaReq ...
type GetDriveFileMetaReq struct {
	UserIDType  *IDType                           `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	RequestDocs []*GetDriveFileMetaReqRequestDocs `json:"request_docs,omitempty"` // 请求文档, 一次不超过200个, 长度范围: `1` ～ `200`
	WithURL     *bool                             `json:"with_url,omitempty"`     // 是否获取文档链接, 示例值: false
}

// GetDriveFileMetaReqRequestDocs ...
type GetDriveFileMetaReqRequestDocs struct {
	DocToken string `json:"doc_token,omitempty"` // 文件的 token, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnfYZzTlvXqZIGTdAHKabcef"
	DocType  string `json:"doc_type,omitempty"`  // 文件类型, 示例值: "doc", 可选值有: doc: 飞书文档, sheet: 飞书电子表格, bitable: 飞书多维表格, mindnote: 飞书思维笔记, file: 飞书文件, wiki: 飞书wiki, docx: 飞书新版文档, folder: 飞书文件夹
}

// GetDriveFileMetaResp ...
type GetDriveFileMetaResp struct {
	Metas      []*GetDriveFileMetaRespMeta   `json:"metas,omitempty"`       // 文档元数据列表
	FailedList []*GetDriveFileMetaRespFailed `json:"failed_list,omitempty"` // 无法获取元数据的文档列表
}

// GetDriveFileMetaRespFailed ...
type GetDriveFileMetaRespFailed struct {
	Token string `json:"token,omitempty"` // 获取元数据失败的文档token
	Code  int64  `json:"code,omitempty"`  // 获取元数据失败的错误码, 可选值有: 970002: Unsupported doc-type, 970003: No permission to access meta, 970005: Record not found (不存在或者已被删除)
}

// GetDriveFileMetaRespMeta ...
type GetDriveFileMetaRespMeta struct {
	DocToken         string `json:"doc_token,omitempty"`          // 文件token
	DocType          string `json:"doc_type,omitempty"`           // 文件类型
	Title            string `json:"title,omitempty"`              // 标题
	OwnerID          string `json:"owner_id,omitempty"`           // 文件所有者
	CreateTime       string `json:"create_time,omitempty"`        // 创建时间（Unix时间戳）
	LatestModifyUser string `json:"latest_modify_user,omitempty"` // 最后编辑者
	LatestModifyTime string `json:"latest_modify_time,omitempty"` // 最后编辑时间（Unix时间戳）
	URL              string `json:"url,omitempty"`                // 文档链接
	SecLabelName     string `json:"sec_label_name,omitempty"`     // 文档密级标签名称, 字段权限要求: 获取文档密级标签名称
}

// getDriveFileMetaResp ...
type getDriveFileMetaResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveFileMetaResp `json:"data,omitempty"`
}
