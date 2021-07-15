package chatroom_client

import (
	"fmt"

	"github.com/lemontree2015/skynet.common.min/gproto"
	"github.com/lemontree2015/skynet/client"
)

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

//////////////////////
// RPC LeaveChatRoom
//////////////////////

type RPCGIMChatRoomLeaveChatRoomIn struct {
	AppId                      string
	ChatRoomId                 string
	Account                    string
	SessionId                  uint64
	GProtoLeaveChatRoomRequest *gproto.GProtoLeaveChatRoomRequest
}

type RPCGIMChatRoomLeaveChatRoomOut struct {
	Code string
}

func LeaveChatRoom(appId, chatRoomId, account string, sessionId uint64, gProtoLeaveChatRoomRequest *gproto.GProtoLeaveChatRoomRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomLeaveChatRoomIn{
		AppId:                      appId,
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
	AppId                string
	ChatRoomId           string
	Account              string
	SessionId            uint64
	GProtoMessageRequest *gproto.GProtoMessageRequest
}

type RPCGIMChatRoomChatRoomMessageOut struct {
	Code string
}

func ChatRoomMessage(appId, chatRoomId, account string, sessionId uint64, gProtoMessageRequest *gproto.GProtoMessageRequest) error {
	// 构造参数
	in := &RPCGIMChatRoomChatRoomMessageIn{
		AppId:                appId,
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
