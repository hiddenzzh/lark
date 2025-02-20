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

// CancelApprovalInstance 对于状态为“审批中”的单个审批实例进行撤销操作, 撤销后审批流程结束。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/cancel
func (r *ApprovalService) CancelApprovalInstance(ctx context.Context, request *CancelApprovalInstanceReq, options ...MethodOptionFunc) (*CancelApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalCancelApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#CancelApprovalInstance mock enable")
		return r.cli.mock.mockApprovalCancelApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CancelApprovalInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/cancel",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(cancelApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalCancelApprovalInstance mock ApprovalCancelApprovalInstance method
func (r *Mock) MockApprovalCancelApprovalInstance(f func(ctx context.Context, request *CancelApprovalInstanceReq, options ...MethodOptionFunc) (*CancelApprovalInstanceResp, *Response, error)) {
	r.mockApprovalCancelApprovalInstance = f
}

// UnMockApprovalCancelApprovalInstance un-mock ApprovalCancelApprovalInstance method
func (r *Mock) UnMockApprovalCancelApprovalInstance() {
	r.mockApprovalCancelApprovalInstance = nil
}

// CancelApprovalInstanceReq ...
type CancelApprovalInstanceReq struct {
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ApprovalCode string  `json:"approval_code,omitempty"` // 审批定义Code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	InstanceCode string  `json:"instance_code,omitempty"` // 审批实例Code, 示例值: "81D31358-93AF-92D6-7425-01A5D67C4E71"
	UserID       string  `json:"user_id,omitempty"`       // 操作用户, 根据user_id_type填写, 示例值: "f7cb567e"
}

// CancelApprovalInstanceResp ...
type CancelApprovalInstanceResp struct {
}

// cancelApprovalInstanceResp ...
type cancelApprovalInstanceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CancelApprovalInstanceResp `json:"data,omitempty"`
}
