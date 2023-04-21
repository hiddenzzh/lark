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

// BatchDeleteBitableTable 删除多个数据表。
//
// ::: note
// 首次调用请参考 [云文档接口快速入门](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)[多维表格接口接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete
func (r *BitableService) BatchDeleteBitableTable(ctx context.Context, request *BatchDeleteBitableTableReq, options ...MethodOptionFunc) (*BatchDeleteBitableTableResp, *Response, error) {
	if r.cli.mock.mockBitableBatchDeleteBitableTable != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchDeleteBitableTable mock enable")
		return r.cli.mock.mockBitableBatchDeleteBitableTable(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchDeleteBitableTable",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchDeleteBitableTableResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchDeleteBitableTable mock BitableBatchDeleteBitableTable method
func (r *Mock) MockBitableBatchDeleteBitableTable(f func(ctx context.Context, request *BatchDeleteBitableTableReq, options ...MethodOptionFunc) (*BatchDeleteBitableTableResp, *Response, error)) {
	r.mockBitableBatchDeleteBitableTable = f
}

// UnMockBitableBatchDeleteBitableTable un-mock BitableBatchDeleteBitableTable method
func (r *Mock) UnMockBitableBatchDeleteBitableTable() {
	r.mockBitableBatchDeleteBitableTable = nil
}

// BatchDeleteBitableTableReq ...
type BatchDeleteBitableTableReq struct {
	AppToken string   `path:"app_token" json:"-"`  // 多维表格的唯一标识符 [app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "appbcbWCzen6D8dezhoCH2RpMAh", 最小长度: `1` 字符
	TableIDs []string `json:"table_ids,omitempty"` // 待删除的数据表的id [table_id 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#735fe883), 示例值: ["tblsRc9GRRXKqhvW"]
}

// BatchDeleteBitableTableResp ...
type BatchDeleteBitableTableResp struct {
}

// batchDeleteBitableTableResp ...
type batchDeleteBitableTableResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteBitableTableResp `json:"data,omitempty"`
}
