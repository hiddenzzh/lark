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

// EventV2DriveFileBitableRecordChangedV1 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 被订阅的多维表格记录发生变更将会触发此事件, 公式字段的值变化不会触发事件。
// ## 概述
// :::html
// <md-table>
// <md-thead>
// <tr>
// <md-th>基本</md-th>
// <md-th></md-th>
// </tr>
// </md-thead>
// <md-tbody>
// <md-tr>
// <md-th>支持的应用类型</md-th>
// <md-td>
// custom, isv
// </md-td>
// </md-tr>
// <md-tr>
// <md-th>
// 权限要求
// </md-th>
// <md-td>
// 查看、评论、编辑和管理多维表格
// 查看、评论、编辑和管理云空间中所有文件
// </md-td>
// </md-tr>
// </md-tbody>
// </md-table>
// 如何订阅文档请点击查看 [订阅云文档事件](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/subscribe)。
// ## 支持的记录变更类型
// | 变更类型         | action           |
// | --------- | --------------- |
// |新增行记录 | `record_added` |
// |删除行记录 | `record_deleted` |
// |修改行记录 | `record_edited` |
// 回调结构中的 `field_value` 字段为 JSON 序列化后的字符串, 序列化前的结构请查看 [数据结构](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/development-guide/bitable-structure)
// ## 回调示例
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/event/list/bitable-record-changed
func (r *EventCallbackService) HandlerEventV2DriveFileBitableRecordChangedV1(f EventV2DriveFileBitableRecordChangedV1Handler) {
	r.cli.eventHandler.eventV2DriveFileBitableRecordChangedV1Handler = f
}

// EventV2DriveFileBitableRecordChangedV1Handler event EventV2DriveFileBitableRecordChangedV1 handler
type EventV2DriveFileBitableRecordChangedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2DriveFileBitableRecordChangedV1) (string, error)

// EventV2DriveFileBitableRecordChangedV1 ...
type EventV2DriveFileBitableRecordChangedV1 struct {
}
