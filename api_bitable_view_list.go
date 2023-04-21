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

// GetBitableViewList 根据 app_token 和 table_id, 获取数据表的所有视图
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list
func (r *BitableService) GetBitableViewList(ctx context.Context, request *GetBitableViewListReq, options ...MethodOptionFunc) (*GetBitableViewListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableViewList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableViewList mock enable")
		return r.cli.mock.mockBitableGetBitableViewList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableViewList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableViewListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableViewList mock BitableGetBitableViewList method
func (r *Mock) MockBitableGetBitableViewList(f func(ctx context.Context, request *GetBitableViewListReq, options ...MethodOptionFunc) (*GetBitableViewListResp, *Response, error)) {
	r.mockBitableGetBitableViewList = f
}

// UnMockBitableGetBitableViewList un-mock BitableGetBitableViewList method
func (r *Mock) UnMockBitableGetBitableViewList() {
	r.mockBitableGetBitableViewList = nil
}

// GetBitableViewListReq ...
type GetBitableViewListReq struct {
	AppToken   string  `path:"app_token" json:"-"`     // base app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh", 最小长度: `1` 字符
	TableID    string  `path:"table_id" json:"-"`      // table id, 示例值: "tblsRc9GRRXKqhvW"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `100`
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "vewTpR1urY"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetBitableViewListResp ...
type GetBitableViewListResp struct {
	Items     []*GetBitableViewListRespItem `json:"items,omitempty"`      // 视图信息
	PageToken string                        `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有更多项
	Total     int64                         `json:"total,omitempty"`      // 总数
}

// GetBitableViewListRespItem ...
type GetBitableViewListRespItem struct {
	ViewID             string `json:"view_id,omitempty"`               // 视图Id
	ViewName           string `json:"view_name,omitempty"`             // 视图名字
	ViewType           string `json:"view_type,omitempty"`             // 视图类型
	ViewPublicLevel    string `json:"view_public_level,omitempty"`     // 视图公共等级 Public、Locked、Private, 可选值有: Public: 公共视图, Locked: 锁定视图, Private: 个人视图
	ViewPrivateOwnerID string `json:"view_private_owner_id,omitempty"` // 个人视图的owner_id, id类型和 user_id_type 参数保持一致
}

// getBitableViewListResp ...
type getBitableViewListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableViewListResp `json:"data,omitempty"`
}
