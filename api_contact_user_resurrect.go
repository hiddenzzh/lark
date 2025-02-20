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

// ResurrectUser 该接口用于恢复已删除用户（已离职的成员）, 仅自建应用可申请, 应用商店应用无权调用接口。
//
// - 仅支持恢复离职 30 天内的成员。恢复后, 部分用户数据仍不可恢复, 请谨慎调用。
// - 待恢复成员的用户 ID 不能被企业内其他成员使用。如有重复, 请先离职对应的成员, 否则接口会报错。
// - 待恢复成员的手机号和邮箱不能被企业内其他成员使用。如有重复, 请先修改对应成员的信息, 否则接口会报错。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/resurrect
func (r *ContactService) ResurrectUser(ctx context.Context, request *ResurrectUserReq, options ...MethodOptionFunc) (*ResurrectUserResp, *Response, error) {
	if r.cli.mock.mockContactResurrectUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#ResurrectUser mock enable")
		return r.cli.mock.mockContactResurrectUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "ResurrectUser",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/users/:user_id/resurrect",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(resurrectUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactResurrectUser mock ContactResurrectUser method
func (r *Mock) MockContactResurrectUser(f func(ctx context.Context, request *ResurrectUserReq, options ...MethodOptionFunc) (*ResurrectUserResp, *Response, error)) {
	r.mockContactResurrectUser = f
}

// UnMockContactResurrectUser un-mock ContactResurrectUser method
func (r *Mock) UnMockContactResurrectUser() {
	r.mockContactResurrectUser = nil
}

// ResurrectUserReq ...
type ResurrectUserReq struct {
	UserID           string            `path:"user_id" json:"-"`             // 用户ID, 需要与查询参数中的user_id_type类型保持一致, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 部门ID类型, 示例值: "department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	Departments      []string          `json:"departments,omitempty"`        // 部门, 最大长度: `50`
	SubscriptionIDs  []string          `json:"subscription_ids,omitempty"`   // 指定恢复后分配的席位, 需开通“分配用户席位”权限, 示例值: ["23213213213123123"]
}

// ResurrectUserResp ...
type ResurrectUserResp struct {
}

// resurrectUserResp ...
type resurrectUserResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *ResurrectUserResp `json:"data,omitempty"`
}
