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

// CreateOKRProgressRecord 创建 OKR 进展记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/progress_record/create
func (r *OKRService) CreateOKRProgressRecord(ctx context.Context, request *CreateOKRProgressRecordReq, options ...MethodOptionFunc) (*CreateOKRProgressRecordResp, *Response, error) {
	if r.cli.mock.mockOKRCreateOKRProgressRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#CreateOKRProgressRecord mock enable")
		return r.cli.mock.mockOKRCreateOKRProgressRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "CreateOKRProgressRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/progress_records",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createOKRProgressRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRCreateOKRProgressRecord mock OKRCreateOKRProgressRecord method
func (r *Mock) MockOKRCreateOKRProgressRecord(f func(ctx context.Context, request *CreateOKRProgressRecordReq, options ...MethodOptionFunc) (*CreateOKRProgressRecordResp, *Response, error)) {
	r.mockOKRCreateOKRProgressRecord = f
}

// UnMockOKRCreateOKRProgressRecord un-mock OKRCreateOKRProgressRecord method
func (r *Mock) UnMockOKRCreateOKRProgressRecord() {
	r.mockOKRCreateOKRProgressRecord = nil
}

// CreateOKRProgressRecordReq ...
type CreateOKRProgressRecordReq struct {
	UserIDType  *IDType                            `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	SourceTitle string                             `json:"source_title,omitempty"` // 进展来源, 示例值: "周报系统"
	SourceURL   string                             `json:"source_url,omitempty"`   // 进展来源链接, 示例值: "https://www.zhoubao.com", 正则校验: `^https?://.*$`
	TargetID    string                             `json:"target_id,omitempty"`    // 目标id, 与target_type对应, 示例值: "7041430377642082323"
	TargetType  int64                              `json:"target_type,omitempty"`  // 目标类型, 示例值: 1, 可选值有: 2: okr的O, 3: okr的KR
	Content     *CreateOKRProgressRecordReqContent `json:"content,omitempty"`      // 进展详情 富文本格式
}

// CreateOKRProgressRecordReqContent ...
type CreateOKRProgressRecordReqContent struct {
	Blocks []*CreateOKRProgressRecordReqContentBlock `json:"blocks,omitempty"` // 文档结构是按行排列的, 每行内容是一个 Block
}

// CreateOKRProgressRecordReqContentBlock ...
type CreateOKRProgressRecordReqContentBlock struct {
	Type      *string                                          `json:"type,omitempty"`      // 文档元素类型, 示例值: "paragraph", 可选值有: paragraph: 文本段落, gallery: 图片
	Paragraph *CreateOKRProgressRecordReqContentBlockParagraph `json:"paragraph,omitempty"` // 文本段落
	Gallery   *CreateOKRProgressRecordReqContentBlockGallery   `json:"gallery,omitempty"`   // 图片
}

// CreateOKRProgressRecordReqContentBlockGallery ...
type CreateOKRProgressRecordReqContentBlockGallery struct {
	ImageList []*CreateOKRProgressRecordReqContentBlockGalleryImageList `json:"imageList,omitempty"` // 图片元素
}

// CreateOKRProgressRecordReqContentBlockGalleryImageList ...
type CreateOKRProgressRecordReqContentBlockGalleryImageList struct {
	FileToken *string  `json:"fileToken,omitempty"` // 图片 token, 通过上传图片接口获取, 示例值: "boxcnOj88GDkmWGm2zsTyCBqoLb"
	Src       *string  `json:"src,omitempty"`       // 图片链接, 示例值: "https://bytedance.feishu.cn/drive/home/"
	Width     *float64 `json:"width,omitempty"`     // 图片宽, 单位px, 示例值: 458
	Height    *float64 `json:"height,omitempty"`    // 图片高, 单位px, 示例值: 372
}

// CreateOKRProgressRecordReqContentBlockParagraph ...
type CreateOKRProgressRecordReqContentBlockParagraph struct {
	Style    *CreateOKRProgressRecordReqContentBlockParagraphStyle     `json:"style,omitempty"`    // 段落样式
	Elements []*CreateOKRProgressRecordReqContentBlockParagraphElement `json:"elements,omitempty"` // 段落元素组成一个段落
}

// CreateOKRProgressRecordReqContentBlockParagraphElement ...
type CreateOKRProgressRecordReqContentBlockParagraphElement struct {
	Type     *string                                                         `json:"type,omitempty"`     // 元素类型, 示例值: "textRun", 可选值有: textRun: 文本型元素, docsLink: 文档链接型元素, person: 艾特用户型元素
	TextRun  *CreateOKRProgressRecordReqContentBlockParagraphElementTextRun  `json:"textRun,omitempty"`  // 文本
	DocsLink *CreateOKRProgressRecordReqContentBlockParagraphElementDocsLink `json:"docsLink,omitempty"` // 飞书云文档
	Person   *CreateOKRProgressRecordReqContentBlockParagraphElementPerson   `json:"person,omitempty"`   // 艾特用户
}

// CreateOKRProgressRecordReqContentBlockParagraphElementDocsLink ...
type CreateOKRProgressRecordReqContentBlockParagraphElementDocsLink struct {
	URL   *string `json:"url,omitempty"`   // 飞书云文档链接地址, 示例值: "https://xxx.feishu.cn/docx/xxxxxxxx"
	Title *string `json:"title,omitempty"` // 飞书云文档标题, 示例值: "项目说明文档"
}

// CreateOKRProgressRecordReqContentBlockParagraphElementPerson ...
type CreateOKRProgressRecordReqContentBlockParagraphElementPerson struct {
	OpenID *string `json:"openId,omitempty"` // 员工的OpenID, 示例值: "ou_3bbe8a09c20e89cce9bff989ed840674"
}

// CreateOKRProgressRecordReqContentBlockParagraphElementTextRun ...
type CreateOKRProgressRecordReqContentBlockParagraphElementTextRun struct {
	Text  *string                                                             `json:"text,omitempty"`  // 具体的文本内容, 示例值: "周报内容"
	Style *CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyle `json:"style,omitempty"` // 文本内容的样式, 支持 BIUS、颜色等
}

// CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyle ...
type CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyle struct {
	Bold          *bool                                                                        `json:"bold,omitempty"`          // 是否加粗, 示例值: true
	StrikeThrough *bool                                                                        `json:"strikeThrough,omitempty"` // 是否删除, 示例值: true
	BackColor     *CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleBackColor `json:"backColor,omitempty"`     // 背景颜色
	TextColor     *CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleTextColor `json:"textColor,omitempty"`     // 字体颜色
	Link          *CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleLink      `json:"link,omitempty"`          // 链接地址
}

// CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleBackColor ...
type CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleBackColor struct {
	Red   *int64   `json:"red,omitempty"`   // 红 取值范围[0, 255], 示例值: 216
	Green *int64   `json:"green,omitempty"` // 绿 取值范围[0, 255], 示例值: 191
	Blue  *int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255], 示例值: 188
	Alpha *float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1], 示例值: 0.1
}

// CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleLink ...
type CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleLink struct {
	URL *string `json:"url,omitempty"` // 链接地址, 示例值: "https://www.xxxxx.com/"
}

// CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleTextColor ...
type CreateOKRProgressRecordReqContentBlockParagraphElementTextRunStyleTextColor struct {
	Red   *int64   `json:"red,omitempty"`   // 红 取值范围[0, 255], 示例值: 216
	Green *int64   `json:"green,omitempty"` // 绿 取值范围[0, 255], 示例值: 191
	Blue  *int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255], 示例值: 188
	Alpha *float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1], 示例值: 0.1
}

// CreateOKRProgressRecordReqContentBlockParagraphStyle ...
type CreateOKRProgressRecordReqContentBlockParagraphStyle struct {
	List *CreateOKRProgressRecordReqContentBlockParagraphStyleList `json:"list,omitempty"` // 有序列表/无序列表/任务列表
}

// CreateOKRProgressRecordReqContentBlockParagraphStyleList ...
type CreateOKRProgressRecordReqContentBlockParagraphStyleList struct {
	Type        *string `json:"type,omitempty"`        // 列表类型, 示例值: "number", 可选值有: number: 有序列表, bullet: 无序列表, checkBox: 任务列表, checkedBox: 已完成的任务列表, indent: tab缩进
	IndentLevel *int64  `json:"indentLevel,omitempty"` // 列表的缩进级别, 支持指定一行的缩进 除代码块以外的列表都支持设置缩进, 支持 1-16 级缩进, 取值范围: [1, 16], 示例值: 1
	Number      *int64  `json:"number,omitempty"`      // 用于指定列表的行号, 仅对有序列表和代码块生效 如果为有序列表设置了缩进, 行号可能会显示为字母或者罗马数字, 示例值: 1
}

