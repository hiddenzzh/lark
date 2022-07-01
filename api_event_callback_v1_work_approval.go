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

// EventV1WorkApproval 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 「审批」应用的表单里如果包含 [加班控件组], 则在此表单审批通过后触发此事件。
// * 依赖权限: [访问审批应用]
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/overtime
func (r *EventCallbackService) HandlerEventV1WorkApproval(f EventV1WorkApprovalHandler) {
	r.cli.eventHandler.eventV1WorkApprovalHandler = f
}

// EventV1WorkApprovalHandler event EventV1WorkApproval handler
type EventV1WorkApprovalHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1WorkApproval) (string, error)

// EventV1WorkApproval ...
type EventV1WorkApproval struct {
	AppID         string `json:"app_id,omitempty"`          // 如: cli_xxx
	TenantKey     string `json:"tenant_key,omitempty"`      // 如: xxx
	Type          string `json:"type,omitempty"`            // 如: work_approval
	InstanceCode  string `json:"instance_code,omitempty"`   // 审批实例Code. 如: xxx
	EmployeeID    string `json:"employee_id,omitempty"`     // 用户id. 如: xxx
	OpenID        string `json:"open_id,omitempty"`         // 用户open_id. 如: ou_xxx
	StartTime     int64  `json:"start_time,omitempty"`      // 审批发起时间. 如: 1502199207
	EndTime       int64  `json:"end_time,omitempty"`        // 审批结束时间. 如: 1502199307
	WorkType      string `json:"work_type,omitempty"`       // 加班类型. 如: xxx
	WorkStartTime string `json:"work_start_time,omitempty"` // 加班开始时间. 如: 2018-12-01 12:00:00
	WorkEndTime   string `json:"work_end_time,omitempty"`   // 加班结束时间. 如: 2018-12-02 12:00:00
	WorkInterval  int64  `json:"work_interval,omitempty"`   // 加班时长, 单位（秒）. 如: 7200
	WorkReason    string `json:"work_reason,omitempty"`     // 加班事由. 如: xxx
}
