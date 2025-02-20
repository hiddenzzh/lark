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

// GetChat 获取群名称、群描述、群头像、群主 ID 等群基本信息。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群里（否则只会返回群名称、群头像等基本信息）
// - 获取内部群信息时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get
func (r *ChatService) GetChat(ctx context.Context, request *GetChatReq, options ...MethodOptionFunc) (*GetChatResp, *Response, error) {
	if r.cli.mock.mockChatGetChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChat mock enable")
		return r.cli.mock.mockChatGetChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChat",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChat mock ChatGetChat method
func (r *Mock) MockChatGetChat(f func(ctx context.Context, request *GetChatReq, options ...MethodOptionFunc) (*GetChatResp, *Response, error)) {
	r.mockChatGetChat = f
}

// UnMockChatGetChat un-mock ChatGetChat method
func (r *Mock) UnMockChatGetChat() {
	r.mockChatGetChat = nil
}

// GetChatReq ...
type GetChatReq struct {
	ChatID     string  `path:"chat_id" json:"-"`       // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetChatResp ...
type GetChatResp struct {
	Avatar                 string                            `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string                            `json:"name,omitempty"`                     // 群名称
	Description            string                            `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                        `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    AddMemberPermission               `json:"add_member_permission,omitempty"`    // 群成员添加权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, 注意: 单聊不返回该字段
	ShareCardPermission    ShareCardPermission               `json:"share_card_permission,omitempty"`    // 群分享权限, 可选值有: `allowed`: 允许, `not_allowed`: 不允许, 注意: 单聊不返回该字段
	AtAllPermission        AtAllPermission                   `json:"at_all_permission,omitempty"`        // at 所有人权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, 注意: 单聊不返回该字段
	EditPermission         EditPermission                    `json:"edit_permission,omitempty"`          // 群编辑权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	OwnerIDType            IDType                            `json:"owner_id_type,omitempty"`            // 群主 ID 对应的ID类型, 与查询参数中的 [user_id_type] 相同。取值为: `open_id`、`user_id`、`union_id`其中之一, 注意: 当群主是机器人时不返回该字段, 单聊不返回该字段
	OwnerID                string                            `json:"owner_id,omitempty"`                 // 群主 ID, ID值与查询参数中的 [user_id_type] 对应；不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 注意: 当群主是机器人时不返回该字段, 单聊不返回该字段
	UserManagerIDList      []string                          `json:"user_manager_id_list,omitempty"`     // 用户管理员列表
	BotManagerIDList       []string                          `json:"bot_manager_id_list,omitempty"`      // 机器人管理员列表
	ChatMode               ChatMode                          `json:"chat_mode,omitempty"`                // 群模式, 可选值有: `group`: 群组, `topic`: 话题, `p2p`: 单聊
	ChatType               ChatType                          `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群, 注意: 单聊不返回该字段
	ChatTag                string                            `json:"chat_tag,omitempty"`                 // 群标签, 如有多个, 则按照下列顺序返回第一个, 可选值有: `inner`: 内部群, `tenant`: 公司群, `department`: 部门群, `edu`: 教育群, `meeting`: 会议群, `customer_service`: 客服群, 注意: 单聊不返回该字段
	JoinMessageVisibility  MessageVisibility                 `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 注意: 单聊不返回该字段
	LeaveMessageVisibility MessageVisibility                 `json:"leave_message_visibility,omitempty"` // 出群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 注意: 单聊不返回该字段
	MembershipApproval     MembershipApproval                `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批, 注意: 单聊不返回该字段
	ModerationPermission   ModerationPermission              `json:"moderation_permission,omitempty"`    // 发言权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, `moderator_list`: 指定群成员
	External               bool                              `json:"external,omitempty"`                 // 是否是外部群
	TenantKey              string                            `json:"tenant_key,omitempty"`               // 租户Key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用中的唯一标识
	UserCount              string                            `json:"user_count,omitempty"`               // 群成员人数
	BotCount               string                            `json:"bot_count,omitempty"`                // 群机器人数
	RestrictedModeSetting  *GetChatRespRestrictedModeSetting `json:"restricted_mode_setting,omitempty"`  // 保密模式设置
}

// GetChatRespRestrictedModeSetting ...
type GetChatRespRestrictedModeSetting struct {
	Status                         bool   `json:"status,omitempty"`                            // 保密模式是否开启
	ScreenshotHasPermissionSetting string `json:"screenshot_has_permission_setting,omitempty"` // 允许截屏录屏, 可选值有: all_members: 所有成员允许截屏录屏, not_anyone: 所有成员禁止截屏录屏
	DownloadHasPermissionSetting   string `json:"download_has_permission_setting,omitempty"`   // 允许下载消息中图片、视频和文件, 可选值有: all_members: 所有成员允许下载资源, not_anyone: 所有成员禁止下载资源
	MessageHasPermissionSetting    string `json:"message_has_permission_setting,omitempty"`    // 允许复制和转发消息, 可选值有: all_members: 所有成员允许复制和转发消息, not_anyone: 所有成员禁止复制和转发消息
}

// getChatResp ...
type getChatResp struct {
	Code int64        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string       `json:"msg,omitempty"`  // 错误描述
	Data *GetChatResp `json:"data,omitempty"`
}
