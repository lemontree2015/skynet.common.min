package api

import (
	"github.com/golang/glog"
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet.common.min/chatroom_client"
	"github.com/lemontree2015/skynet.common.min/errors"
	"github.com/lemontree2015/skynet.common.min/gproto"
	"github.com/lemontree2015/skynet.common.min/server_client"
	"github.com/lemontree2015/skynet.common.min/session_client"
)

/////////////////
// Server APIs
/////////////////

// Route一条KickoffNotify消息到目标GIM Server[N] Service
func RemoteRouteKickoffNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoKickOffNotify *gproto.GProtoKickOffNotify) error {
	return server_client.KickoffNotify(serviceKey, account, sessionId, gProtoKickOffNotify)
}

// 查询在线用户到目标GIM Server[N] Service
func RemoteGetAndDelMessage(serviceKey *skynet.ServiceKey, account string) (error, map[string]*gproto.GProtoMessageNotify) {
	return server_client.GetAndDelMessage(serviceKey, account)
}

// Route一条MessageNotify消息到目标GIM Server[N] Service
//
// 备注:
// 如果目标Session不存在, 则将消息写入DB
func RouteMessageNotify(account string, gProtoMessageNotify *gproto.GProtoMessageNotify, isSave bool) error {
	if serviceKey, sessionId, _ := GetSession(account); serviceKey != nil {
		// 找到目标account对应的session
		return RemoteRouteMessageNotify(serviceKey, account, sessionId, gProtoMessageNotify)
	} else {
		if isSave {
			// 写入DB
			glog.Infof("APIMessage(OfflineMessage) create offline_error: account=%v, msgId=%v",
				account, gProtoMessageNotify.MsgId)
		}
		return nil
	}
}

// Route一条MessageNotify消息到目标GIM Server[N] Service
func RemoteRouteMessageNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoMessageNotify *gproto.GProtoMessageNotify) error {
	return server_client.MessageNotify(serviceKey, account, sessionId, gProtoMessageNotify)
}

// Route一条EnterChatRoomResponse消息到目标GIM Server[N] Service
func RouteEnterChatRoomResponse(account string, sessionId uint64, gProtoEnterChatRoomResponse *gproto.GProtoEnterChatRoomResponse) error {
	if serviceKey, _, _ := GetSession(account); serviceKey != nil {
		// 找到目标account对应的session
		return RemoteRouteEnterChatRoomResponse(serviceKey, account, sessionId, gProtoEnterChatRoomResponse)
	} else {
		return errors.NoSessionError
	}
}

// Route一条EnterChatRoomResponse消息到目标GIM Server[N] Service
func RemoteRouteEnterChatRoomResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoEnterChatRoomResponse *gproto.GProtoEnterChatRoomResponse) error {
	return server_client.EnterChatRoomResponse(serviceKey, account, sessionId, gProtoEnterChatRoomResponse)
}

// Route一条EnterChatRoomNotify消息到目标GIM Server[N] Service
func RemoteRouteEnterChatRoomNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoEnterChatRoomNotify *gproto.GProtoEnterChatRoomNotify) error {
	return server_client.EnterChatRoomNotify(serviceKey, account, sessionId, gProtoEnterChatRoomNotify)
}

// Route一条LeaveChatRoomResponse消息到目标GIM Server[N] Service
func RouteLeaveChatRoomResponse(account string, sessionId uint64, gProtoLeaveChatRoomResponse *gproto.GProtoLeaveChatRoomResponse) error {
	if serviceKey, _, _ := GetSession(account); serviceKey != nil {
		// 找到目标account对应的session
		return RemoteRouteLeaveChatRoomResponse(serviceKey, account, sessionId, gProtoLeaveChatRoomResponse)
	} else {
		return errors.NoSessionError
	}
}

// Route一条LeaveChatRoomResponse消息到目标GIM Server[N] Service
func RemoteRouteLeaveChatRoomResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoLeaveChatRoomResponse *gproto.GProtoLeaveChatRoomResponse) error {
	return server_client.LeaveChatRoomResponse(serviceKey, account, sessionId, gProtoLeaveChatRoomResponse)
}

// Route一条LeaveChatRoomNotify消息到目标GIM Server[N] Service
func RemoteRouteLeaveChatRoomNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoLeaveChatRoomNotify *gproto.GProtoLeaveChatRoomNotify) error {
	return server_client.LeaveChatRoomNotify(serviceKey, account, sessionId, gProtoLeaveChatRoomNotify)
}

// Route一条MessageResponse消息到目标GIM Server[N] Service
//
// 备注:
// 如果目标Session不存在, 则将消息写入DB
func RouteMessageResponse(account string, sessionId uint64, gProtoMessageResponse *gproto.GProtoMessageResponse) error {
	if serviceKey, _, _ := GetSession(account); serviceKey != nil {
		// 找到目标account对应的session
		return RemoteRouteMessageResponse(serviceKey, account, sessionId, gProtoMessageResponse)
	} else {
		return errors.NoSessionError
	}
}

// Route一条MessageResponse消息到目标GIM Server[N] Service
func RemoteRouteMessageResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoMessageResponse *gproto.GProtoMessageResponse) error {
	return server_client.MessageResponse(serviceKey, account, sessionId, gProtoMessageResponse)
}

/////////////////
// Session APIs
/////////////////

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有对应的Session存在
func TouchSession(account string) (*skynet.ServiceKey, uint64, error) {
	return session_client.TouchSession(account)
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有对应的Session存在
func GetSession(account string) (*skynet.ServiceKey, uint64, error) {
	return session_client.GetSession(account)
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有OldServiceKey
func SetSession(account string, serviceKey *skynet.ServiceKey, remoteSessionId uint64) (*skynet.ServiceKey, uint64, error) {
	return session_client.SetSession(account, serviceKey, remoteSessionId)
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有OldServiceKey
func DelSession(account string) (*skynet.ServiceKey, uint64, error) {
	return session_client.DelSession(account)
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否删除了OldServiceKey
func DelSessionWithSessionId(account string, sessionId uint64) (*skynet.ServiceKey, uint64, error) {
	return session_client.DelSessionWithSessionId(account, sessionId)
}

/////////////////
// ChatRoom APIs
/////////////////

func EnterChatRoom(chatRoomId, account string, sessionId uint64, gProtoEnterChatRoomRequest *gproto.GProtoEnterChatRoomRequest) error {
	return chatroom_client.EnterChatRoom(chatRoomId, account, sessionId, gProtoEnterChatRoomRequest)
}

func LeaveChatRoom(chatRoomId, account string, sessionId uint64, gProtoLeaveChatRoomRequest *gproto.GProtoLeaveChatRoomRequest) error {
	return chatroom_client.LeaveChatRoom(chatRoomId, account, sessionId, gProtoLeaveChatRoomRequest)
}

func ChatRoomMessage(chatRoomId, account string, sessionId uint64, gProtoMessageRequest *gproto.GProtoMessageRequest) error {
	return chatroom_client.ChatRoomMessage(chatRoomId, account, sessionId, gProtoMessageRequest)
}
