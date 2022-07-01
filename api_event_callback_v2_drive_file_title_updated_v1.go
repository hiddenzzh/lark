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

// EventV2DriveFileTitleUpdatedV1 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 文件标题变更时将触发此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-title-update
func (r *EventCallbackService) HandlerEventV2DriveFileTitleUpdatedV1(f EventV2DriveFileTitleUpdatedV1Handler) {
	r.cli.eventHandler.eventV2DriveFileTitleUpdatedV1Handler = f
}

// EventV2DriveFileTitleUpdatedV1Handler event EventV2DriveFileTitleUpdatedV1 handler
type EventV2DriveFileTitleUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2DriveFileTitleUpdatedV1) (string, error)

// EventV2DriveFileTitleUpdatedV1 ...
type EventV2DriveFileTitleUpdatedV1 struct {
}