// CreateOKRProgressRecordResp ...
type CreateOKRProgressRecordResp struct {
	ProgressID string                              `json:"progress_id,omitempty"` // OKR 进展ID
	ModifyTime string                              `json:"modify_time,omitempty"` // 进展更新时间 毫秒
	Content    *CreateOKRProgressRecordRespContent `json:"content,omitempty"`     // 进展 对应的 Content 详细内容
}

// CreateOKRProgressRecordRespContent ...
type CreateOKRProgressRecordRespContent struct {
	Blocks []*CreateOKRProgressRecordRespContentBlock `json:"blocks,omitempty"` // 文档结构是按行排列的, 每行内容是一个 Block
}

// CreateOKRProgressRecordRespContentBlock ...
type CreateOKRProgressRecordRespContentBlock struct {
	Type      string                                            `json:"type,omitempty"`      // 文档元素类型, 可选值有: paragraph: 文本段落, gallery: 图片
	Paragraph *CreateOKRProgressRecordRespContentBlockParagraph `json:"paragraph,omitempty"` // 文本段落
	Gallery   *CreateOKRProgressRecordRespContentBlockGallery   `json:"gallery,omitempty"`   // 图片
}

// CreateOKRProgressRecordRespContentBlockGallery ...
type CreateOKRProgressRecordRespContentBlockGallery struct {
	ImageList []*CreateOKRProgressRecordRespContentBlockGalleryImageList `json:"imageList,omitempty"` // 图片元素
}

// CreateOKRProgressRecordRespContentBlockGalleryImageList ...
type CreateOKRProgressRecordRespContentBlockGalleryImageList struct {
	FileToken string  `json:"fileToken,omitempty"` // 图片 token, 通过上传图片接口获取
	Src       string  `json:"src,omitempty"`       // 图片链接
	Width     float64 `json:"width,omitempty"`     // 图片宽, 单位px
	Height    float64 `json:"height,omitempty"`    // 图片高, 单位px
}

// CreateOKRProgressRecordRespContentBlockParagraph ...
type CreateOKRProgressRecordRespContentBlockParagraph struct {
	Style    *CreateOKRProgressRecordRespContentBlockParagraphStyle     `json:"style,omitempty"`    // 段落样式
	Elements []*CreateOKRProgressRecordRespContentBlockParagraphElement `json:"elements,omitempty"` // 段落元素组成一个段落
}

// CreateOKRProgressRecordRespContentBlockParagraphElement ...
type CreateOKRProgressRecordRespContentBlockParagraphElement struct {
	Type     string                                                           `json:"type,omitempty"`     // 元素类型, 可选值有: textRun: 文本型元素, docsLink: 文档链接型元素, person: 艾特用户型元素
	TextRun  *CreateOKRProgressRecordRespContentBlockParagraphElementTextRun  `json:"textRun,omitempty"`  // 文本
	DocsLink *CreateOKRProgressRecordRespContentBlockParagraphElementDocsLink `json:"docsLink,omitempty"` // 飞书云文档
	Person   *CreateOKRProgressRecordRespContentBlockParagraphElementPerson   `json:"person,omitempty"`   // 艾特用户
}

