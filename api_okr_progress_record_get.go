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

// GetOKRProgressRecord 根据 ID 获取 OKR 进展记录详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/progress_record/get
func (r *OKRService) GetOKRProgressRecord(ctx context.Context, request *GetOKRProgressRecordReq, options ...MethodOptionFunc) (*GetOKRProgressRecordResp, *Response, error) {
	if r.cli.mock.mockOKRGetOKRProgressRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetOKRProgressRecord mock enable")
		return r.cli.mock.mockOKRGetOKRProgressRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetOKRProgressRecord",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/progress_records/:progress_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getOKRProgressRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRGetOKRProgressRecord mock OKRGetOKRProgressRecord method
func (r *Mock) MockOKRGetOKRProgressRecord(f func(ctx context.Context, request *GetOKRProgressRecordReq, options ...MethodOptionFunc) (*GetOKRProgressRecordResp, *Response, error)) {
	r.mockOKRGetOKRProgressRecord = f
}

// UnMockOKRGetOKRProgressRecord un-mock OKRGetOKRProgressRecord method
func (r *Mock) UnMockOKRGetOKRProgressRecord() {
	r.mockOKRGetOKRProgressRecord = nil
}

// GetOKRProgressRecordReq ...
type GetOKRProgressRecordReq struct {
	ProgressID string  `path:"progress_id" json:"-"`   // 待查询的 OKR进展记录 ID, 示例值: "7041857032248410131"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetOKRProgressRecordResp ...
type GetOKRProgressRecordResp struct {
	ProgressID string                           `json:"progress_id,omitempty"` // OKR 进展ID
	ModifyTime string                           `json:"modify_time,omitempty"` // 进展更新时间 毫秒
	Content    *GetOKRProgressRecordRespContent `json:"content,omitempty"`     // 进展 对应的 Content 详细内容
}

// GetOKRProgressRecordRespContent ...
type GetOKRProgressRecordRespContent struct {
	Blocks []*GetOKRProgressRecordRespContentBlock `json:"blocks,omitempty"` // 文档结构是按行排列的, 每行内容是一个 Block
}

// GetOKRProgressRecordRespContentBlock ...
type GetOKRProgressRecordRespContentBlock struct {
	Type      string                                         `json:"type,omitempty"`      // 文档元素类型, 可选值有: paragraph: 文本段落, gallery: 图片
	Paragraph *GetOKRProgressRecordRespContentBlockParagraph `json:"paragraph,omitempty"` // 文本段落
	Gallery   *GetOKRProgressRecordRespContentBlockGallery   `json:"gallery,omitempty"`   // 图片
}

// GetOKRProgressRecordRespContentBlockGallery ...
type GetOKRProgressRecordRespContentBlockGallery struct {
	ImageList []*GetOKRProgressRecordRespContentBlockGalleryImageList `json:"imageList,omitempty"` // 图片元素
}

// GetOKRProgressRecordRespContentBlockGalleryImageList ...
type GetOKRProgressRecordRespContentBlockGalleryImageList struct {
	FileToken string  `json:"fileToken,omitempty"` // 图片 token, 通过上传图片接口获取
	Src       string  `json:"src,omitempty"`       // 图片链接
	Width     float64 `json:"width,omitempty"`     // 图片宽, 单位px
	Height    float64 `json:"height,omitempty"`    // 图片高, 单位px
}

// GetOKRProgressRecordRespContentBlockParagraph ...
type GetOKRProgressRecordRespContentBlockParagraph struct {
	Style    *GetOKRProgressRecordRespContentBlockParagraphStyle     `json:"style,omitempty"`    // 段落样式
	Elements []*GetOKRProgressRecordRespContentBlockParagraphElement `json:"elements,omitempty"` // 段落元素组成一个段落
}

// GetOKRProgressRecordRespContentBlockParagraphElement ...
type GetOKRProgressRecordRespContentBlockParagraphElement struct {
	Type     string                                                        `json:"type,omitempty"`     // 元素类型, 可选值有: textRun: 文本型元素, docsLink: 文档链接型元素, person: 艾特用户型元素
	TextRun  *GetOKRProgressRecordRespContentBlockParagraphElementTextRun  `json:"textRun,omitempty"`  // 文本
	DocsLink *GetOKRProgressRecordRespContentBlockParagraphElementDocsLink `json:"docsLink,omitempty"` // 飞书云文档
	Person   *GetOKRProgressRecordRespContentBlockParagraphElementPerson   `json:"person,omitempty"`   // 艾特用户
}

// GetOKRProgressRecordRespContentBlockParagraphElementDocsLink ...
type GetOKRProgressRecordRespContentBlockParagraphElementDocsLink struct {
	URL   string `json:"url,omitempty"`   // 飞书云文档链接地址
	Title string `json:"title,omitempty"` // 飞书云文档标题
}

// GetOKRProgressRecordRespContentBlockParagraphElementPerson ...
type GetOKRProgressRecordRespContentBlockParagraphElementPerson struct {
	OpenID string `json:"openId,omitempty"` // 员工的OpenID
}

// GetOKRProgressRecordRespContentBlockParagraphElementTextRun ...
type GetOKRProgressRecordRespContentBlockParagraphElementTextRun struct {
	Text  string                                                            `json:"text,omitempty"`  // 具体的文本内容
	Style *GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyle `json:"style,omitempty"` // 文本内容的样式, 支持 BIUS、颜色等
}

// GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyle ...
type GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyle struct {
	Bold          bool                                                                       `json:"bold,omitempty"`          // 是否加粗
	StrikeThrough bool                                                                       `json:"strikeThrough,omitempty"` // 是否删除
	BackColor     *GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor `json:"backColor,omitempty"`     // 背景颜色
	TextColor     *GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor `json:"textColor,omitempty"`     // 字体颜色
	Link          *GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink      `json:"link,omitempty"`          // 链接地址
}

// GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor ...
type GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor struct {
	Red   int64   `json:"red,omitempty"`   // 红 取值范围[0, 255]
	Green int64   `json:"green,omitempty"` // 绿 取值范围[0, 255]
	Blue  int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255]
	Alpha float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1]
}

// GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink ...
type GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink struct {
	URL string `json:"url,omitempty"` // 链接地址
}

// GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor ...
type GetOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor struct {
	Red   int64   `json:"red,omitempty"`   // 红 取值范围[0, 255]
	Green int64   `json:"green,omitempty"` // 绿 取值范围[0, 255]
	Blue  int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255]
	Alpha float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1]
}

// GetOKRProgressRecordRespContentBlockParagraphStyle ...
type GetOKRProgressRecordRespContentBlockParagraphStyle struct {
	List *GetOKRProgressRecordRespContentBlockParagraphStyleList `json:"list,omitempty"` // 有序列表/无序列表/任务列表
}

// GetOKRProgressRecordRespContentBlockParagraphStyleList ...
type GetOKRProgressRecordRespContentBlockParagraphStyleList struct {
	Type        string `json:"type,omitempty"`        // 列表类型, 可选值有: number: 有序列表, bullet: 无序列表, checkBox: 任务列表, checkedBox: 已完成的任务列表, indent: tab缩进
	IndentLevel int64  `json:"indentLevel,omitempty"` // 列表的缩进级别, 支持指定一行的缩进 除代码块以外的列表都支持设置缩进, 支持 1-16 级缩进, 取值范围: [1, 16]
	Number      int64  `json:"number,omitempty"`      // 用于指定列表的行号, 仅对有序列表和代码块生效 如果为有序列表设置了缩进, 行号可能会显示为字母或者罗马数字
}

// getOKRProgressRecordResp ...
type getOKRProgressRecordResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetOKRProgressRecordResp `json:"data,omitempty"`
}
