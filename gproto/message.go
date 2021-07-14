package gproto

import (
	"fmt"
)

///////////////////
// 消息相关的模块
///////////////////

// 全局错误码
const (
	CODE_SUCCESS                  = 0  // Success
	CODE_IGNORE                   = 1  // Ignore
	CODE_ERROR_NO_USER            = 2  // User Not Found
	CODE_ERROR_PASSWORD_ERROR     = 3  // Password Error
	CODE_ERROR_UNKNOWN            = 4  // Unknown Error
	CODE_ERROR_REPLACED           = 5  // Session Replaced
	CODE_ERROR_CLIENT             = 6  // Client Logci Error
	CODE_ERROR_SERVER             = 7  // Server Internal Error
	CODE_ERROR_DB                 = 8  // DB Ops Error
	CODE_ERROR_NOT_FOUND          = 9  // Not Found Error
	CODE_ERROR_AUTH_ERROR         = 10 // AUTH Error
	CODE_ERROR_UNSUPPORT_API      = 11 // UnSupport API
	CODE_ERROR_INVALID_JSON       = 12 // Invalid JSON Format
	CODE_ERROR_ALREADY_REGISTERED = 13 // Already Registered Error
	CODE_ERROR_NO_APP             = 14 // No App Error
	CODE_ERROR_ALREADY_EXIST      = 15 // Already Exist
)

func CodeToMessage(code int) string {
	switch code {
	case CODE_SUCCESS:
		return "Success"
	case CODE_IGNORE:
		return "Ignore"
	case CODE_ERROR_NO_USER:
		return "User Not Found"
	case CODE_ERROR_PASSWORD_ERROR:
		return "Password Error"
	case CODE_ERROR_UNKNOWN:
		return "Unknown Error"
	case CODE_ERROR_REPLACED:
		return "Session Replaced"
	case CODE_ERROR_CLIENT:
		return "Client Logci Error"
	case CODE_ERROR_SERVER:
		return "Server Internal Error"
	case CODE_ERROR_DB:
		return "DB Ops Error"
	case CODE_ERROR_NOT_FOUND:
		return "Not Found Error"
	case CODE_ERROR_AUTH_ERROR:
		return "Auth Error"
	case CODE_ERROR_UNSUPPORT_API:
		return "UnSupport API"
	case CODE_ERROR_INVALID_JSON:
		return "Invalid JSON Format"
	case CODE_ERROR_ALREADY_REGISTERED:
		return "Already Registered Error"
	case CODE_ERROR_NO_APP:
		return "No App Error"
	default:
		return fmt.Sprintf("Unknown Code: %v", code)
	}
}

