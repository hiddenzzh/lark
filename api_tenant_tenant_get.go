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

// GetTenant 获取企业名称、企业编号等企业信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query
func (r *TenantService) GetTenant(ctx context.Context, request *GetTenantReq, options ...MethodOptionFunc) (*GetTenantResp, *Response, error) {
	if r.cli.mock.mockTenantGetTenant != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Tenant#GetTenant mock enable")
		return r.cli.mock.mockTenantGetTenant(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Tenant",
		API:                   "GetTenant",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/tenant/v2/tenant/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getTenantResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTenantGetTenant mock TenantGetTenant method
func (r *Mock) MockTenantGetTenant(f func(ctx context.Context, request *GetTenantReq, options ...MethodOptionFunc) (*GetTenantResp, *Response, error)) {
	r.mockTenantGetTenant = f
}

// UnMockTenantGetTenant un-mock TenantGetTenant method
func (r *Mock) UnMockTenantGetTenant() {
	r.mockTenantGetTenant = nil
}

// GetTenantReq ...
type GetTenantReq struct {
}

// GetTenantResp ...
type GetTenantResp struct {
	Tenant *GetTenantRespTenant `json:"tenant,omitempty"` // 企业信息
}

// GetTenantRespTenant ...
type GetTenantRespTenant struct {
	Name      string                     `json:"name,omitempty"`       // 企业名称
	DisplayID string                     `json:"display_id,omitempty"` // 企业编号, 平台内唯一
	TenantTag int64                      `json:"tenant_tag,omitempty"` // 个人版/团队版标志, 可选值有: 0: 团队版, 2: 个人版
	TenantKey string                     `json:"tenant_key,omitempty"` // 企业标识
	Avatar    *GetTenantRespTenantAvatar `json:"avatar,omitempty"`     // 企业头像
}

// GetTenantRespTenantAvatar ...
type GetTenantRespTenantAvatar struct {
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 企业头像
	Avatar72     string `json:"avatar_72,omitempty"`     // 企业头像 72x72
	Avatar240    string `json:"avatar_240,omitempty"`    // 企业头像 240x240
	Avatar640    string `json:"avatar_640,omitempty"`    // 企业头像 640x640
}

// getTenantResp ...
type getTenantResp struct {
	Code int64          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *GetTenantResp `json:"data,omitempty"`
}
