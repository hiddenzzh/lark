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

// GetDocxDocumentRawContent 获取文档的纯文本内容。
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnby5Y0yoACL3PdfZqrJEm6f#doxcnQeqI4wiKIMis6GNvCOBuqg), 了解相关规则及约束。
// 应用频率限制: 单个应用调用频率上限为每秒 5 次, 超过该频率限制, 接口将返回 HTTP 状态码 400 及错误码 99991400。当请求被限频, 应用需要处理限频状态码, 并使用指数退避算法或其它一些频控策略降低对 API 的调用速率。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/raw_content
func (r *DriveService) GetDocxDocumentRawContent(ctx context.Context, request *GetDocxDocumentRawContentReq, options ...MethodOptionFunc) (*GetDocxDocumentRawContentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDocxDocumentRawContent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDocxDocumentRawContent mock enable")
		return r.cli.mock.mockDriveGetDocxDocumentRawContent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDocxDocumentRawContent",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/raw_content",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDocxDocumentRawContentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDocxDocumentRawContent mock DriveGetDocxDocumentRawContent method
func (r *Mock) MockDriveGetDocxDocumentRawContent(f func(ctx context.Context, request *GetDocxDocumentRawContentReq, options ...MethodOptionFunc) (*GetDocxDocumentRawContentResp, *Response, error)) {
	r.mockDriveGetDocxDocumentRawContent = f
}

// UnMockDriveGetDocxDocumentRawContent un-mock DriveGetDocxDocumentRawContent method
func (r *Mock) UnMockDriveGetDocxDocumentRawContent() {
	r.mockDriveGetDocxDocumentRawContent = nil
}

// GetDocxDocumentRawContentReq ...
type GetDocxDocumentRawContentReq struct {
	DocumentID string `path:"document_id" json:"-"` // 文档的唯一标识。对应新版文档 Token, [点击了解如何获取云文档 Token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doxbcmEtbFrbbq10nPNu8gO1F3b", 长度范围: `27` ～ `27` 字符
	Lang       *int64 `query:"lang" json:"-"`       // 语言（用于 MentionUser 语言的选取）, 示例值: 0, 可选值有: 0: 中文, 1: 英文, 2: 日文, 默认值: `0`
}

// GetDocxDocumentRawContentResp ...
type GetDocxDocumentRawContentResp struct {
	Content string `json:"content,omitempty"` // 文档纯文本
}

// getDocxDocumentRawContentResp ...
type getDocxDocumentRawContentResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetDocxDocumentRawContentResp `json:"data,omitempty"`
}
