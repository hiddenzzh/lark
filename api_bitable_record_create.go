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

// CreateBitableRecord 该接口用于在数据表中新增一条记录
//
// 该接口支持调用频率上限为 10 QPS（Query Per Second, 每秒请求率）
// ::: note
// 首次调用请参考 [云文档接口快速入门](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)[多维表格接口接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/create
func (r *BitableService) CreateBitableRecord(ctx context.Context, request *CreateBitableRecordReq, options ...MethodOptionFunc) (*CreateBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableRecord mock enable")
		return r.cli.mock.mockBitableCreateBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableRecord mock BitableCreateBitableRecord method
func (r *Mock) MockBitableCreateBitableRecord(f func(ctx context.Context, request *CreateBitableRecordReq, options ...MethodOptionFunc) (*CreateBitableRecordResp, *Response, error)) {
	r.mockBitableCreateBitableRecord = f
}

// UnMockBitableCreateBitableRecord un-mock BitableCreateBitableRecord method
func (r *Mock) UnMockBitableCreateBitableRecord() {
	r.mockBitableCreateBitableRecord = nil
}

// CreateBitableRecordReq ...
type CreateBitableRecordReq struct {
	AppToken   string                 `path:"app_token" json:"-"`     // 多维表格的唯一标识符 [app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "bascng7vrxcxpig7geggXiCtadY"
	TableID    string                 `path:"table_id" json:"-"`      // 多维表格数据表的唯一标识符 [table_id 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#735fe883), 示例值: "tblUa9vcYjWQYJCj"
	UserIDType *IDType                `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Fields     map[string]interface{} `json:"fields,omitempty"`       // 数据表的字段, 即数据表的列, 当前接口支持的字段类型请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#31f78a3c), 不同类型字段的数据结构请参考[数据结构概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/development-guide/bitable-structure)
}

// CreateBitableRecordResp ...
type CreateBitableRecordResp struct {
	Record *CreateBitableRecordRespRecord `json:"record,omitempty"` // 新增的记录的内容
}

// CreateBitableRecordRespRecord ...
type CreateBitableRecordRespRecord struct {
	RecordID         string                                       `json:"record_id,omitempty"`          // 一条记录的唯一标识 id [record_id 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#15d8db94)
	CreatedBy        *CreateBitableRecordRespRecordCreatedBy      `json:"created_by,omitempty"`         // 该记录的创建人
	CreatedTime      int64                                        `json:"created_time,omitempty"`       // 该记录的创建时间
	LastModifiedBy   *CreateBitableRecordRespRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 该记录最新一次更新的修改人
	LastModifiedTime int64                                        `json:"last_modified_time,omitempty"` // 该记录最近一次的更新时间
	Fields           map[string]interface{}                       `json:"fields,omitempty"`             // 数据表的字段, 即数据表的列, 当前接口支持的字段类型请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#31f78a3c), 不同类型字段的数据结构请参考[数据结构概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/development-guide/bitable-structure)
}

// CreateBitableRecordRespRecordCreatedBy ...
type CreateBitableRecordRespRecordCreatedBy struct {
	ID     string `json:"id,omitempty"`      // 用户id, id类型等于user_id_type所指定的类型。
	Name   string `json:"name,omitempty"`    // 用户的中文名称
	EnName string `json:"en_name,omitempty"` // 用户的英文名称
	Email  string `json:"email,omitempty"`   // 用户的邮箱
}

// CreateBitableRecordRespRecordLastModifiedBy ...
type CreateBitableRecordRespRecordLastModifiedBy struct {
	ID     string `json:"id,omitempty"`      // 用户id, id类型等于user_id_type所指定的类型。
	Name   string `json:"name,omitempty"`    // 用户的中文名称
	EnName string `json:"en_name,omitempty"` // 用户的英文名称
	Email  string `json:"email,omitempty"`   // 用户的邮箱
}

// createBitableRecordResp ...
type createBitableRecordResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableRecordResp `json:"data,omitempty"`
}
