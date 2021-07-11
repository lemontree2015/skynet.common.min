package session_client

import (
	"fmt"
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet/client"
)

////////////////////
// RPC TouchSession
////////////////////


type RPCGIMSessionTouchSessionIn struct {
	Account string
}

type RPCGIMSessionTouchSessionOut struct {
	Code            string
	Account         string
	ServiceKey      *skynet.ServiceKey
	RemoteSessionId uint64
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有对应的Session存在
func TouchSession(account string) (*skynet.ServiceKey, uint64, error) {
	// 构造参数
	in := &RPCGIMSessionTouchSessionIn{
		Account: account,
	}
	out := &RPCGIMSessionTouchSessionOut{}

	// RPC请求
	if err := client.GetClient(SessionServiceName(account), "1.0.0").Send("TouchSession", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return out.ServiceKey, out.RemoteSessionId, nil
		} else {
			return out.ServiceKey, out.RemoteSessionId, fmt.Errorf(out.Code)
		}
	} else {
		return nil, 0, err
	}
}

//////////////////
// RPC GetSession
//////////////////

type RPCGIMSessionGetSessionIn struct {
	Account string
}

type RPCGIMSessionGetSessionOut struct {
	Code            string
	Account         string
	ServiceKey      *skynet.ServiceKey
	RemoteSessionId uint64
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有对应的Session存在
func GetSession(account string) (*skynet.ServiceKey, uint64, error) {
	// 构造参数
	in := &RPCGIMSessionGetSessionIn{
		Account: account,
	}
	out := &RPCGIMSessionGetSessionOut{}

	// RPC请求
	if err := client.GetClient(SessionServiceName(account), "1.0.0").Send("GetSession", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return out.ServiceKey, out.RemoteSessionId, nil
		} else {
			return out.ServiceKey, out.RemoteSessionId, fmt.Errorf(out.Code)
		}
	} else {
		return nil, 0, err
	}
}

//////////////////
// RPC SetSession
//////////////////

type RPCGIMSessionSetSessionIn struct {
	Account         string
	ServiceKey      *skynet.ServiceKey
	RemoteSessionId uint64
}

type RPCGIMSessionSetSessionOut struct {
	Code               string
	Account            string
	OldServiceKey      *skynet.ServiceKey
	OldRemoteSessionId uint64
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有OldServiceKey
func SetSession(account string, serviceKey *skynet.ServiceKey, remoteSessionId uint64) (*skynet.ServiceKey, uint64, error) {
	// 构造参数
	in := &RPCGIMSessionSetSessionIn{
		Account:         account,
		ServiceKey:      serviceKey,
		RemoteSessionId: remoteSessionId,
	}
	out := &RPCGIMSessionSetSessionOut{}

	// RPC请求
	if err := client.GetClient(SessionServiceName(account), "1.0.0").Send("SetSession", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return out.OldServiceKey, out.OldRemoteSessionId, nil
		} else {
			return out.OldServiceKey, out.OldRemoteSessionId, fmt.Errorf(out.Code)
		}
	} else {
		return nil, 0, err
	}
}

//////////////////
// RPC DelSession
//////////////////

type RPCGIMSessionDelSessionIn struct {
	Account string
}

type RPCGIMSessionDelSessionOut struct {
	Code               string
	Account            string
	OldServiceKey      *skynet.ServiceKey
	OldRemoteSessionId uint64
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有OldServiceKey
func DelSession(account string) (*skynet.ServiceKey, uint64, error) {
	// 构造参数
	in := &RPCGIMSessionDelSessionIn{
		Account: account,
	}
	out := &RPCGIMSessionDelSessionOut{}

	// RPC请求
	if err := client.GetClient(SessionServiceName(account), "1.0.0").Send("DelSession", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return out.OldServiceKey, out.OldRemoteSessionId, nil
		} else {
			return out.OldServiceKey, out.OldRemoteSessionId, fmt.Errorf(out.Code)
		}
	} else {
		return nil, 0, err
	}
}

///////////////////////////////
// RPC DelSessionWithSessionID
///////////////////////////////

type RPCGIMSessionDelSessionWithSessionIdIn struct {
	Account   string
	SessionId uint64
}

type RPCGIMSessionDelSessionWithSessionIdOut struct {
	Code               string
	Account            string
	OldServiceKey      *skynet.ServiceKey
	OldRemoteSessionId uint64
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否删除了Session
func DelSessionWithSessionId(account string, sessionId uint64) (*skynet.ServiceKey, uint64, error) {
	// 构造参数
	in := &RPCGIMSessionDelSessionWithSessionIdIn{
		Account:   account,
		SessionId: sessionId,
	}
	out := &RPCGIMSessionDelSessionOut{}

	// RPC请求
	if err := client.GetClient(SessionServiceName(account), "1.0.0").Send("DelSessionWithSessionId", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return out.OldServiceKey, out.OldRemoteSessionId, nil
		} else {
			return out.OldServiceKey, out.OldRemoteSessionId, fmt.Errorf(out.Code)
		}
	} else {
		return nil, 0, err
	}
}
