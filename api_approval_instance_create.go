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
	OpenID                 *string                                            `json:"open_id,omitempty"`                    // 发起审批用户 open id, 与user_id需二者传一, 如果传了 user_id 则优先使用 user_id, 示例值: "ou_3cda9c969f737aaa05e6915dce306cb9"
	DepartmentID           *string                                            `json:"department_id,omitempty"`              // 发起审批用户部门id, 如果用户只属于一个部门, 可以不填。如果属于多个部门, 默认会选择部门列表第一个部门, 示例值: "9293493ccacbdb9a"
	Form                   ApprovalWidgetList                                 `json:"form,omitempty"`                       // json 数组, 控件值, 示例值: "[{\"id\":\"111\", \"type\": \"input\", \"value\":\"test\"}]"
	NodeApproverUserIDList []*CreateApprovalInstanceReqNodeApproverUserIDList `json:"node_approver_user_id_list,omitempty"` // 如果有发起人自选节点, 则需要填写对应节点的审批人
	NodeApproverOpenIDList []*CreateApprovalInstanceReqNodeApproverOpenIDList `json:"node_approver_open_id_list,omitempty"` // 审批人发起人自选 open id, 与上述node_approver_user_id_list字段取并集
	NodeCcUserIDList       []*CreateApprovalInstanceReqNodeCcUserIDList       `json:"node_cc_user_id_list,omitempty"`       // 如果有发起人自选节点, 则可填写对应节点的抄送人, 单个节点最多选择20位抄送人, 最大长度: `20`
	NodeCcOpenIDList       []*CreateApprovalInstanceReqNodeCcOpenIDList       `json:"node_cc_open_id_list,omitempty"`       // 抄送人发起人自选 open id 单个节点最多选择20位抄送人, 最大长度: `20`
	UUID                   *string                                            `json:"uuid,omitempty"`                       // 审批实例 uuid, 用于幂等操作, 每个租户下面的唯一key, 同一个 uuid 只能用于创建一个审批实例, 如果冲突, 返回错误码 60012, 格式建议为 XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX, 不区分大小写, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A87", 长度范围: `1` ～ `64` 字符
	AllowResubmit          *bool                                              `json:"allow_resubmit,omitempty"`             // 可配置是否可以重新提交, 适用于审批人退回场景, 提单人在同一实例重新提交单据, 示例值: true
	AllowSubmitAgain       *bool                                              `json:"allow_submit_again,omitempty"`         // 可配置是否可以再次提交, 适用于周期性提单场景, 按照当前表单内容再次发起一个新实例, 示例值: true
	CancelBotNotification  *string                                            `json:"cancel_bot_notification,omitempty"`    // 配置bot是否取消通知结果, 示例值: "0"
	ForbidRevoke           *bool                                              `json:"forbid_revoke,omitempty"`              // 配置是否可以禁止撤销, 示例值: false, 默认值: `false`
	I18nResources          []*CreateApprovalInstanceReqI18nResource           `json:"i18n_resources,omitempty"`             // 国际化文案。目前只支单行、多行文本的值。
}

// CreateApprovalInstanceReqI18nResource ...
type CreateApprovalInstanceReqI18nResource struct {
	Locale    string                                       `json:"locale,omitempty"`     // 语言可选值有: zh-CN: 中文 en-US: 英文 ja-JP: 日文, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	Texts     []*CreateApprovalInstanceReqI18nResourceText `json:"texts,omitempty"`      // 文案 key, value, i18n key 以 @i18n@ 开头； 该字段主要用于做国际化, 允许用户同时传多个语言的文案, 审批中心会根据用户当前的语音环境使用对应的文案, 如果没有传用户当前的语音环境文案, 则会使用默认的语言文案, 示例值: { "@i18n@1": "权限申请", "@i18n@2": "OA审批", "@i18n@3": "Permission" }
	IsDefault bool                                         `json:"is_default,omitempty"` // 是否默认语言, 默认语言需要包含所有key, 非默认语言如果key不存在会使用默认语言代替, 示例值: true
}

// CreateApprovalInstanceReqI18nResourceText ...
type CreateApprovalInstanceReqI18nResourceText struct {
	Key   string `json:"key,omitempty"`   // 文案key, 示例值: "@i18n@1"
	Value string `json:"value,omitempty"` // 文案, 示例值: "people"
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
