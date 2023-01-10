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
	"io"
)

// UploadAdminBadgeImage 通过该接口可以上传勋章详情图、挂饰图的文件, 获取对应的文件key。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge_image/create
func (r *AdminService) UploadAdminBadgeImage(ctx context.Context, request *UploadAdminBadgeImageReq, options ...MethodOptionFunc) (*UploadAdminBadgeImageResp, *Response, error) {
	if r.cli.mock.mockAdminUploadAdminBadgeImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#UploadAdminBadgeImage mock enable")
		return r.cli.mock.mockAdminUploadAdminBadgeImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "UploadAdminBadgeImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badge_images",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadAdminBadgeImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminUploadAdminBadgeImage mock AdminUploadAdminBadgeImage method
func (r *Mock) MockAdminUploadAdminBadgeImage(f func(ctx context.Context, request *UploadAdminBadgeImageReq, options ...MethodOptionFunc) (*UploadAdminBadgeImageResp, *Response, error)) {
	r.mockAdminUploadAdminBadgeImage = f
}

// UnMockAdminUploadAdminBadgeImage un-mock AdminUploadAdminBadgeImage method
func (r *Mock) UnMockAdminUploadAdminBadgeImage() {
	r.mockAdminUploadAdminBadgeImage = nil
}

// UploadAdminBadgeImageReq ...
type UploadAdminBadgeImageReq struct {
	ImageFile io.Reader `json:"image_file,omitempty"` // 勋章图片的文件, 仅支持 PNG 格式, 320 x 320 像素, 大小不超过 1024 KB, 示例值: file binary
	ImageType int64     `json:"image_type,omitempty"` // 图片的类型, 示例值: 1, 可选值有: 1: 勋章详情图, 2: 勋章挂饰图, 取值范围: `1` ～ `2`
}

// UploadAdminBadgeImageResp ...
type UploadAdminBadgeImageResp struct {
	ImageKey string `json:"image_key,omitempty"` // 图片的key
}

// uploadAdminBadgeImageResp ...
type uploadAdminBadgeImageResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *UploadAdminBadgeImageResp `json:"data,omitempty"`
}
