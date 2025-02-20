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

// ClearPublicMailboxMember 删除公共邮箱所有成员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear
func (r *MailService) ClearPublicMailboxMember(ctx context.Context, request *ClearPublicMailboxMemberReq, options ...MethodOptionFunc) (*ClearPublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailClearPublicMailboxMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#ClearPublicMailboxMember mock enable")
		return r.cli.mock.mockMailClearPublicMailboxMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "ClearPublicMailboxMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(clearPublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailClearPublicMailboxMember mock MailClearPublicMailboxMember method
func (r *Mock) MockMailClearPublicMailboxMember(f func(ctx context.Context, request *ClearPublicMailboxMemberReq, options ...MethodOptionFunc) (*ClearPublicMailboxMemberResp, *Response, error)) {
	r.mockMailClearPublicMailboxMember = f
}

// UnMockMailClearPublicMailboxMember un-mock MailClearPublicMailboxMember method
func (r *Mock) UnMockMailClearPublicMailboxMember() {
	r.mockMailClearPublicMailboxMember = nil
}

// ClearPublicMailboxMemberReq ...
type ClearPublicMailboxMemberReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值: "xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
}

// ClearPublicMailboxMemberResp ...
type ClearPublicMailboxMemberResp struct {
}

// clearPublicMailboxMemberResp ...
type clearPublicMailboxMemberResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *ClearPublicMailboxMemberResp `json:"data,omitempty"`
}
