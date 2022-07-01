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

// SetSheetValueImage 该接口用于根据 spreadsheetToken 和 range 向单个格子写入图片。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDNxYjL1QTM24SN0EjN
func (r *DriveService) SetSheetValueImage(ctx context.Context, request *SetSheetValueImageReq, options ...MethodOptionFunc) (*SetSheetValueImageResp, *Response, error) {
	if r.cli.mock.mockDriveSetSheetValueImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#SetSheetValueImage mock enable")
		return r.cli.mock.mockDriveSetSheetValueImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "SetSheetValueImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(setSheetValueImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveSetSheetValueImage mock DriveSetSheetValueImage method
func (r *Mock) MockDriveSetSheetValueImage(f func(ctx context.Context, request *SetSheetValueImageReq, options ...MethodOptionFunc) (*SetSheetValueImageResp, *Response, error)) {
	r.mockDriveSetSheetValueImage = f
}

// UnMockDriveSetSheetValueImage un-mock DriveSetSheetValueImage method
func (r *Mock) UnMockDriveSetSheetValueImage() {
	r.mockDriveSetSheetValueImage = nil
}

// SetSheetValueImageReq ...
type SetSheetValueImageReq struct {
	SpreadSheetToken string `path:"spreadsheetToken" json:"-"` // spreadsheet的token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Range            string `json:"range,omitempty"`           // 查询范围  range=<sheetId>!<开始格子>:<结束格子> 如: xxxx!A1:D5, 详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)。此处限定为一个格子, 如: xxxx!A1:A1
	Image            []byte `json:"image,omitempty"`           // 需要写入的图片二进制流, 支持  "PNG", "JPEG", "JPG", "GIF", "BMP", "JFIF", "EXIF", "TIFF", "BPG", "WEBP", "HEIC" 等图片格式
	Name             string `json:"name,omitempty"`            // 写入的图片名字
}

// SetSheetValueImageResp ...
type SetSheetValueImageResp struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	Revision         int64  `json:"revision,omitempty"`         // spreadsheet 的版本号
	UpdateRange      string `json:"updateRange,omitempty"`      // 写入图片的range
}

// setSheetValueImageResp ...
type setSheetValueImageResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *SetSheetValueImageResp `json:"data,omitempty"`
}
