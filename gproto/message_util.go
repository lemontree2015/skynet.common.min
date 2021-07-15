package gproto

import (
	"github.com/lemontree2015/skynet/misc"
)

func MessageRequestToMessageNotify(gProtoMessageRequest *GProtoMessageRequest) *GProtoMessageNotify {
	timestamp := misc.UnixTimestamp()
	return &GProtoMessageNotify{
		MsgId:       gProtoMessageRequest.MsgId,
		MsgType:     gProtoMessageRequest.MsgType,
		Content:     gProtoMessageRequest.Content,
		RoomId:      gProtoMessageRequest.RoomId,
		From:        gProtoMessageRequest.From,
		To:          gProtoMessageRequest.To,
		SendTime:    uint64(timestamp),
		SendTimeAck: uint64(timestamp),
		Extend:      gProtoMessageRequest.Extend,
	}
}
