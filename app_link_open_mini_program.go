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
	"github.com/chyroc/lark/internal"
)

// OpenMiniProgram 打开小程序
//
// 打开一个小程序或者小程序中的一个页面
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-a-gadget
func (r *AppLinkService) OpenMiniProgram(req *OpenMiniProgramReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/mini_program/open", req)
}

// OpenMiniProgramReq ...
type OpenMiniProgramReq struct {
	AppID       string  `json:"appId,omitempty"`        // 小程序 appId(可从「开发者后台-凭证与基础信息」获取)
	Mode        *string `json:"mode,omitempty"`         // PC小程序启动模式, 枚举值包括: `sidebar-semi`: 聊天的侧边栏打开 `appCenter`: 工作台中打开 `window`: 独立大窗口打开 `window-semi`: 独立小窗口打开, 飞书3.33版本开始支持此模式
	Height      *string `json:"height,omitempty"`       // 自定义独立窗口高度（仅当`mode`为`window`时生效）, 飞书5.12版本开始支持 最小值: 480 最大值: 屏幕的高度 默认值: 飞书窗口的高度
	Width       *string `json:"width,omitempty"`        // 自定义独立窗口宽度（仅当`mode`为`window`时生效）, 飞书5.12版本开始支持 最小值: 640 最大值: 屏幕的宽度 默认值: 飞书窗口的宽度
	Relaunch    *string `json:"relaunch,omitempty"`     // 是否重新加载指定页面。该参数仅当applink中传入path参数时才会生效。枚举值包括: `false`: 与[小程序打开逻辑](https://open.feishu.cn/document/uYjL24iN/uMjNzUjLzYzM14yM2MTN#c2bcfa33)一致。如果用户已打开相同path的页面与参数, 则保持页面原先状态, 不会重新加载；其他情况下会清空原来的页面栈, 打开指定页 `true`: 无论用户是否打开相同path的页面与参数, 一定会清空原来的页面栈, 打开指定页（与[relaunch](https://open.feishu.cn/document/uYjL24iN/uEDM04SMwQjLxADN)的逻辑一致）<BR>飞书5.9版本开始支持该参数, 默认值为`false`
	Path        *string `json:"path,omitempty"`         // 需要跳转的页面路径, 路径后可以带参数。也可以使用 path_android、path_ios、path_pc 参数对不同的客户端指定不同的path
	PathAndroid *string `json:"path_android,omitempty"` // 同 path 参数, Android 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数
	PathIos     *string `json:"path_ios,omitempty"`     // 同 path 参数, iOS 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数
	PathPc      *string `json:"path_pc,omitempty"`      // 同 path 参数, PC 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数
	MinLkVer    *string `json:"min_lk_ver,omitempty"`   // 指定 AppLink 协议能够兼容的最小飞书版本, 使用三位版本号 x.y.z。如果当前飞书版本号小于min_lk_ver, 打开该 AppLink 会显示为兼容页面
}
