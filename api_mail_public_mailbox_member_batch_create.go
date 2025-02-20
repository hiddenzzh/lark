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

// BatchCreatePublicMailboxMember 一次请求可以给一个公共邮箱添加多个成员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/batch_create
func (r *MailService) BatchCreatePublicMailboxMember(ctx context.Context, request *BatchCreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*BatchCreatePublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailBatchCreatePublicMailboxMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#BatchCreatePublicMailboxMember mock enable")
		return r.cli.mock.mockMailBatchCreatePublicMailboxMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "BatchCreatePublicMailboxMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchCreatePublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailBatchCreatePublicMailboxMember mock MailBatchCreatePublicMailboxMember method
func (r *Mock) MockMailBatchCreatePublicMailboxMember(f func(ctx context.Context, request *BatchCreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*BatchCreatePublicMailboxMemberResp, *Response, error)) {
	r.mockMailBatchCreatePublicMailboxMember = f
}

// UnMockMailBatchCreatePublicMailboxMember un-mock MailBatchCreatePublicMailboxMember method
func (r *Mock) UnMockMailBatchCreatePublicMailboxMember() {
	r.mockMailBatchCreatePublicMailboxMember = nil
}

// BatchCreatePublicMailboxMemberReq ...
type BatchCreatePublicMailboxMemberReq struct {
	PublicMailboxID string                                   `path:"public_mailbox_id" json:"-"` // The unique ID or email address of a public mailbox, 示例值: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx"
	UserIDType      *IDType                                  `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Items           []*BatchCreatePublicMailboxMemberReqItem `json:"items,omitempty"`            // 本次调用添加的公共邮箱成员列表, 长度范围: `1` ～ `200`
}

// BatchCreatePublicMailboxMemberReqItem ...
type BatchCreatePublicMailboxMemberReqItem struct {
	MemberID *string       `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识（在请求体中不用填）, 示例值: "xxxxxxxxxxxxxxx"
	UserID   *string       `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）, 示例值: "xxxxxxxxxx"
	Type     *MailUserType `json:"type,omitempty"`      // 成员类型, 示例值: "USER", 可选值有: USER: 内部用户
}

// BatchCreatePublicMailboxMemberResp ...
type BatchCreatePublicMailboxMemberResp struct {
	Items []*BatchCreatePublicMailboxMemberRespItem `json:"items,omitempty"` // 添加成功后的公共邮箱成员信息列表
}

// BatchCreatePublicMailboxMemberRespItem ...
type BatchCreatePublicMailboxMemberRespItem struct {
	MemberID string       `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识（在请求体中不用填）
	UserID   string       `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）
	Type     MailUserType `json:"type,omitempty"`      // 成员类型, 可选值有: USER: 内部用户
}

// batchCreatePublicMailboxMemberResp ...
type batchCreatePublicMailboxMemberResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreatePublicMailboxMemberResp `json:"data,omitempty"`
}
