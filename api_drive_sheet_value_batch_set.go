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

// BatchSetSheetValue 该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据, 若范围内有数据, 将被更新覆盖；单次写入不超过5000行, 100列, 每个格子不超过5万字符。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEjMzUjLxIzM14SMyMTN
func (r *DriveService) BatchSetSheetValue(ctx context.Context, request *BatchSetSheetValueReq, options ...MethodOptionFunc) (*BatchSetSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveBatchSetSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetValue mock enable")
		return r.cli.mock.mockDriveBatchSetSheetValue(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchSetSheetValue",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchSetSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveBatchSetSheetValue mock DriveBatchSetSheetValue method
func (r *Mock) MockDriveBatchSetSheetValue(f func(ctx context.Context, request *BatchSetSheetValueReq, options ...MethodOptionFunc) (*BatchSetSheetValueResp, *Response, error)) {
	r.mockDriveBatchSetSheetValue = f
}

// UnMockDriveBatchSetSheetValue un-mock DriveBatchSetSheetValue method
func (r *Mock) UnMockDriveBatchSetSheetValue() {
	r.mockDriveBatchSetSheetValue = nil
}

// BatchSetSheetValueReq ...
type BatchSetSheetValueReq struct {
	SpreadSheetToken string                             `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	ValueRanges      []*BatchSetSheetValueReqValueRange `json:"valueRanges,omitempty"`     // 需要更新的多个范围
}

// BatchSetSheetValueReqValueRange ...
type BatchSetSheetValueReqValueRange struct {
	Range  string           `json:"range,omitempty"`  // 更新范围, 包含 sheetId 与单元格范围两部分, 目前支持三种索引方式, 详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)。range所表示的范围需要大于等于values占用的范围。
	Values [][]SheetContent `json:"values,omitempty"` // 需要写入的值, 如要写入公式、超链接、email、@人等, 可详看附录[sheet 支持写入数据类型](https://open.feishu.cn/document/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN)
}

// BatchSetSheetValueResp ...
type BatchSetSheetValueResp struct {
	Responses        []*BatchSetSheetValueRespResponse `json:"responses,omitempty"`        // 响应
	Revision         int64                             `json:"revision,omitempty"`         // sheet 的版本号
	SpreadSheetToken string                            `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
}

// BatchSetSheetValueRespResponse ...
type BatchSetSheetValueRespResponse struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 写入的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 写入的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 写入的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 写入的单元格总数
}

// batchSetSheetValueResp ...
type batchSetSheetValueResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *BatchSetSheetValueResp `json:"data,omitempty"`
}
