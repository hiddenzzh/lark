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

// GetUserListOld 基于部门ID获取部门下直属用户列表。
//
// [常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。
// - 使用 user_access_token 情况下根据个人组织架构的通讯录可见范围进行权限过滤, 返回个人组织架构通讯录范围（[登陆企业管理后台进行权限配置](https://www.feishu.cn/admin/security/permission/visibility)）内可见的用户数据。
// -  tenant_access_token  基于应用通讯录范围进行权限鉴定。由于 department_id 是非必填参数, 填与不填存在<b>两种数据权限校验与返回</b>情况: 1、请求设置了 department_id
// （根部门为0）, 会检验所带部门ID是否具有通讯录权限（如果带上
// department_id=0 会校验是否有全员权限）, 有则返回部门下直属的成员列表, 否则提示无部门权限的错误码返回。 2、请求未带
// department_id 参数, 则会返回权限范围内的独立用户（权限范围直接包含了某用户, 则该用户视为权限范围内的独立用户）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list
//
// Deprecated
func (r *ContactService) GetUserListOld(ctx context.Context, request *GetUserListOldReq, options ...MethodOptionFunc) (*GetUserListOldResp, *Response, error) {
	if r.cli.mock.mockContactGetUserListOld != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetUserListOld mock enable")
		return r.cli.mock.mockContactGetUserListOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetUserListOld",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/users",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getUserListOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetUserListOld mock ContactGetUserListOld method
func (r *Mock) MockContactGetUserListOld(f func(ctx context.Context, request *GetUserListOldReq, options ...MethodOptionFunc) (*GetUserListOldResp, *Response, error)) {
	r.mockContactGetUserListOld = f
}

// UnMockContactGetUserListOld un-mock ContactGetUserListOld method
func (r *Mock) UnMockContactGetUserListOld() {
	r.mockContactGetUserListOld = nil
}

// GetUserListOldReq ...
type GetUserListOldReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_type", 可选值有: `department_id`: 以自定义department_id来标识部门, `open_department_id`: 以open_department_id来标识部门, 默认值: `open_department_id`
	DepartmentID     *string           `query:"department_id" json:"-"`      // 填写该字段表示获取部门下所有用户, 选填, 示例值: "od-xxxxxxxxxxxxx"
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS%2BJKiSIkdexPw="
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 最大值: `100`
}

// GetUserListOldResp ...
type GetUserListOldResp struct {
	HasMore   bool                      `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                    `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetUserListOldRespItem `json:"items,omitempty"`
}

// GetUserListOldRespItem ...
type GetUserListOldRespItem struct {
	UnionID         string                              `json:"union_id,omitempty"`          // 用户的union_id, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID          string                              `json:"user_id,omitempty"`           // 租户内用户的唯一标识, 用户的user_id, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	OpenID          string                              `json:"open_id,omitempty"`           // 用户的open_id, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Name            string                              `json:"name,omitempty"`              // 用户名, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	EnName          string                              `json:"en_name,omitempty"`           // 英文名, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	Email           string                              `json:"email,omitempty"`             // 邮箱, 注意: 1. 非中国大陆手机号成员必须同时添加邮箱, 2. 邮箱不可重复, 字段权限要求: 获取用户邮箱信息
	Mobile          string                              `json:"mobile,omitempty"`            // 手机号, 在本企业内不可重复；未认证企业仅支持添加中国大陆手机号, 通过飞书认证的企业允许添加海外手机号, 注意国际电话区号前缀中必须包含加号 +, 字段权限要求: 获取用户手机号
	MobileVisible   bool                                `json:"mobile_visible,omitempty"`    // 手机号码可见性, true 为可见, false 为不可见, 目前默认为 true。不可见时, 组织员工将无法查看该员工的手机号码
	Gender          int64                               `json:"gender,omitempty"`            // 性别, 可选值有: `0`: 保密, `1`: 男, `2`: 女, 字段权限要求（满足任一）: 获取用户性别, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	Avatar          *GetUserListOldRespItemAvatar       `json:"avatar,omitempty"`            // 用户头像信息, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	Status          *GetUserListOldRespItemStatus       `json:"status,omitempty"`            // 用户状态, 枚举类型, 包括is_frozen、is_resigned、is_activated、is_exited, 用户状态转移参见: [用户状态图](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/field-overview#4302b5a1), 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	DepartmentIDs   []string                            `json:"department_ids,omitempty"`    // 用户所属部门的ID列表, 一个用户可属于多个部门, ID值的类型与查询参数中的department_id_type 对应, 不同 ID 的说明与department_id的获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	LeaderUserID    string                              `json:"leader_user_id,omitempty"`    // 用户的直接主管的用户ID, ID值与查询参数中的user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 获取方式参见[如何获取user_id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	City            string                              `json:"city,omitempty"`              // 工作城市, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	Country         string                              `json:"country,omitempty"`           // 国家或地区Code缩写, 具体写入格式请参考 [国家/地区码表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description), 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	WorkStation     string                              `json:"work_station,omitempty"`      // 工位, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	JoinTime        int64                               `json:"join_time,omitempty"`         // 入职时间, 时间戳格式, 表示从1970年1月1日开始所经过的秒数, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	IsTenantManager bool                                `json:"is_tenant_manager,omitempty"` // 是否是租户超级管理员, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	EmployeeNo      string                              `json:"employee_no,omitempty"`       // 工号, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	EmployeeType    int64                               `json:"employee_type,omitempty"`     // 员工类型, 可选值有: `1`: 正式员工, `2`: 实习生, `3`: 外包, `4`: 劳务, `5`: 顾问, 同时可读取到自定义员工类型的 int 值, 可通过下方接口获取到该租户的自定义员工类型的名称, [获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	Orders          []*GetUserListOldRespItemOrder      `json:"orders,omitempty"`            // 用户排序信息, 用于标记通讯录下组织架构的人员顺序, 人员可能存在多个部门中, 且有不同的排序, 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	CustomAttrs     []*GetUserListOldRespItemCustomAttr `json:"custom_attrs,omitempty"`      // 自定义字段, 请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“, 否则该字段不会生效/返回, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	EnterpriseEmail string                              `json:"enterprise_email,omitempty"`  // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
	JobTitle        string                              `json:"job_title,omitempty"`         // 职务, 字段权限要求（满足任一）: 获取用户雇佣信息, 以应用身份读取通讯录, 读取通讯录, 以应用身份访问通讯录
}

// GetUserListOldRespItemAvatar ...
type GetUserListOldRespItemAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// GetUserListOldRespItemCustomAttr ...
type GetUserListOldRespItemCustomAttr struct {
	Type  string                                 `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, [自定义字段相关常见问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                                 `json:"id,omitempty"`    // 自定义字段ID
	Value *GetUserListOldRespItemCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// GetUserListOldRespItemCustomAttrValue ...
type GetUserListOldRespItemCustomAttrValue struct {
	Text        string                                            `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                            `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                            `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionValue string                                            `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                            `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                            `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *GetUserListOldRespItemCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// GetUserListOldRespItemCustomAttrValueGenericUser ...
type GetUserListOldRespItemCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型    1: 用户
}

// GetUserListOldRespItemOrder ...
type GetUserListOldRespItemOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, ID值与查询参数中的department_id_type 对应, 表示用户所在的、且需要排序的部门, 不同 ID 的说明参见及获取方式 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview)
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序, 数值越大, 排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序, 数值越大, 排序越靠前
}

// GetUserListOldRespItemStatus ...
type GetUserListOldRespItemStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

// getUserListOldResp ...
type getUserListOldResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetUserListOldResp `json:"data,omitempty"`
}
