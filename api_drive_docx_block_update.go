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

// UpdateDocxBlock 更新指定的块。
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnby5Y0yoACL3PdfZqrJEm6f#doxcnEeyS0I8MMqoieIMpK7jm8g), 了解相关规则及约束。
// 应用频率限制: 单个应用调用频率上限为每秒 3 次, 超过该频率限制, 接口将返回 HTTP 状态码 400 及错误码 99991400；
// 文档频率限制: 单篇文档并发编辑上限为每秒 3 次, 超过该频率限制, 接口将返回 HTTP 状态码 429, 编辑操作包括:
// - 创建块
// - 删除块
// - 更新块
// - 批量更新块
// 当请求被限频, 应用需要处理限频状态码, 并使用指数退避算法或其它一些频控策略降低对 API 的调用速率。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/patch
func (r *DriveService) UpdateDocxBlock(ctx context.Context, request *UpdateDocxBlockReq, options ...MethodOptionFunc) (*UpdateDocxBlockResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDocxBlock != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDocxBlock mock enable")
		return r.cli.mock.mockDriveUpdateDocxBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDocxBlock",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/:block_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDocxBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDocxBlock mock DriveUpdateDocxBlock method
func (r *Mock) MockDriveUpdateDocxBlock(f func(ctx context.Context, request *UpdateDocxBlockReq, options ...MethodOptionFunc) (*UpdateDocxBlockResp, *Response, error)) {
	r.mockDriveUpdateDocxBlock = f
}

// UnMockDriveUpdateDocxBlock un-mock DriveUpdateDocxBlock method
func (r *Mock) UnMockDriveUpdateDocxBlock() {
	r.mockDriveUpdateDocxBlock = nil
}

// UpdateDocxBlockReq ...
type UpdateDocxBlockReq struct {
	DocumentID                 string                                        `path:"document_id" json:"-"`                     // 文档的唯一标识。对应新版文档 Token, [点击了解如何获取云文档 Token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doxcnePuYufKa49ISjhD8Ih0ikh"
	BlockID                    string                                        `path:"block_id" json:"-"`                        // Block 的唯一标识, 示例值: "doxcnO6UW6wAw2qIcYf4hZpFIth"
	DocumentRevisionID         *int64                                        `query:"document_revision_id" json:"-"`           // 操作的文档版本, 1 表示文档最新版本。若此时操作的版本为文档最新版本, 则需要持有文档的阅读权限；若此时操作的版本为文档的历史版本, 则需要持有文档的编辑权限, 示例值:1, 默认值: `-1`, 最小值: `-1`
	ClientToken                *string                                       `query:"client_token" json:"-"`                   // 操作的唯一标识, 与接口返回值的 client_token 相对应, 用于幂等的进行更新操作。此值为空表示将发起一次新的请求, 此值非空表示幂等的进行更新操作, 示例值: "0e2633a3-aa1a-4171-af9e-0768ff863566"
	UserIDType                 *IDType                                       `query:"user_id_type" json:"-"`                   // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UpdateTextElements         *UpdateDocxBlockReqUpdateTextElements         `json:"update_text_elements,omitempty"`           // 更新文本元素请求
	UpdateTextStyle            *UpdateDocxBlockReqUpdateTextStyle            `json:"update_text_style,omitempty"`              // 更新文本样式请求
	UpdateTableProperty        *UpdateDocxBlockReqUpdateTableProperty        `json:"update_table_property,omitempty"`          // 更新表格属性请求
	InsertTableRow             *UpdateDocxBlockReqInsertTableRow             `json:"insert_table_row,omitempty"`               // 表格插入新行请求
	InsertTableColumn          *UpdateDocxBlockReqInsertTableColumn          `json:"insert_table_column,omitempty"`            // 表格插入新列请求
	DeleteTableRows            *UpdateDocxBlockReqDeleteTableRows            `json:"delete_table_rows,omitempty"`              // 表格批量删除行请求
	DeleteTableColumns         *UpdateDocxBlockReqDeleteTableColumns         `json:"delete_table_columns,omitempty"`           // 表格批量删除列请求
	MergeTableCells            *UpdateDocxBlockReqMergeTableCells            `json:"merge_table_cells,omitempty"`              // 表格合并单元格请求
	UnmergeTableCells          *UpdateDocxBlockReqUnmergeTableCells          `json:"unmerge_table_cells,omitempty"`            // 表格取消单元格合并状态请求
	InsertGridColumn           *UpdateDocxBlockReqInsertGridColumn           `json:"insert_grid_column,omitempty"`             // 分栏插入新的分栏列请求
	DeleteGridColumn           *UpdateDocxBlockReqDeleteGridColumn           `json:"delete_grid_column,omitempty"`             // 分栏删除列请求
	UpdateGridColumnWidthRatio *UpdateDocxBlockReqUpdateGridColumnWidthRatio `json:"update_grid_column_width_ratio,omitempty"` // 更新分栏列宽比例请求
	ReplaceImage               *UpdateDocxBlockReqReplaceImage               `json:"replace_image,omitempty"`                  // 替换图片请求
	ReplaceFile                *UpdateDocxBlockReqReplaceFile                `json:"replace_file,omitempty"`                   // 替换附件请求
	UpdateText                 *UpdateDocxBlockReqUpdateText                 `json:"update_text,omitempty"`                    // 更新文本元素及样式请求
}

// UpdateDocxBlockReqDeleteGridColumn ...
type UpdateDocxBlockReqDeleteGridColumn struct {
	ColumnIndex int64 `json:"column_index,omitempty"` // 删除列索引, 从 0 开始, 如 0 表示删除第一列（-1表示删除最后一列）, 示例值: 0, 最小值: `-1`
}

// UpdateDocxBlockReqDeleteTableColumns ...
type UpdateDocxBlockReqDeleteTableColumns struct {
	ColumnStartIndex int64 `json:"column_start_index,omitempty"` // 列开始索引（区间左闭右开）, 示例值: 0, 最小值: `0`
	ColumnEndIndex   int64 `json:"column_end_index,omitempty"`   // 列结束索引（区间左闭右开）, 示例值: 1, 最小值: `1`
}

// UpdateDocxBlockReqDeleteTableRows ...
type UpdateDocxBlockReqDeleteTableRows struct {
	RowStartIndex int64 `json:"row_start_index,omitempty"` // 行开始索引（区间左闭右开）, 示例值: 0, 最小值: `0`
	RowEndIndex   int64 `json:"row_end_index,omitempty"`   // 行结束索引（区间左闭右开）, 示例值: 1, 最小值: `1`
}

// UpdateDocxBlockReqInsertGridColumn ...
type UpdateDocxBlockReqInsertGridColumn struct {
	ColumnIndex int64 `json:"column_index,omitempty"` // 插入列索引, 从 1 开始, 如 1 表示在第一列后插入, 注意不允许传 0（-1表示在最后一列后插入）, 示例值: 1, 最小值: `-1`
}

// UpdateDocxBlockReqInsertTableColumn ...
type UpdateDocxBlockReqInsertTableColumn struct {
	ColumnIndex int64 `json:"column_index,omitempty"` // 插入的列在表格中的索引。（-1表示在表格末尾插入一列）, 示例值:1, 最小值: `-1`
}

// UpdateDocxBlockReqInsertTableRow ...
type UpdateDocxBlockReqInsertTableRow struct {
	RowIndex int64 `json:"row_index,omitempty"` // 插入的行在表格中的索引。（-1表示在表格末尾插入一行）, 示例值:1, 最小值: `-1`
}

// UpdateDocxBlockReqMergeTableCells ...
type UpdateDocxBlockReqMergeTableCells struct {
	RowStartIndex    int64 `json:"row_start_index,omitempty"`    // 行起始索引（区间左闭右开）, 示例值: 0, 最小值: `0`
	RowEndIndex      int64 `json:"row_end_index,omitempty"`      // 行结束索引（区间左闭右开）, 示例值: 1, 最小值: `1`
	ColumnStartIndex int64 `json:"column_start_index,omitempty"` // 列起始索引（区间左闭右开）, 示例值: 0, 最小值: `0`
	ColumnEndIndex   int64 `json:"column_end_index,omitempty"`   // 列结束索引（区间左闭右开）, 示例值: 1, 最小值: `1`
}

// UpdateDocxBlockReqReplaceFile ...
type UpdateDocxBlockReqReplaceFile struct {
	Token string `json:"token,omitempty"` // 附件 token, 示例值: "boxbckbfvfcqEg22hAzN8Dh9gJd"
}

// UpdateDocxBlockReqReplaceImage ...
type UpdateDocxBlockReqReplaceImage struct {
	Token string `json:"token,omitempty"` // 图片 token, 示例值: "boxbckbfvfcqEg22hAzN8Dh9gJd"
}

// UpdateDocxBlockReqUnmergeTableCells ...
type UpdateDocxBlockReqUnmergeTableCells struct {
	RowIndex    int64 `json:"row_index,omitempty"`    // table 行索引, 示例值: 0, 最小值: `0`
	ColumnIndex int64 `json:"column_index,omitempty"` // table 列索引, 示例值: 0, 最小值: `0`
}

// UpdateDocxBlockReqUpdateGridColumnWidthRatio ...
type UpdateDocxBlockReqUpdateGridColumnWidthRatio struct {
	WidthRatios []int64 `json:"width_ratios,omitempty"` // 更新列宽比例时, 需要传入所有列宽占比, 示例值: 50, 长度范围: `1` ～ `99`
}

// UpdateDocxBlockReqUpdateTableProperty ...
type UpdateDocxBlockReqUpdateTableProperty struct {
	ColumnWidth int64 `json:"column_width,omitempty"` // 表格列宽, 示例值: 100, 最小值: `50`
	ColumnIndex int64 `json:"column_index,omitempty"` // 需要修改列宽的表格列的索引, 示例值: 0, 最小值: `0`
}

// UpdateDocxBlockReqUpdateText ...
type UpdateDocxBlockReqUpdateText struct {
	Elements []*DocxTextElement `json:"elements,omitempty"` // 更新的文本元素列表, 单次更新中 reminder 上限 30 个, mention_doc 上限 50 个, mention_user 上限 100 个, 最小长度: `1`
	Style    *DocxTextStyle     `json:"style,omitempty"`    // 更新的文本样式
	Fields   []int64            `json:"fields,omitempty"`   // 文本样式中应更新的字段, 必须至少指定一个字段。例如, 要调整 Block 对齐方式, 请设置 fields 为 [1], 示例值: [1], 可选值有: 1: 修改 Block 的对齐方式, 2: 修改 todo block 的完成状态, 3: 修改 block 的折叠状态, 4: 修改代码块的语言类型, 5: 修改代码块的折叠状态
}

// UpdateDocxBlockReqUpdateTextElements ...
type UpdateDocxBlockReqUpdateTextElements struct {
	Elements []*DocxTextElement `json:"elements,omitempty"` // 更新的文本元素列表, 单次更新中 reminder 上限 30 个, mention_doc 上限 50 个, mention_user 上限 100 个, 最小长度: `1`
}

// UpdateDocxBlockReqUpdateTextStyle ...
type UpdateDocxBlockReqUpdateTextStyle struct {
	Style  *UpdateDocxBlockReqUpdateTextStyleStyle `json:"style,omitempty"`  // 文本样式
	Fields []int64                                 `json:"fields,omitempty"` // 应更新的字段, 必须至少指定一个字段。例如, 要调整 Block 对齐方式, 请设置 fields 为 [1], 示例值: 修改的文字样式属性, 可选值有: 1: 修改 Block 的对齐方式, 2: 修改 todo block 的完成状态, 3: 修改 block 的折叠状态, 4: 修改代码块的语言类型, 5: 修改代码块的折叠状态
}

// UpdateDocxBlockReqUpdateTextStyleStyle ...
type UpdateDocxBlockReqUpdateTextStyleStyle struct {
	Align    *int64 `json:"align,omitempty"`    // 对齐方式, 示例值: 1, 可选值有: 1: 居左排版, 2: 居中排版, 3: 居右排版, 默认值: `1`
	Done     *bool  `json:"done,omitempty"`     // todo 的完成状态, 示例值: true, 默认值: `false`
	Folded   *bool  `json:"folded,omitempty"`   // 文本的折叠状态, 示例值: true, 默认值: `false`
	Language *int64 `json:"language,omitempty"` // 代码块语言, 示例值: 1, 可选值有: 1: PlainText, 2: ABAP, 3: Ada, 4: Apache, 5: Apex, 6: Assembly Language, 7: Bash, 8: CSharp, 9: C++, 10: C, 11: COBOL, 12: CSS, 13: CoffeeScript, 14: D, 15: Dart, 16: Delphi, 17: Django, 18: Dockerfile, 19: Erlang, 20: Fortran, 21: FoxPro, 22: Go, 23: Groovy, 24: HTML, 25: HTMLBars, 26: HTTP, 27: Haskell, 28: JSON, 29: Java, 30: JavaScript, 31: Julia, 32: Kotlin, 33: LateX, 34: Lisp, 35: Logo, 36: Lua, 37: MATLAB, 38: Makefile, 39: Markdown, 40: Nginx, 41: Objective-C, 42: OpenEdgeABL, 43: PHP, 44: Perl, 45: PostScript, 46: Power Shell, 47: Prolog, 48: ProtoBuf, 49: Python, 50: R, 51: RPG, 52: Ruby, 53: Rust, 54: SAS, 55: SCSS, 56: SQL, 57: Scala, 58: Scheme, 59: Scratch, 60: Shell, 61: Swift, 62: Thrift, 63: TypeScript, 64: VBScript, 65: Visual Basic, 66: XML, 67: YAML, 68: CMake, 69: Diff, 70: Gherkin, 71: GraphQL, 72: OpenGL Shading Language, 73: Properties, 74: Solidity, 75: TOML
	Wrap     *bool  `json:"wrap,omitempty"`     // 代码块是否自动换行, 示例值: true, 默认值: `false`
}

// UpdateDocxBlockResp ...
type UpdateDocxBlockResp struct {
	Block              *DocxBlock `json:"block,omitempty"`                // 更新后的 block 信息
	DocumentRevisionID int64      `json:"document_revision_id,omitempty"` // 当前更新成功后文档的版本号
	ClientToken        string     `json:"client_token,omitempty"`         // 操作的唯一标识, 更新请求中使用此值表示幂等的进行此次更新
}

// updateDocxBlockResp ...
type updateDocxBlockResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDocxBlockResp `json:"data,omitempty"`
}
