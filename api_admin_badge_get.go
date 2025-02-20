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

// GetAdminBadge 可以通过该接口查询勋章的详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge/get
func (r *AdminService) GetAdminBadge(ctx context.Context, request *GetAdminBadgeReq, options ...MethodOptionFunc) (*GetAdminBadgeResp, *Response, error) {
	if r.cli.mock.mockAdminGetAdminBadge != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#GetAdminBadge mock enable")
		return r.cli.mock.mockAdminGetAdminBadge(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "GetAdminBadge",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badges/:badge_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAdminBadgeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminGetAdminBadge mock AdminGetAdminBadge method
func (r *Mock) MockAdminGetAdminBadge(f func(ctx context.Context, request *GetAdminBadgeReq, options ...MethodOptionFunc) (*GetAdminBadgeResp, *Response, error)) {
	r.mockAdminGetAdminBadge = f
}

// UnMockAdminGetAdminBadge un-mock AdminGetAdminBadge method
func (r *Mock) UnMockAdminGetAdminBadge() {
	r.mockAdminGetAdminBadge = nil
}

// GetAdminBadgeReq ...
type GetAdminBadgeReq struct {
	BadgeID string `path:"badge_id" json:"-"` // 勋章id, 示例值: "m_DjMzaK", 长度范围: `1` ～ `64` 字符
}

// GetAdminBadgeResp ...
type GetAdminBadgeResp struct {
	Badge *GetAdminBadgeRespBadge `json:"badge,omitempty"` // 勋章信息
}

// GetAdminBadgeRespBadge ...
type GetAdminBadgeRespBadge struct {
	ID              string                                 `json:"id,omitempty"`               // 租户内勋章的唯一标识, 该值由系统随机生成。
	Name            string                                 `json:"name,omitempty"`             // 租户内唯一的勋章名称, 最多30个字符。
	Explanation     string                                 `json:"explanation,omitempty"`      // 勋章的描述文案, 最多100个字符。
	DetailImage     string                                 `json:"detail_image,omitempty"`     // 企业勋章的详情图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key。
	ShowImage       string                                 `json:"show_image,omitempty"`       // 企业勋章的头像挂饰图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key。
	I18nName        *GetAdminBadgeRespBadgeI18nName        `json:"i18n_name,omitempty"`        // 勋章的多语言名称, 同name字段限制, 最多30个字符。
	I18nExplanation *GetAdminBadgeRespBadgeI18nExplanation `json:"i18n_explanation,omitempty"` // 勋章的多语言描述文案, 同explanation字段限制, 最多100个字符。
}

// GetAdminBadgeRespBadgeI18nExplanation ...
type GetAdminBadgeRespBadgeI18nExplanation struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文文案
	EnUs string `json:"en_us,omitempty"` // 英文文案
	JaJp string `json:"ja_jp,omitempty"` // 日文文案
}

// GetAdminBadgeRespBadgeI18nName ...
type GetAdminBadgeRespBadgeI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文文案
	EnUs string `json:"en_us,omitempty"` // 英文文案
	JaJp string `json:"ja_jp,omitempty"` // 日文文案
}

// getAdminBadgeResp ...
type getAdminBadgeResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetAdminBadgeResp `json:"data,omitempty"`
}
