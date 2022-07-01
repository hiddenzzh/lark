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

// GetApplicationAppAdminUserList 查询审核应用的管理员列表, 返回最新10个管理员账户id列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN
func (r *ApplicationService) GetApplicationAppAdminUserList(ctx context.Context, request *GetApplicationAppAdminUserListReq, options ...MethodOptionFunc) (*GetApplicationAppAdminUserListResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationAppAdminUserList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationAppAdminUserList mock enable")
		return r.cli.mock.mockApplicationGetApplicationAppAdminUserList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationAppAdminUserList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/user/v4/app_admin_user/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationAppAdminUserListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationAppAdminUserList mock ApplicationGetApplicationAppAdminUserList method
func (r *Mock) MockApplicationGetApplicationAppAdminUserList(f func(ctx context.Context, request *GetApplicationAppAdminUserListReq, options ...MethodOptionFunc) (*GetApplicationAppAdminUserListResp, *Response, error)) {
	r.mockApplicationGetApplicationAppAdminUserList = f
}

// UnMockApplicationGetApplicationAppAdminUserList un-mock ApplicationGetApplicationAppAdminUserList method
func (r *Mock) UnMockApplicationGetApplicationAppAdminUserList() {
	r.mockApplicationGetApplicationAppAdminUserList = nil
}

// GetApplicationAppAdminUserListReq ...
type GetApplicationAppAdminUserListReq struct {
}

// GetApplicationAppAdminUserListResp ...
type GetApplicationAppAdminUserListResp struct {
	UserList []*GetApplicationAppAdminUserListRespUser `json:"user_list,omitempty"` // 管理员列表
}

// GetApplicationAppAdminUserListRespUser ...
type GetApplicationAppAdminUserListRespUser struct {
	OpenID  string `json:"open_id,omitempty"`  // 某管理员的open_id
	UserID  string `json:"user_id,omitempty"`  // 某管理员的user_id
	UnionID string `json:"union_id,omitempty"` // 某管理员的union_id
}

// getApplicationAppAdminUserListResp ...
type getApplicationAppAdminUserListResp struct {
	Code int64                               `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 返回码描述
	Data *GetApplicationAppAdminUserListResp `json:"data,omitempty"`
}