const (
	MSG_PING_REQUEST                 = 0x1001 // C -> S
	MSG_PING_RESPONSE                = 0x1002 // S -> C
	MSG_ECHO_REQUEST                 = 0x1003 // C -> S
	MSG_ECHO_RESPONSE                = 0x1004 // S -> C
	MSG_KICKOFF_NOTIFY               = 0x1005 // S -> C
	MSG_AUTH_REQUEST                 = 0x1006 // C -> S
	MSG_AUTH_RESPONSE                = 0x1007 // S -> C
	MSG_SET_PROFILE_REQUEST          = 0x1008 // C -> S
	MSG_SET_PROFILE_RESPONSE         = 0x1009 // S -> C
	MSG_SET_FRIEND_MARKNAME_REQUEST  = 0x100a // C -> S
	MSG_SET_FRIEND_MARKNAME_RESPONSE = 0x100b // S -> C
	MSG_GET_PROFILES_REQUEST         = 0x100c // C -> S
	MSG_GET_PROFILES_RESPONSE        = 0x100d // S -> C
	MSG_MESSAGE_REQUEST              = 0x100e // C -> S
	MSG_MESSAGE_RESPONSE             = 0x100f // S -> C
	MSG_MESSAGE_NOTIFY               = 0x1010 // S -> C
	MSG_MESSAGE_ACK                  = 0x1011 // C -> S
	MSG_ADD_FRIEND_REQUEST           = 0x1012 // C -> S
	MSG_ADD_FRIEND_RESPONSE          = 0x1013 // S -> C
	MSG_ADD_FRIEND_NOTIFY            = 0x1014 // S -> C
	MSG_ANSWER_FRIEND_REQUEST        = 0x1015 // C -> S
	MSG_ANSWER_FRIEND_RESPONSE       = 0x1016 // S -> C
	MSG_ANSWER_FRIEND_NOTIFY         = 0x1017 // S -> C
	MSG_DEL_FRIEND_REQUEST           = 0x1018 // C -> S
	MSG_DEL_FRIEND_RESPONSE          = 0x1019 // S -> C
	MSG_DEL_FRIEND_NOTIFY            = 0x101a // S -> C
	MSG_GET_FRIENDS_REQUEST          = 0x101b // C -> S
	MSG_GET_FRIENDS_RESPONSE         = 0x101c // S -> C
	MSG_GET_CHATROOM_INFO_REQUEST    = 0x101d // C -> S
	MSG_GET_CHATROOM_INFO_RESPONSE   = 0x101e // S -> S
	MSG_ENTER_CHATROOM_REQUEST       = 0x101f // C -> S
	MSG_ENTER_CHATROOM_RESPONSE      = 0x1020 // S -> C
	MSG_ENTER_CHATROOM_NOTIFY        = 0x1021 // S -> C
	MSG_LEAVE_CHATROOM_REQUEST       = 0x1022 // C -> S
	MSG_LEAVE_CHATROOM_RESPONSE      = 0x1023 // S -> C
	MSG_LEAVE_CHATROOM_NOTIFY        = 0x1024 // S -> C
	MSG_CHATROOM_KICKOFF_NOTIFY      = 0x1025 // S -> C
	MSG_CREATE_GROUP_REQUEST         = 0x1026 // C -> S
	MSG_CREATE_GROUP_RESPONSE        = 0x1027 // S -> C
	MSG_GET_GROUP_INFO_REQUEST       = 0x1028 // C -> S
	MSG_GET_GROUP_INFO_RESPONSE      = 0x1029 // S -> S
	MSG_SET_GROUP_INFO_REQUEST       = 0x102a // C -> S
	MSG_SET_GROUP_INFO_RESPONSE      = 0x102b // S -> S
	MSG_GET_GROUP_LIST_REQUEST       = 0x102c // C -> S
	MSG_GET_GROUP_LIST_RESPONSE      = 0x102d // S -> S
	MSG_ENTER_GROUP_REQUEST          = 0x102e // C -> S
	MSG_ENTER_GROUP_RESPONSE         = 0x1030 // S -> C
	MSG_ENTER_GROUP_NOTIFY           = 0x1031 // S -> C
	MSG_LEAVE_GROUP_REQUEST          = 0x1032 // C -> S
	MSG_LEAVE_GROUP_RESPONSE         = 0x1033 // S -> C
	MSG_LEAVE_GROUP_NOTIFY           = 0x1034 // S -> C
	MSG_GROUP_KICKOFF_NOTIFY         = 0x1035 // S -> C
	MSG_ADD_BLACKLIST_REQUEST        = 0x1035 // C -> S
	MSG_ADD_BLACKLIST_RESPONSE       = 0x1036 // S -> C
	MSG_ADD_BLACKLIST_NOTIFY         = 0x1037 // S -> C
	MSG_DEL_BLACKLIST_REQUEST        = 0x1038 // C -> S
	MSG_DEL_BLACKLIST_RESPONSE       = 0x1039 // S -> C
	MSG_DEL_BLACKLIST_NOTIFY         = 0x103a // S -> C
	MSG_GET_BLACKLISTS_REQUEST       = 0x103b // C -> S
	MSG_GET_BLACKLISTS_RESPONSE      = 0x103c // S -> C
	MSG_SET_GROUP_MARKNAME_REQUEST   = 0x103d // C -> S
	MSG_SET_GROUP_MARKNAME_RESPONSE  = 0x103e // S -> C
)

// optional字段前缀
const (
	OPTIONAL_TRUE  = 1 // 字段写入
	OPTIONAL_FALSE = 0 // 字段不写(忽略), 默认值
)

////////////////////////////////////
// message PingRequest {
// }

type GProtoPingRequest struct {
}

////////////////////////////////////

func (pingRequest *GProtoPingRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		return nil, nil
	}
	return nil, InvalidVersionError
}