// CreateOKRProgressRecordRespContentBlockParagraphElementDocsLink ...
type CreateOKRProgressRecordRespContentBlockParagraphElementDocsLink struct {
	URL   string `json:"url,omitempty"`   // 飞书云文档链接地址
	Title string `json:"title,omitempty"` // 飞书云文档标题
}

// CreateOKRProgressRecordRespContentBlockParagraphElementPerson ...
type CreateOKRProgressRecordRespContentBlockParagraphElementPerson struct {
	OpenID string `json:"openId,omitempty"` // 员工的OpenID
}

// CreateOKRProgressRecordRespContentBlockParagraphElementTextRun ...
type CreateOKRProgressRecordRespContentBlockParagraphElementTextRun struct {
	Text  string                                                               `json:"text,omitempty"`  // 具体的文本内容
	Style *CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyle `json:"style,omitempty"` // 文本内容的样式, 支持 BIUS、颜色等
}

// CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyle ...
type CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyle struct {
	Bold          bool                                                                          `json:"bold,omitempty"`          // 是否加粗
	StrikeThrough bool                                                                          `json:"strikeThrough,omitempty"` // 是否删除
	BackColor     *CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor `json:"backColor,omitempty"`     // 背景颜色
	TextColor     *CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor `json:"textColor,omitempty"`     // 字体颜色
	Link          *CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink      `json:"link,omitempty"`          // 链接地址
}

// CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor ...
type CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleBackColor struct {
	Red   int64   `json:"red,omitempty"`   // 红 取值范围[0, 255]
	Green int64   `json:"green,omitempty"` // 绿 取值范围[0, 255]
	Blue  int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255]
	Alpha float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1]
}

// CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink ...
type CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleLink struct {
	URL string `json:"url,omitempty"` // 链接地址
}

// CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor ...
type CreateOKRProgressRecordRespContentBlockParagraphElementTextRunStyleTextColor struct {
	Red   int64   `json:"red,omitempty"`   // 红 取值范围[0, 255]
	Green int64   `json:"green,omitempty"` // 绿 取值范围[0, 255]
	Blue  int64   `json:"blue,omitempty"`  // 蓝 取值范围[0, 255]
	Alpha float64 `json:"alpha,omitempty"` // 透明度 取值范围[0, 1]
}

// CreateOKRProgressRecordRespContentBlockParagraphStyle ...
type CreateOKRProgressRecordRespContentBlockParagraphStyle struct {
	List *CreateOKRProgressRecordRespContentBlockParagraphStyleList `json:"list,omitempty"` // 有序列表/无序列表/任务列表
}

// CreateOKRProgressRecordRespContentBlockParagraphStyleList ...
type CreateOKRProgressRecordRespContentBlockParagraphStyleList struct {
	Type        string `json:"type,omitempty"`        // 列表类型, 可选值有: number: 有序列表, bullet: 无序列表, checkBox: 任务列表, checkedBox: 已完成的任务列表, indent: tab缩进
	IndentLevel int64  `json:"indentLevel,omitempty"` // 列表的缩进级别, 支持指定一行的缩进 除代码块以外的列表都支持设置缩进, 支持 1-16 级缩进, 取值范围: [1, 16]
	Number      int64  `json:"number,omitempty"`      // 用于指定列表的行号, 仅对有序列表和代码块生效 如果为有序列表设置了缩进, 行号可能会显示为字母或者罗马数字
}

// createOKRProgressRecordResp ...
type createOKRProgressRecordResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *CreateOKRProgressRecordResp `json:"data,omitempty"`
}
