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

// DeleteBitableView 删除数据表中的视图
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/delete
func (r *BitableService) DeleteBitableView(ctx context.Context, request *DeleteBitableViewReq, options ...MethodOptionFunc) (*DeleteBitableViewResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableView mock enable")
		return r.cli.mock.mockBitableDeleteBitableView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "DeleteBitableView",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views/:view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteBitableViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableDeleteBitableView mock BitableDeleteBitableView method
func (r *Mock) MockBitableDeleteBitableView(f func(ctx context.Context, request *DeleteBitableViewReq, options ...MethodOptionFunc) (*DeleteBitableViewResp, *Response, error)) {
	r.mockBitableDeleteBitableView = f
}

// UnMockBitableDeleteBitableView un-mock BitableDeleteBitableView method
func (r *Mock) UnMockBitableDeleteBitableView() {
	r.mockBitableDeleteBitableView = nil
}

// DeleteBitableViewReq ...
type DeleteBitableViewReq struct {
	AppToken string `path:"app_token" json:"-"` // base app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值: "tblsRc9GRRXKqhvW"
	ViewID   string `path:"view_id" json:"-"`   // 视图Id, 示例值: "vewTpR1urY"
}

// DeleteBitableViewResp ...
type DeleteBitableViewResp struct {
}

// deleteBitableViewResp ...
type deleteBitableViewResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteBitableViewResp `json:"data,omitempty"`
}
