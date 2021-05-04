package lark

import (
	"context"
)

// DeleteMessage 机器人撤回机器人自己发送的消息或群主撤回群内消息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)  ，撤回消息时机器人仍需要在会话内
// - 机器人可以撤回单聊和群组内，自己发送 且 发送时间不超过1天(24小时)的消息
// - 若机器人要撤回群内他人发送的消息，则机器人必须是该群的群主 或者 得到群主的授权，且消息发送时间不超过1天（24小时）
// - 无法撤回通过「批量发送消息接口」发送的消息
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete
func (r *MessageAPI) DeleteMessage(ctx context.Context, request *DeleteMessageReq) (*DeleteMessageResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(deleteMessageResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "DeleteMessage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待撤回的消息的ID,**示例值**："om_dc13264520392913993dd051dba21dcf"
}

type deleteMessageResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMessageResp `json:"data,omitempty"`
}

type DeleteMessageResp struct{}
