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

// GetJssdkTicket 通过在你的网页中引入飞书网页组件, 可在你的网页中直接使用飞书的功能。
//
// - 网页组件只适用于自建应用, 暂不支持商店应用。
// - 网页组件适用于普通web网页, 不建议在小程序webview中调用此组件
// ## 准备工作
// - 在开发者后台[创建一个企业自建应用](https://open.feishu.cn/document/home/introduction-to-custom-app-development/self-built-application-development-process)。
// - 引入组件库。在网页 html 中引入如下代码:
// ```html
// <script src="https://lf1-cdn-tos.bytegoofy.com/goofy/locl/lark/external_js_sdk/h5-js-sdk-1.1.2.js"></script>
// ```
// ::: warning
// 若要使用成员卡片组件, SDK需要在`<body>`加载后引入。
//
// doc: https://open.feishu.cn/document/uYjL24iN/uUDO3YjL1gzN24SN4cjN
func (r *JssdkService) GetJssdkTicket(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error) {
	if r.cli.mock.mockJssdkGetJssdkTicket != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Jssdk#GetJssdkTicket mock enable")
		return r.cli.mock.mockJssdkGetJssdkTicket(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Jssdk",
		API:                 "GetJssdkTicket",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/jssdk/ticket/get",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedAppAccessToken:  true,
		NeedUserAccessToken: true,
	}
	resp := new(getJssdkTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockJssdkGetJssdkTicket mock JssdkGetJssdkTicket method
func (r *Mock) MockJssdkGetJssdkTicket(f func(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error)) {
	r.mockJssdkGetJssdkTicket = f
}

// UnMockJssdkGetJssdkTicket un-mock JssdkGetJssdkTicket method
func (r *Mock) UnMockJssdkGetJssdkTicket() {
	r.mockJssdkGetJssdkTicket = nil
}

// GetJssdkTicketReq ...
type GetJssdkTicketReq struct {
}

// GetJssdkTicketResp ...
type GetJssdkTicketResp struct {
	ExpireIn int64  `json:"expire_in,omitempty"` // jsapi_ticket 的有效时间
	Ticket   string `json:"ticket,omitempty"`    // jsapi_ticket
}

// getJssdkTicketResp ...
type getJssdkTicketResp struct {
	Code int64               `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 对返回码的文本描述
	Data *GetJssdkTicketResp `json:"data,omitempty"` // 返回内容
}
