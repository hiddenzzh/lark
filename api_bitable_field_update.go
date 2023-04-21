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

// UpdateBitableField 该接口用于在数据表中更新一个字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/update
func (r *BitableService) UpdateBitableField(ctx context.Context, request *UpdateBitableFieldReq, options ...MethodOptionFunc) (*UpdateBitableFieldResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateBitableField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#UpdateBitableField mock enable")
		return r.cli.mock.mockBitableUpdateBitableField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "UpdateBitableField",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateBitableFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableUpdateBitableField mock BitableUpdateBitableField method
func (r *Mock) MockBitableUpdateBitableField(f func(ctx context.Context, request *UpdateBitableFieldReq, options ...MethodOptionFunc) (*UpdateBitableFieldResp, *Response, error)) {
	r.mockBitableUpdateBitableField = f
}

// UnMockBitableUpdateBitableField un-mock BitableUpdateBitableField method
func (r *Mock) UnMockBitableUpdateBitableField() {
	r.mockBitableUpdateBitableField = nil
}

// UpdateBitableFieldReq ...
type UpdateBitableFieldReq struct {
	AppToken    string                            `path:"app_token" json:"-"`    // base app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID     string                            `path:"table_id" json:"-"`     // table id, 示例值: "tblsRc9GRRXKqhvW"
	FieldID     string                            `path:"field_id" json:"-"`     // field id, 示例值: "fldPTb0U2y"
	FieldName   string                            `json:"field_name,omitempty"`  // 多维表格字段名, 请注意: 1. 名称中的首尾空格将会被去除, 示例值: "多行文本"
	Type        int64                             `json:"type,omitempty"`        // 多维表格字段类型, 示例值: 1, 可选值有: 1: 多行文本, 2: 数字, 3: 单选, 4: 多选, 5: 日期, 7: 复选框, 11: 人员, 15: 超链接, 17: 附件, 18: 关联, 20: 公式, 21: 双向关联, 1001: 创建时间, 1002: 最后更新时间, 1003: 创建人, 1004: 修改人, 1005: 自动编号, 13: 电话号码, 22: 地理位置
	Property    *UpdateBitableFieldReqProperty    `json:"property,omitempty"`    // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
	Description *UpdateBitableFieldReqDescription `json:"description,omitempty"` // 字段的描述
}

// UpdateBitableFieldReqDescription ...
type UpdateBitableFieldReqDescription struct {
	DisableSync *bool   `json:"disable_sync,omitempty"` // 是否禁止同步, 如果为true, 表示禁止同步该描述内容到表单的问题描述（只在新增、修改字段时生效）, 示例值: true, 默认值: `true`
	Text        *string `json:"text,omitempty"`         // 字段描述内容, 示例值: "这是一个字段描述"
}

// UpdateBitableFieldReqProperty ...
type UpdateBitableFieldReqProperty struct {
	Options           []*UpdateBitableFieldReqPropertyOption   `json:"options,omitempty"`            // 单选、多选字段的选项信息
	Formatter         *string                                  `json:"formatter,omitempty"`          // 数字、公式字段的显示格式, 示例值: "0"
	DateFormatter     *string                                  `json:"date_formatter,omitempty"`     // 日期、创建时间、最后更新时间字段的显示格式, 示例值: "日期格式"
	AutoFill          *bool                                    `json:"auto_fill,omitempty"`          // 日期字段中新纪录自动填写创建时间, 示例值: false
	Multiple          *bool                                    `json:"multiple,omitempty"`           // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录, 示例值: false
	TableID           *string                                  `json:"table_id,omitempty"`           // 单向关联、双向关联字段中关联的数据表的id, 示例值: "tblsRc9GRRXKqhvW"
	TableName         *string                                  `json:"table_name,omitempty"`         // 单向关联、双向关联字段中关联的数据表的名字, 示例值: ""table2""
	BackFieldName     *string                                  `json:"back_field_name,omitempty"`    // 双向关联字段中关联的数据表中对应的双向关联字段的名字, 示例值: ""table1-双向关联""
	AutoSerial        *UpdateBitableFieldReqPropertyAutoSerial `json:"auto_serial,omitempty"`        // 自动编号类型
	Location          *UpdateBitableFieldReqPropertyLocation   `json:"location,omitempty"`           // 地理位置输入方式
	FormulaExpression *string                                  `json:"formula_expression,omitempty"` // 公式字段的表达式, 示例值: "bitable::$table[tblNj92WQBAasdEf].$field[fldMV60rYs]*2"
}

// UpdateBitableFieldReqPropertyAutoSerial ...
type UpdateBitableFieldReqPropertyAutoSerial struct {
	Type    string                                           `json:"type,omitempty"`    // 自动编号类型, 示例值: "auto_increment_number", 可选值有: custom: 自定义编号, auto_increment_number: 自增数字
	Options []*UpdateBitableFieldReqPropertyAutoSerialOption `json:"options,omitempty"` // 自动编号规则列表
}

// UpdateBitableFieldReqPropertyAutoSerialOption ...
type UpdateBitableFieldReqPropertyAutoSerialOption struct {
	Type  string `json:"type,omitempty"`  // 自动编号的可选规则项类型, 示例值: "created_time", 可选值有: system_number: 自增数字位, value范围1-9, fixed_text: 固定字符, 最大长度: 20, created_time: 创建时间, 支持格式 "yyyyMMdd"、"yyyyMM"、"yyyy"、"MMdd"、"MM"、"dd"
	Value string `json:"value,omitempty"` // 与自动编号的可选规则项类型相对应的取值, 示例值: "yyyyMMdd"
}

// UpdateBitableFieldReqPropertyLocation ...
type UpdateBitableFieldReqPropertyLocation struct {
	InputType string `json:"input_type,omitempty"` // 地理位置输入限制, 示例值: "not_limit", 可选值有: only_mobile: 只允许移动端上传, not_limit: 无限制
}

// UpdateBitableFieldReqPropertyOption ...
type UpdateBitableFieldReqPropertyOption struct {
	Name  *string `json:"name,omitempty"`  // 选项名, 示例值: "红色"
	ID    *string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID, 示例值: "optKl35lnG"
	Color *int64  `json:"color,omitempty"` // 选项颜色, 示例值: 0, 取值范围: `0` ～ `54`
}

// UpdateBitableFieldResp ...
type UpdateBitableFieldResp struct {
	Field *UpdateBitableFieldRespField `json:"field,omitempty"` // 字段
}

// UpdateBitableFieldRespField ...
type UpdateBitableFieldRespField struct {
	FieldID     string                                  `json:"field_id,omitempty"`    // 多维表格字段 id
	FieldName   string                                  `json:"field_name,omitempty"`  // 多维表格字段名, 请注意: 1. 名称中的首尾空格将会被去除。
	Type        int64                                   `json:"type,omitempty"`        // 多维表格字段类型, 可选值有: 1: 多行文本, 2: 数字, 3: 单选, 4: 多选, 5: 日期, 7: 复选框, 11: 人员, 15: 超链接, 17: 附件, 18: 关联, 20: 公式, 21: 双向关联, 1001: 创建时间, 1002: 最后更新时间, 1003: 创建人, 1004: 修改人, 1005: 自动编号, 13: 电话号码, 22: 地理位置
	Property    *UpdateBitableFieldRespFieldProperty    `json:"property,omitempty"`    // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
	Description *UpdateBitableFieldRespFieldDescription `json:"description,omitempty"` // 字段的描述
	IsPrimary   bool                                    `json:"is_primary,omitempty"`  // 是否是索引列
}

// UpdateBitableFieldRespFieldDescription ...
type UpdateBitableFieldRespFieldDescription struct {
	DisableSync bool   `json:"disable_sync,omitempty"` // 是否禁止同步, 如果为true, 表示禁止同步该描述内容到表单的问题描述（只在新增、修改字段时生效）
	Text        string `json:"text,omitempty"`         // 字段描述内容
}

// UpdateBitableFieldRespFieldProperty ...
type UpdateBitableFieldRespFieldProperty struct {
	Options           []*UpdateBitableFieldRespFieldPropertyOption   `json:"options,omitempty"`            // 单选、多选字段的选项信息
	Formatter         string                                         `json:"formatter,omitempty"`          // 数字、公式字段的显示格式
	DateFormatter     string                                         `json:"date_formatter,omitempty"`     // 日期、创建时间、最后更新时间字段的显示格式
	AutoFill          bool                                           `json:"auto_fill,omitempty"`          // 日期字段中新纪录自动填写创建时间
	Multiple          bool                                           `json:"multiple,omitempty"`           // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录
	TableID           string                                         `json:"table_id,omitempty"`           // 单向关联、双向关联字段中关联的数据表的id
	TableName         string                                         `json:"table_name,omitempty"`         // 单向关联、双向关联字段中关联的数据表的名字
	BackFieldName     string                                         `json:"back_field_name,omitempty"`    // 双向关联字段中关联的数据表中对应的双向关联字段的名字
	AutoSerial        *UpdateBitableFieldRespFieldPropertyAutoSerial `json:"auto_serial,omitempty"`        // 自动编号类型
	Location          *UpdateBitableFieldRespFieldPropertyLocation   `json:"location,omitempty"`           // 地理位置输入方式
	FormulaExpression string                                         `json:"formula_expression,omitempty"` // 公式字段的表达式
}

// UpdateBitableFieldRespFieldPropertyAutoSerial ...
type UpdateBitableFieldRespFieldPropertyAutoSerial struct {
	Type    string                                                 `json:"type,omitempty"`    // 自动编号类型, 可选值有: custom: 自定义编号, auto_increment_number: 自增数字
	Options []*UpdateBitableFieldRespFieldPropertyAutoSerialOption `json:"options,omitempty"` // 自动编号规则列表
}

// UpdateBitableFieldRespFieldPropertyAutoSerialOption ...
type UpdateBitableFieldRespFieldPropertyAutoSerialOption struct {
	Type  string `json:"type,omitempty"`  // 自动编号的可选规则项类型, 可选值有: system_number: 自增数字位, value范围1-9, fixed_text: 固定字符, 最大长度: 20, created_time: 创建时间, 支持格式 "yyyyMMdd"、"yyyyMM"、"yyyy"、"MMdd"、"MM"、"dd"
	Value string `json:"value,omitempty"` // 与自动编号的可选规则项类型相对应的取值
}

// UpdateBitableFieldRespFieldPropertyLocation ...
type UpdateBitableFieldRespFieldPropertyLocation struct {
	InputType string `json:"input_type,omitempty"` // 地理位置输入限制, 可选值有: only_mobile: 只允许移动端上传, not_limit: 无限制
}

// UpdateBitableFieldRespFieldPropertyOption ...
type UpdateBitableFieldRespFieldPropertyOption struct {
	Name  string `json:"name,omitempty"`  // 选项名
	ID    string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID
	Color int64  `json:"color,omitempty"` // 选项颜色
}

// updateBitableFieldResp ...
type updateBitableFieldResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBitableFieldResp `json:"data,omitempty"`
}
