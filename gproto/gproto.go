package gproto

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	MagicWord = 0x7fffffff // 4 bytes
)

var (
	InvalidDecodeBufferError = errors.New("invalid Decode Buffer")
	InvalidVersionError      = errors.New("invalid Version")
)

func init() {
	// 初始化Message Descriptions
	messageDescriptions[MSG_PING_REQUEST] = "MSG_PING_REQUEST"
	messageDescriptions[MSG_PING_RESPONSE] = "MSG_PING_RESPONSE"
	messageDescriptions[MSG_ECHO_REQUEST] = "MSG_ECHO_REQUEST"
	messageDescriptions[MSG_ECHO_RESPONSE] = "MSG_ECHO_RESPONSE"
	messageDescriptions[MSG_KICKOFF_NOTIFY] = "MSG_KICKOFF_NOTIFY"
	messageDescriptions[MSG_AUTH_REQUEST] = "MSG_AUTH_REQUEST"
	messageDescriptions[MSG_AUTH_RESPONSE] = "MSG_AUTH_RESPONSE"
	messageDescriptions[MSG_MESSAGE_REQUEST] = "MSG_MESSAGE_REQUEST"
	messageDescriptions[MSG_MESSAGE_RESPONSE] = "MSG_MESSAGE_RESPONSE"
	messageDescriptions[MSG_MESSAGE_NOTIFY] = "MSG_MESSAGE_NOTIFY"
	messageDescriptions[MSG_MESSAGE_ACK] = "MSG_MESSAGE_ACK"
	messageDescriptions[MSG_GET_CHATROOM_INFO_REQUEST] = "MSG_GET_CHATROOM_INFO_REQUEST"
	messageDescriptions[MSG_GET_CHATROOM_INFO_RESPONSE] = "MSG_GET_CHATROOM_INFO_RESPONSE"
	messageDescriptions[MSG_ENTER_CHATROOM_REQUEST] = "MSG_ENTER_CHATROOM_REQUEST"
	messageDescriptions[MSG_ENTER_CHATROOM_RESPONSE] = "MSG_ENTER_CHATROOM_RESPONSE"
	messageDescriptions[MSG_ENTER_CHATROOM_NOTIFY] = "MSG_ENTER_CHATROOM_NOTIFY"
	messageDescriptions[MSG_LEAVE_CHATROOM_REQUEST] = "MSG_LEAVE_CHATROOM_REQUEST"
	messageDescriptions[MSG_LEAVE_CHATROOM_RESPONSE] = "MSG_LEAVE_CHATROOM_RESPONSE"
	messageDescriptions[MSG_LEAVE_CHATROOM_NOTIFY] = "MSG_LEAVE_CHATROOM_NOTIFY"
	messageDescriptions[MSG_CHATROOM_KICKOFF_NOTIFY] = "MSG_CHATROOM_KICKOFF_NOTIFY"
	messageDescriptions[MSG_CREATE_GROUP_REQUEST] = "MSG_CREATE_GROUP_REQUEST"
	messageDescriptions[MSG_CREATE_GROUP_RESPONSE] = "MSG_CREATE_GROUP_RESPONSE"
	messageDescriptions[MSG_GET_GROUP_INFO_REQUEST] = "MSG_GET_GROUP_INFO_REQUEST"
	messageDescriptions[MSG_GET_GROUP_INFO_RESPONSE] = "MSG_GET_GROUP_INFO_RESPONSE"
	messageDescriptions[MSG_SET_GROUP_INFO_REQUEST] = "MSG_SET_GROUP_INFO_REQUEST"
	messageDescriptions[MSG_SET_GROUP_INFO_RESPONSE] = "MSG_SET_GROUP_INFO_RESPONSE"
	messageDescriptions[MSG_GET_GROUP_LIST_REQUEST] = "MSG_GET_GROUP_LIST_REQUEST"
	messageDescriptions[MSG_GET_GROUP_LIST_RESPONSE] = "MSG_GET_GROUP_LIST_RESPONSE"
	messageDescriptions[MSG_ENTER_GROUP_REQUEST] = "MSG_ENTER_GROUP_REQUEST"
	messageDescriptions[MSG_ENTER_GROUP_RESPONSE] = "MSG_ENTER_GROUP_RESPONSE"
	messageDescriptions[MSG_ENTER_GROUP_NOTIFY] = "MSG_ENTER_GROUP_NOTIFY"
	messageDescriptions[MSG_LEAVE_GROUP_REQUEST] = "MSG_LEAVE_GROUP_REQUEST"
	messageDescriptions[MSG_LEAVE_GROUP_RESPONSE] = "MSG_LEAVE_GROUP_RESPONSE"
	messageDescriptions[MSG_LEAVE_GROUP_NOTIFY] = "MSG_LEAVE_GROUP_NOTIFY"
	messageDescriptions[MSG_GROUP_KICKOFF_NOTIFY] = "MSG_GROUP_KICKOFF_NOTIFY"
	messageDescriptions[MSG_SET_GROUP_MARKNAME_REQUEST] = "MSG_SET_GROUP_MARKNAME_REQUEST"
	messageDescriptions[MSG_SET_GROUP_MARKNAME_RESPONSE] = "MSG_SET_GROUP_MARKNAME_RESPONSE"

	// 初始化Message Creators
	messageCreators[MSG_PING_REQUEST] = func() IMessage { return new(GProtoPingRequest) }
	messageCreators[MSG_PING_RESPONSE] = func() IMessage { return new(GProtoPingResponse) }
	messageCreators[MSG_ECHO_REQUEST] = func() IMessage { return new(GProtoEchoRequest) }
	messageCreators[MSG_ECHO_RESPONSE] = func() IMessage { return new(GProtoEchoResponse) }
	messageCreators[MSG_KICKOFF_NOTIFY] = func() IMessage { return new(GProtoKickOffNotify) }
	messageCreators[MSG_AUTH_REQUEST] = func() IMessage { return new(GProtoAuthRequest) }
	messageCreators[MSG_AUTH_RESPONSE] = func() IMessage { return new(GProtoAuthResponse) }
	messageCreators[MSG_MESSAGE_REQUEST] = func() IMessage { return new(GProtoMessageRequest) }
	messageCreators[MSG_MESSAGE_RESPONSE] = func() IMessage { return new(GProtoMessageResponse) }
	messageCreators[MSG_MESSAGE_NOTIFY] = func() IMessage { return new(GProtoMessageNotify) }
	messageCreators[MSG_MESSAGE_ACK] = func() IMessage { return new(GProtoMessageAck) }
	messageCreators[MSG_GET_CHATROOM_INFO_REQUEST] = func() IMessage { return new(GProtoGetChatRoomInfoRequest) }
	messageCreators[MSG_GET_CHATROOM_INFO_RESPONSE] = func() IMessage { return new(GProtoGetChatRoomInfoResponse) }
	messageCreators[MSG_ENTER_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoGetChatRoomInfoResponse) }
	messageCreators[MSG_ENTER_CHATROOM_REQUEST] = func() IMessage { return new(GProtoEnterChatRoomRequest) }
	messageCreators[MSG_ENTER_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoEnterChatRoomResponse) }
	messageCreators[MSG_ENTER_CHATROOM_NOTIFY] = func() IMessage { return new(GProtoEnterChatRoomNotify) }
	messageCreators[MSG_LEAVE_CHATROOM_REQUEST] = func() IMessage { return new(GProtoLeaveChatRoomRequest) }
	messageCreators[MSG_LEAVE_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoLeaveChatRoomResponse) }
	messageCreators[MSG_LEAVE_CHATROOM_NOTIFY] = func() IMessage { return new(GProtoLeaveChatRoomNotify) }
	messageCreators[MSG_CHATROOM_KICKOFF_NOTIFY] = func() IMessage { return new(GProtoKickOffNotify) }
}

