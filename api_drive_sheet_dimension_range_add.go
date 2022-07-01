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

// AddSheetDimensionRange 该接口用于根据 spreadsheetToken 和长度, 在末尾增加空行/列；单次操作不超过5000行或列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUjMzUjL1IzM14SNyMTN
func (r *DriveService) AddSheetDimensionRange(ctx context.Context, request *AddSheetDimensionRangeReq, options ...MethodOptionFunc) (*AddSheetDimensionRangeResp, *Response, error) {
	if r.cli.mock.mockDriveAddSheetDimensionRange != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#AddSheetDimensionRange mock enable")
		return r.cli.mock.mockDriveAddSheetDimensionRange(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "AddSheetDimensionRange",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(addSheetDimensionRangeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveAddSheetDimensionRange mock DriveAddSheetDimensionRange method
func (r *Mock) MockDriveAddSheetDimensionRange(f func(ctx context.Context, request *AddSheetDimensionRangeReq, options ...MethodOptionFunc) (*AddSheetDimensionRangeResp, *Response, error)) {
	r.mockDriveAddSheetDimensionRange = f
}

// UnMockDriveAddSheetDimensionRange un-mock DriveAddSheetDimensionRange method
func (r *Mock) UnMockDriveAddSheetDimensionRange() {
	r.mockDriveAddSheetDimensionRange = nil
}

// AddSheetDimensionRangeReq ...
type AddSheetDimensionRangeReq struct {
	SpreadSheetToken string                              `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token, 详见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Dimension        *AddSheetDimensionRangeReqDimension `json:"dimension,omitempty"`       // 需要增加行列的维度信息
}

// AddSheetDimensionRangeReqDimension ...
type AddSheetDimensionRangeReqDimension struct {
	SheetID        string  `json:"sheetId,omitempty"`        // sheetId
	MajorDimension *string `json:"majorDimension,omitempty"` // 默认 ROWS, 可选 ROWS、COLUMNS
	Length         int64   `json:"length,omitempty"`         // 要增加的行/列数, 0<length<5000
}

// AddSheetDimensionRangeResp ...
type AddSheetDimensionRangeResp struct {
	AddCount       int64  `json:"addCount,omitempty"`       // 增加的行/列数
	MajorDimension string `json:"majorDimension,omitempty"` // 插入维度
}

// addSheetDimensionRangeResp ...
type addSheetDimensionRangeResp struct {
	Code int64                       `json:"code,omitempty"`
	Msg  string                      `json:"msg,omitempty"`
	Data *AddSheetDimensionRangeResp `json:"data,omitempty"`
}