func (pingRequest *GProtoPingRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message PingResponse {
//     required uint8 code = 1;
// }

type GProtoPingResponse struct {
	Code uint8
}

////////////////////////////////////

func (pingResponse *GProtoPingResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(pingResponse.Code); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (pingResponse *GProtoPingResponse) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&pingResponse.Code); err != nil {
			return err
		}
		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EchoRequest {
//     required string data = 1;
// }

type GProtoEchoRequest struct {
	Data string
}

////////////////////////////////////

func (echoRequest *GProtoEchoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(echoRequest.Data); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (echoRequest *GProtoEchoRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 2 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if echoRequest.Data, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message EchoResponse {
//     required uint8 code = 1;
//     required string data = 2;
// }

type GProtoEchoResponse struct {
	Code uint8
	Data string
}

////////////////////////////////////

func (echoResponse *GProtoEchoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(echoResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(echoResponse.Data); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (echoResponse *GProtoEchoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&echoResponse.Code); err != nil {
			return err
		}
		if echoResponse.Data, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message KickoffNotify {
//     required uint8 code = 1;
// }

type GProtoKickOffNotify struct {
	Code uint8
}

////////////////////////////////////

func (kickOffNotify *GProtoKickOffNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(kickOffNotify.Code); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (kickOffNotify *GProtoKickOffNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&kickOffNotify.Code); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message AuthRequest {
//     required string appId = 1;
//     required string appSecret = 2;
//     required string account = 3;
//     required string token = 4;
//     required string customVersion = 5;
//     optional uint8 groupFlag = 6;
// }

type GProtoAuthRequest struct {
	Account       string
	Token         string
	CustomVersion string
}

////////////////////////////////////

func (authRequest *GProtoAuthRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(authRequest.Account); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(authRequest.Token); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(authRequest.CustomVersion); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (authRequest *GProtoAuthRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if authRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		if authRequest.Token, err = buffer.ReadString(); err != nil {
			return err
		}
		if authRequest.CustomVersion, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message AuthResponse {
//     required uint8 code = 1;
//     required string account = 2;
// }

type GProtoAuthResponse struct {
	Code    uint8
	Account string
}

////////////////////////////////////

func (authResponse *GProtoAuthResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(authResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(authResponse.Account); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (authResponse *GProtoAuthResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 2 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&authResponse.Code); err != nil {
			return err
		}
		if authResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message MessageRequest {
//     required string msgId = 1;
//     required uint8 msgType = 2;
//     required string content = 3;
//     required string from = 4;
//     required string to = 5;
//     optional uint32 record_time = 6;
//     optional string extend = 7;
// }

const (
	// 0 - 9
	MESSAGE_TYPE_PLAYER_TEXT              = 0
	MESSAGE_TYPE_PLAYER_IMAGE             = 1
	MESSAGE_TYPE_PLAYER_SHORT_VOICE       = 2
	MESSAGE_TYPE_PLAYER_SHORT_VEDEO       = 3
	MESSAGE_TYPE_PLAYER_CUSTOMIZE_ONLINE  = 4 // 自定义格式消息(不会离线存储)
	MESSAGE_TYPE_PLAYER_CUSTOMIZE_OFFLINE = 5 // 自定义格式消息(会离线存储)

	// 10 - 19
	MESSAGE_TYPE_GROUP_TEXT              = 10
	MESSAGE_TYPE_GROUP_IMAGE             = 11
	MESSAGE_TYPE_GROUP_SHORT_VOICE       = 12
	MESSAGE_TYPE_GROUP_SHORT_VEDEO       = 13
	MESSAGE_TYPE_GROUP_CUSTOMIZE_ONLINE  = 14 // 自定义格式消息(不会离线存储)
	MESSAGE_TYPE_GROUP_CUSTOMIZE_OFFLINE = 15 // 自定义格式消息(会离线存储)

	// 20 - 29
	MESSAGE_TYPE_CHATROOM_TEXT             = 20
	MESSAGE_TYPE_CHATROOM_IMAGE            = 21
	MESSAGE_TYPE_CHATROOM_SHORT_VOICE      = 22
	MESSAGE_TYPE_CHATROOM_SHORT_VEDEO      = 23
	MESSAGE_TYPE_CHATROOM_CUSTOMIZE_ONLINE = 24 // 自定义格式消息(不会离线存储)

	// 30 - 39
	MESSAGE_TYPE_BROADCAST_TEXT             = 30 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_IMAGE            = 31 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_SHORT_VOICE      = 32 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_SHORT_VIDEO      = 33 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_CUSTOMIZE_ONLINE = 34 // 自定义格式消息(不会离线存储)

)

type GProtoMessageRequest struct {
	MsgId   string
	MsgType uint8
	Content string
	From    string
	To      string
	Extend  string
}

func (messageRequest *GProtoMessageRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(messageRequest.MsgId); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(messageRequest.MsgType); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageRequest.Content); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageRequest.From); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageRequest.To); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageRequest.Extend); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (messageRequest *GProtoMessageRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if messageRequest.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&messageRequest.MsgType); err != nil {
			return err
		}
		if messageRequest.Content, err = buffer.ReadString(); err != nil {
			return err
		}
		if messageRequest.From, err = buffer.ReadString(); err != nil {
			return err
		}
		if messageRequest.To, err = buffer.ReadString(); err != nil {
			return err
		}
		if messageRequest.Extend, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message MessageResponse {
//     required uint8 code = 1;
//     required string msgId = 2;
//     optional uint8 msgType = 3;
//     optional string content = 4;
// }

type GProtoMessageResponse struct {
	Code    uint8
	MsgId   string
	MsgType uint8
	Content string
}

func (messageResponse *GProtoMessageResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(messageResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageResponse.MsgId); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(messageResponse.MsgType); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageResponse.Content); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (messageResponse *GProtoMessageResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&messageResponse.Code); err != nil {
			return err
		}
		if messageResponse.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&messageResponse.MsgType); err != nil {
			return err
		}
		if messageResponse.Content, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message MessageNotify {
//     required string msgId = 1;
//     required uint8 msgType = 2;
//     required string content = 3;
//     required string from = 4;
//     required string to = 5;
//     required uint64 send_time = 6;
//	   required uint64 send_time_ack = 7;
//     optional string extend = 8;
// }

type GProtoMessageNotify struct {
	MsgId       string
	MsgType     uint8
	Content     string
	From        string
	To          string
	SendTime    uint64
	SendTimeAck uint64
	Extend      string
}

func (messageNotify *GProtoMessageNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(messageNotify.MsgId); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(messageNotify.MsgType); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageNotify.Content); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageNotify.From); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageNotify.To); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt64(messageNotify.SendTime); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt64(messageNotify.SendTimeAck); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(messageNotify.Extend); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (messageNotify *GProtoMessageNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if messageNotify.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&messageNotify.MsgType); err != nil {
			return err
		}
		if messageNotify.Content, err = buffer.ReadString(); err != nil {
			return err
		}
		if messageNotify.From, err = buffer.ReadString(); err != nil {
			return err
		}
		if messageNotify.To, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt64(&messageNotify.SendTime); err != nil {
			return err
		}
		if err = buffer.ReadUInt64(&messageNotify.SendTimeAck); err != nil {
			return err
		}
		if messageNotify.Extend, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message MessageAck {
//     required string msgId = 1;
// }

type GProtoMessageAck struct {
	MsgId string
}

func (messageAck *GProtoMessageAck) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(messageAck.MsgId); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (messageAck *GProtoMessageAck) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if messageAck.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message GetChatRoomInfoRequest {
//     required string chatroom_id = 1;
// }

type GProtoGetChatRoomInfoRequest struct {
	ChatRoomId string
}

func (getChatRoomInfoRequest *GProtoGetChatRoomInfoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(getChatRoomInfoRequest.ChatRoomId); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (getChatRoomInfoRequest *GProtoGetChatRoomInfoRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if getChatRoomInfoRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message GetChatRoomInfoResponse {
//     required uint8 code = 1;
//     optional ChatRoomInfo info = 2;
// }

type GProtoGetChatRoomInfoResponse struct {
	Code uint8
	Info *GProtoChatRoomInfo
}

func (getChatRoomInfoResponse *GProtoGetChatRoomInfoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(getChatRoomInfoResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteStruct(1, getChatRoomInfoResponse.Info); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (getChatRoomInfoResponse *GProtoGetChatRoomInfoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&getChatRoomInfoResponse.Code); err != nil {
			return err
		}
		if err = buffer.ReadStruct(1, getChatRoomInfoResponse.Info); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomRequest {
//     required string chatroom_id = 1;
// }

type GProtoEnterChatRoomRequest struct {
	ChatRoomId string
}

func (enterChatRoomRequest *GProtoEnterChatRoomRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(enterChatRoomRequest.ChatRoomId); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (enterChatRoomRequest *GProtoEnterChatRoomRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if enterChatRoomRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomResponse {
//     required uint8 code = 1;
//     optional ChatRoomInfo info = 2;
//     repeated OtherUser users = 3;
// }

type GProtoEnterChatRoomResponse struct {
	Code  uint8
	Users []*GProtoOtherUser
}

func (enterChatRoomResponse *GProtoEnterChatRoomResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(enterChatRoomResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteArray(1, enterChatRoomResponse.Users); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (enterChatRoomResponse *GProtoEnterChatRoomResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&enterChatRoomResponse.Code); err != nil {
			return err
		}
		var usersNum uint16
		if err = buffer.ReadUInt16(&usersNum); err != nil {
			return err
		} else {
			enterChatRoomResponse.Users = make([]*GProtoOtherUser, usersNum, usersNum)
			for i := 0; i < int(usersNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						enterChatRoomResponse.Users[i] = &GProtoOtherUser{}
						if err = enterChatRoomResponse.Users[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomNotify {
//     required string chatroom_id = 1;
//     required OtherUser user = 2;
// }

type GProtoEnterChatRoomNotify struct {
	ChatRoomId string
	User       *GProtoOtherUser
}

func (enterChatRoomNotify *GProtoEnterChatRoomNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(enterChatRoomNotify.ChatRoomId); err != nil {
			return nil, err
		}
		if err = buffer.WriteStruct(1, enterChatRoomNotify.User); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (enterChatRoomNotify *GProtoEnterChatRoomNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if enterChatRoomNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		enterChatRoomNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, enterChatRoomNotify.User); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomRequest {
//     required string chatroom_id = 1;
// }

type GProtoLeaveChatRoomRequest struct {
	ChatRoomId string
}

func (leaveChatRoomRequest *GProtoLeaveChatRoomRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(leaveChatRoomRequest.ChatRoomId); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (leaveChatRoomRequest *GProtoLeaveChatRoomRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if leaveChatRoomRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomResponse {
//     required uint8 code = 1;
//     optional string chatroom_id = 2;
// }

type GProtoLeaveChatRoomResponse struct {
	Code       uint8
	ChatRoomId string
}

func (leaveChatRoomResponse *GProtoLeaveChatRoomResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteUInt8(leaveChatRoomResponse.Code); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(leaveChatRoomResponse.ChatRoomId); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (leaveChatRoomResponse *GProtoLeaveChatRoomResponse) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if err = buffer.ReadUInt8(&leaveChatRoomResponse.Code); err != nil {
			return err
		}
		if leaveChatRoomResponse.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomNotify {
//     required string chatroom_id = 1;
//     required string account = 2;
// }

type GProtoLeaveChatRoomNotify struct {
	ChatRoomId string
	Account    string
}

func (leaveChatRoomNotify *GProtoLeaveChatRoomNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(leaveChatRoomNotify.ChatRoomId); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(leaveChatRoomNotify.Account); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (leaveChatRoomNotify *GProtoLeaveChatRoomNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if leaveChatRoomNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		if leaveChatRoomNotify.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

////////////////////////////////////
// message ChatRoomKickoffNotify {
//     required string chatroom_id = 1;
//     required string account = 2;
// }
type GProtoChatRoomKickoffNotify struct {
	ChatRoomId string
	Account    string
}

func (chatRoomKickoffNotify *GProtoChatRoomKickoffNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(chatRoomKickoffNotify.ChatRoomId); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(chatRoomKickoffNotify.Account); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (chatRoomKickoffNotify *GProtoChatRoomKickoffNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if chatRoomKickoffNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		if chatRoomKickoffNotify.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}
