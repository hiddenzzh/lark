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

// EventV1AppOpen 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 当租户第一次安装并启用此应用时触发此事件, 启用应用包含以下场景:
// - 当租户管理员后台首次开通应用
// - 租户内的普通成员首次安装此应用
// 只有应用商店应用才能订阅此事件。自建应用无此事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app-first-enabled
func (r *EventCallbackService) HandlerEventV1AppOpen(f EventV1AppOpenHandler) {
	r.cli.eventHandler.eventV1AppOpenHandler = f
}

// EventV1AppOpenHandler event EventV1AppOpen handler
type EventV1AppOpenHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1AppOpen) (string, error)

// EventV1AppOpen ...
type EventV1AppOpen struct {
	AppID             string                                `json:"app_id,omitempty"`             // 开通的应用ID. 如: cli_xxx
	TenantKey         string                                `json:"tenant_key,omitempty"`         // 开通应用的企业唯一标识. 如: xxx
	Type              string                                `json:"type,omitempty"`               // 事件类型. 如: app_open
	Applicants        []*EventV1AppOpenEventApplicant       `json:"applicants,omitempty"`         // 应用的申请者, 可能有多个
	Installer         *EventV1AppOpenEventInstaller         `json:"installer,omitempty"`          // 当应用被管理员安装时, 返回此字段。如果是自动安装或由普通成员获取时, 没有此字段
	InstallerEmployee *EventV1AppOpenEventInstallerEmployee `json:"installer_employee,omitempty"` // 当应用被普通成员安装时, 返回此字段
}

// EventV1AppOpenEventApplicant ...
type EventV1AppOpenEventApplicant struct {
	OpenID string `json:"open_id,omitempty"` // 用户对此应用的唯一标识, 同一用户对不同应用的open_id不同. 如: xxx
}

// EventV1AppOpenEventInstaller ...
type EventV1AppOpenEventInstaller struct {
	OpenID string `json:"open_id,omitempty"` // 用户对此应用的唯一标识, 同一用户对不同应用的open_id不同. 如: xxx
}

// EventV1AppOpenEventInstallerEmployee ...
type EventV1AppOpenEventInstallerEmployee struct {
	OpenID string `json:"open_id,omitempty"` // 用户对此应用的唯一标识, 同一用户对不同应用的open_id不同. 如: xxx
}
