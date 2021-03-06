package chatroom_client

import (
	"fmt"
	"github.com/lemontree2015/skynet.common.min/gproto"
	"github.com/lemontree2015/skynet/client"
)

type RPCGIMChatRoomUserCountIn struct {
	RoomId string
}

type RPCGIMChatRoomUserCountOut struct {
	Code  string
	Count int
}

// 查询房间在线人数
func ChatRoomUserCount(rId string) (error, int) {
	// 构造参数
	in := &RPCGIMChatRoomUserCountIn{RoomId: rId}
	out := &RPCGIMChatRoomUserCountOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(in.RoomId), "1.0.0").Send("ChatRoomUserCount", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil, out.Count
		} else {
			return fmt.Errorf(out.Code), 0
		}
	} else {
		return err, 0
	}
}

//////////////////////
// RPC GetChatRoomInfo
//////////////////////

type RPCGIMChatRoomGetChatRoomInfoIn struct {
	AppId                        string
	ChatRoomId                   string
	Account                      string
	SessionId                    uint64
	GProtoGetChatRoomInfoRequest *gproto.GProtoGetChatRoomInfoRequest
}

type RPCGIMChatRoomGetChatRoomInfoOut struct {
	Code string
}

func GetChatRoomInfo(appId, chatRoomId, account string, sessionId uint64, gProtoGetChatRoomInfoRequest *gproto.GProtoGetChatRoomInfoRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomGetChatRoomInfoIn{
		AppId:                        appId,
		ChatRoomId:                   chatRoomId,
		Account:                      account,
		SessionId:                    sessionId,
		GProtoGetChatRoomInfoRequest: gProtoGetChatRoomInfoRequest,
	}
	out := &RPCGIMChatRoomGetChatRoomInfoOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(chatRoomId), "1.0.0").Send("GetChatRoomInfo", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////
// RPC EnterChatRoom
//////////////////////

type RPCGIMChatRoomEnterChatRoomIn struct {
	ChatRoomId                 string
	Account                    string
	SessionId                  uint64
	GProtoEnterChatRoomRequest *gproto.GProtoEnterChatRoomRequest
}

type RPCGIMChatRoomEnterChatRoomOut struct {
	Code string
}

func EnterChatRoom(chatRoomId, account string, sessionId uint64, gProtoEnterChatRoomRequest *gproto.GProtoEnterChatRoomRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomEnterChatRoomIn{
		ChatRoomId:                 chatRoomId,
		Account:                    account,
		SessionId:                  sessionId,
		GProtoEnterChatRoomRequest: gProtoEnterChatRoomRequest,
	}
	out := &RPCGIMChatRoomEnterChatRoomOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(chatRoomId), "1.0.0").Send("EnterChatRoom", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

type RPCGIMChatRoomTouchChatRoomIn struct {
	ChatRoomId                 string
	Account                    string
	SessionId                  uint64
	GProtoTouchChatRoomRequest *gproto.GProtoTouchChatRoomRequest
}

type RPCGIMChatRoomTouchChatRoomOut struct {
	Code string
}

func TouchChatRoom(chatRoomId, account string, sessionId uint64, gProtoTouchChatRoomRequest *gproto.GProtoTouchChatRoomRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomTouchChatRoomIn{
		ChatRoomId:                 chatRoomId,
		Account:                    account,
		SessionId:                  sessionId,
		GProtoTouchChatRoomRequest: gProtoTouchChatRoomRequest,
	}
	out := &RPCGIMChatRoomTouchChatRoomOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(chatRoomId), "1.0.0").Send("TouchChatRoom", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////
// RPC LeaveChatRoom
//////////////////////

type RPCGIMChatRoomLeaveChatRoomIn struct {
	ChatRoomId                 string
	Account                    string
	SessionId                  uint64
	GProtoLeaveChatRoomRequest *gproto.GProtoLeaveChatRoomRequest
}

type RPCGIMChatRoomLeaveChatRoomOut struct {
	Code string
}

func LeaveChatRoom(chatRoomId, account string, sessionId uint64, gProtoLeaveChatRoomRequest *gproto.GProtoLeaveChatRoomRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomLeaveChatRoomIn{
		ChatRoomId:                 chatRoomId,
		Account:                    account,
		SessionId:                  sessionId,
		GProtoLeaveChatRoomRequest: gProtoLeaveChatRoomRequest,
	}
	out := &RPCGIMChatRoomLeaveChatRoomOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(chatRoomId), "1.0.0").Send("LeaveChatRoom", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////
// RPC ChatRoomMessage
//////////////////////

type RPCGIMChatRoomChatRoomMessageIn struct {
	ChatRoomId           string
	Account              string
	SessionId            uint64
	GProtoMessageRequest *gproto.GProtoMessageRequest
}

type RPCGIMChatRoomChatRoomMessageOut struct {
	Code string
}

func ChatRoomMessage(chatRoomId, account string, sessionId uint64, gProtoMessageRequest *gproto.GProtoMessageRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomChatRoomMessageIn{
		ChatRoomId:           chatRoomId,
		Account:              account,
		SessionId:            sessionId,
		GProtoMessageRequest: gProtoMessageRequest,
	}
	out := &RPCGIMChatRoomChatRoomMessageOut{}

	// RPC请求
	if err := client.GetClient(ChatRoomServiceName(chatRoomId), "1.0.0").Send("ChatRoomMessage", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}
