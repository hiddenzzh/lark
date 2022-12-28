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

// CreateApprovalInstance 创建一个审批实例, 调用方需对审批定义的表单有详细了解, 将按照定义的表单结构, 将表单 Value 通过接口传入。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/create
func (r *ApprovalService) CreateApprovalInstance(ctx context.Context, request *CreateApprovalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalCreateApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#CreateApprovalInstance mock enable")
		return r.cli.mock.mockApprovalCreateApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CreateApprovalInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalCreateApprovalInstance mock ApprovalCreateApprovalInstance method
func (r *Mock) MockApprovalCreateApprovalInstance(f func(ctx context.Context, request *CreateApprovalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalInstanceResp, *Response, error)) {
	r.mockApprovalCreateApprovalInstance = f
}

// UnMockApprovalCreateApprovalInstance un-mock ApprovalCreateApprovalInstance method
func (r *Mock) UnMockApprovalCreateApprovalInstance() {
	r.mockApprovalCreateApprovalInstance = nil
}

// CreateApprovalInstanceReq ...
type CreateApprovalInstanceReq struct {
	ApprovalCode           string                                             `json:"approval_code,omitempty"`              // 审批定义 code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	UserID                 *string                                            `json:"user_id,omitempty"`                    // 发起审批用户, 示例值: "f7cb567e"
	OpenID                 string                                             `json:"open_id,omitempty"`                    // 发起审批用户 open id, 如果传了 user_id 则优先使用 user_id, 示例值: "ou_3cda9c969f737aaa05e6915dce306cb9"
	DepartmentID           *string                                            `json:"department_id,omitempty"`              // 发起审批用户部门id, 如果用户只属于一个部门, 可以不填。如果属于多个部门, 默认会选择部门列表第一个部门, 示例值: "9293493ccacbdb9a"
	Form                   ApprovalWidgetList                                 `json:"form,omitempty"`                       // json 数组, 控件值, 示例值: "[{\"id\":\"111\", \"type\": \"input\", \"value\":\"test\"}]"
	NodeApproverUserIDList []*CreateApprovalInstanceReqNodeApproverUserIDList `json:"node_approver_user_id_list,omitempty"` // 如果有发起人自选节点, 则需要填写对应节点的审批人
	NodeApproverOpenIDList []*CreateApprovalInstanceReqNodeApproverOpenIDList `json:"node_approver_open_id_list,omitempty"` // 审批人发起人自选 open id, 与上述node_approver_user_id_list字段取并集
	NodeCcUserIDList       []*CreateApprovalInstanceReqNodeCcUserIDList       `json:"node_cc_user_id_list,omitempty"`       // 如果有发起人自选节点, 则可填写对应节点的抄送人, 单个节点最多选择20位抄送人, 最大长度: `20`
	NodeCcOpenIDList       []*CreateApprovalInstanceReqNodeCcOpenIDList       `json:"node_cc_open_id_list,omitempty"`       // 抄送人发起人自选 open id 单个节点最多选择20位抄送人, 最大长度: `20`
	UUID                   *string                                            `json:"uuid,omitempty"`                       // 审批实例 uuid, 用于幂等操作, 每个租户下面的唯一key, 同一个 uuid 只能用于创建一个审批实例, 如果冲突, 返回错误码 60012, 格式建议为 XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX, 不区分大小写, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A87", 长度范围: `1` ～ `64` 字符
	AllowResubmit          *bool                                              `json:"allow_resubmit,omitempty"`             // 可配置是否可以再次提交, 示例值: true
	AllowSubmitAgain       *bool                                              `json:"allow_submit_again,omitempty"`         // 可配置是否可以重新提交, 示例值: true
}

// CreateApprovalInstanceReqNodeApproverOpenIDList ...
type CreateApprovalInstanceReqNodeApproverOpenIDList struct {
	Key   *string  `json:"key,omitempty"`   // node id 或 custom node id, 通过 [查看审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get) 获取, 示例值: "46e6d96cfa756980907209209ec03b64"
	Value []string `json:"value,omitempty"` // value: 审批人列表, 示例值: ["f7cb567e"]
}

// CreateApprovalInstanceReqNodeApproverUserIDList ...
type CreateApprovalInstanceReqNodeApproverUserIDList struct {
	Key   *string  `json:"key,omitempty"`   // node id 或 custom node id, 通过 [查看审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get) 获取, 示例值: "46e6d96cfa756980907209209ec03b64"
	Value []string `json:"value,omitempty"` // value: 审批人列表, 示例值: ["f7cb567e"]
}

// CreateApprovalInstanceReqNodeCcOpenIDList ...
type CreateApprovalInstanceReqNodeCcOpenIDList struct {
	Key   *string  `json:"key,omitempty"`   // node id, 通过 [查看审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get) 获取, 示例值: "46e6d96cfa756980907209209ec03b75"
	Value []string `json:"value,omitempty"` // value: 审批人列表, 示例值: ["f7cb567e"]
}

// CreateApprovalInstanceReqNodeCcUserIDList ...
type CreateApprovalInstanceReqNodeCcUserIDList struct {
	Key   *string  `json:"key,omitempty"`   // node id, 通过 [查看审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get) 获取, 示例值: "46e6d96cfa756980907209209ec03b75"
	Value []string `json:"value,omitempty"` // value: 审批人列表, 示例值: ["f7cb567e"]
}

// CreateApprovalInstanceResp ...
type CreateApprovalInstanceResp struct {
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例 Code
}

// createApprovalInstanceResp ...
type createApprovalInstanceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateApprovalInstanceResp `json:"data,omitempty"`
}
