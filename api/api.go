package api

import (
	"github.com/golang/glog"
	"github.com/lemontree2015/skynet"
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