// 全局的Message Descriptions
var messageDescriptions map[uint16]string = make(map[uint16]string)

// 全局的Message Creators
type MessageCreator func() IMessage

var messageCreators map[uint16]MessageCreator = make(map[uint16]MessageCreator)

// Command
type Command uint16

func (cmd Command) String() string {
	c := uint16(cmd)
	if desc, ok := messageDescriptions[c]; ok {
		return desc
	} else {
		return fmt.Sprintf("Unknown Command=%v", c)
	}
}

type IMessage interface {
	// 编码IMessage, 如果error=nil, 则编码成功
	Encode(version uint16) ([]byte, error)

	// 解码IMessage, 如果error=nil, 则解码成功
	Decode(version uint16, buf []byte) error
}

type Message struct {
	MagicWord     uint32
	Version       uint16
	Command       uint16
	PayloadLength uint32
	Body          interface{} // IMessage
}

func NewEmptyMessage() *Message {
	return &Message{}
}

func NewMessage(version, command uint16, body interface{}) *Message {
	return &Message{
		MagicWord:     MagicWord,
		Version:       version,
		Command:       command,
		PayloadLength: 0, // 发送的时候会计算 & 长度
		Body:          body,
	}
}

func (message *Message) EncodeHead() ([]byte, error) {
	buf := make([]byte, 12, 12)
	binary.BigEndian.PutUint32(buf[0:], message.MagicWord)
	binary.BigEndian.PutUint16(buf[4:], message.Version)
	binary.BigEndian.PutUint16(buf[6:], message.Command)
	binary.BigEndian.PutUint32(buf[8:], message.PayloadLength)
	return buf[0:], nil
}

func (message *Message) DecodeHead(buf []byte) error {
	if len(buf) != 12 {
		return InvalidDecodeBufferError
	}
	message.MagicWord = binary.BigEndian.Uint32(buf[0:4])
	message.Version = binary.BigEndian.Uint16(buf[4:6])
	message.Command = binary.BigEndian.Uint16(buf[6:8])
	message.PayloadLength = binary.BigEndian.Uint32(buf[8:])
	return nil
}

// 编码Message.Body, 如果error=nil, 则编码成功
func (message *Message) Encode() ([]byte, error) {
	if message.Body != nil {
		if m, ok := message.Body.(IMessage); ok {
			return m.Encode(message.Version)
		}
		panic(fmt.Errorf("Message.Body is NOT IMessage"))
	} else {
		return nil, nil
	}
}

// 解码Message.Body, 如果error=nil, 则解码成功
func (message *Message) Decode(buf []byte) error {
	if creator, ok := messageCreators[message.Command]; ok {
		iMessage := creator()
		err := iMessage.Decode(message.Version, buf)
		message.Body = iMessage // 解码的结果
		return err
	}

	panic(fmt.Errorf("Message.Command Invalid %v", message.Command))
}
