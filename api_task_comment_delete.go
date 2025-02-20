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

// DeleteTaskComment 该接口用于通过评论ID删除评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/delete
func (r *TaskService) DeleteTaskComment(ctx context.Context, request *DeleteTaskCommentReq, options ...MethodOptionFunc) (*DeleteTaskCommentResp, *Response, error) {
	if r.cli.mock.mockTaskDeleteTaskComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#DeleteTaskComment mock enable")
		return r.cli.mock.mockTaskDeleteTaskComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "DeleteTaskComment",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteTaskCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskDeleteTaskComment mock TaskDeleteTaskComment method
func (r *Mock) MockTaskDeleteTaskComment(f func(ctx context.Context, request *DeleteTaskCommentReq, options ...MethodOptionFunc) (*DeleteTaskCommentResp, *Response, error)) {
	r.mockTaskDeleteTaskComment = f
}

// UnMockTaskDeleteTaskComment un-mock TaskDeleteTaskComment method
func (r *Mock) UnMockTaskDeleteTaskComment() {
	r.mockTaskDeleteTaskComment = nil
}

// DeleteTaskCommentReq ...
type DeleteTaskCommentReq struct {
	TaskID    string `path:"task_id" json:"-"`    // 任务ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	CommentID string `path:"comment_id" json:"-"` // 评论ID, 示例值: "6937231762296684564"
}

// DeleteTaskCommentResp ...
type DeleteTaskCommentResp struct {
}

// deleteTaskCommentResp ...
type deleteTaskCommentResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteTaskCommentResp `json:"data,omitempty"`
}
