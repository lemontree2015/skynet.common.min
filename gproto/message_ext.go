package gproto

import (
	"fmt"
)

// message SelfUser {
//     required string account = 1;
// }

type GProtoSelfUser struct {
	Account string
}

func EmptyGProtoSelfUser(account string) *GProtoSelfUser {
	return &GProtoSelfUser{
		Account: account,
	}
}

func (selfUser *GProtoSelfUser) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(selfUser.Account); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (selfUser *GProtoSelfUser) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if selfUser.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

func (selfUser *GProtoSelfUser) String() string {
	//return fmt.Sprintf("%v_%v_%v_%v", selfUser.Account, selfUser.Nickname, selfUser.Gender, selfUser.CustomVersion)
	return fmt.Sprintf("%v", selfUser.Account)
}

// message OtherUser {
//     required string account = 1;
// }

type GProtoOtherUser struct {
	Account string
}

func EmptyGProtoOtherUser(account string) *GProtoOtherUser {
	return &GProtoOtherUser{
		Account: account,
	}
}

func (otherUser *GProtoOtherUser) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(otherUser.Account); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (otherUser *GProtoOtherUser) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if otherUser.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

func (otherUser *GProtoOtherUser) String() string {
	return fmt.Sprintf("%v", otherUser.Account)
	//return fmt.Sprintf("%v_%v_%v", otherUser.Account, otherUser.Nickname, otherUser.Gender)
}

///////////////////////////////////////
// message Friend {
//     required string account = 1;
//     required flag flag = 2;
//     optional string markname = 3;
// }

type GProtoFriend struct {
	Account      string
	Flag         uint8
	Markname     string
	MakrnameFlag uint8
}

func (friend *GProtoFriend) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(friend.Account); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(friend.Flag); err != nil {
			return nil, err
		}
		if friend.MakrnameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(friend.Markname); err != nil {
				return nil, err
			}
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (friend *GProtoFriend) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if friend.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&friend.Flag); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&friend.MakrnameFlag); err != nil {
			return err
		}
		if friend.MakrnameFlag == OPTIONAL_TRUE {
			if friend.Markname, err = buffer.ReadString(); err != nil {
				return err
			}
		}
		return nil
	}
	return InvalidVersionError
}

///////////////////////////////////////
// message Blacklist {
//     required string account = 1;
//     required flag flag = 2;
// }

type GProtoBlacklist struct {
	Account string
	Flag    uint8
}

func (blacklist *GProtoBlacklist) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(blacklist.Account); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(blacklist.Flag); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (blacklist *GProtoBlacklist) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if blacklist.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&blacklist.Flag); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

///////////////////////////////////////
// message ChatRoomInfo {
//     required string chatroom_id = 1;
//     requierd string name = 2;
//     requierd string description = 3;
//     required string owner_account = 4;
//     required uint16 max_user = 5;
//     required string data = 6;
//     required int64 created_timestamp = 7;
// }

type GProtoChatRoomInfo struct {
	ChatRoomId       string
	Name             string
	Description      string
	OwnerAccount     string
	MaxUser          uint16
	CreatedTimestamp int64
}

func (chatRoomInfo *GProtoChatRoomInfo) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(chatRoomInfo.ChatRoomId); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(chatRoomInfo.Name); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(chatRoomInfo.Description); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(chatRoomInfo.OwnerAccount); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt16(chatRoomInfo.MaxUser); err != nil {
			return nil, err
		}
		if err = buffer.WriteInt64(chatRoomInfo.CreatedTimestamp); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (chatRoomInfo *GProtoChatRoomInfo) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if chatRoomInfo.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		if chatRoomInfo.Name, err = buffer.ReadString(); err != nil {
			return err
		}
		if chatRoomInfo.Description, err = buffer.ReadString(); err != nil {
			return err
		}
		if chatRoomInfo.OwnerAccount, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt16(&chatRoomInfo.MaxUser); err != nil {
			return err
		}
		if err = buffer.ReadInt64(&chatRoomInfo.CreatedTimestamp); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}

///////////////////////////////////////
// message GroupInfo {
//     required string group_id = 1;
//     requierd string name = 2;
//     requierd string avatar = 3;
//     requierd string description = 4;
//     required string owner_account = 5;
//     required uint16 max_user = 6;
//     required string data = 7;
//     required uint8 flag = 8;
//     required int64 created_timestamp = 9;
//     repeated OtherUser users = 10;
//     repeated GroupMarkname marknames = 11;
// }

type GProtoGroupInfo struct {
	GroupId          string
	Name             string
	Avatar           string
	Description      string
	OwnerAccount     string
	MaxUser          uint16
	Flag             uint8
	CreatedTimestamp int64
	Users            []*GProtoOtherUser
	Marknames        []*GProtoGroupMarkname
}

func (groupInfo *GProtoGroupInfo) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(groupInfo.GroupId); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(groupInfo.Name); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(groupInfo.Avatar); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(groupInfo.Description); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(groupInfo.OwnerAccount); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt16(groupInfo.MaxUser); err != nil {
			return nil, err
		}
		if err = buffer.WriteUInt8(groupInfo.Flag); err != nil {
			return nil, err
		}
		if err = buffer.WriteInt64(groupInfo.CreatedTimestamp); err != nil {
			return nil, err
		}
		if err = buffer.WriteArray(1, groupInfo.Users); err != nil {
			return nil, err
		}
		if err = buffer.WriteArray(1, groupInfo.Marknames); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (groupInfo *GProtoGroupInfo) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if groupInfo.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}
		if groupInfo.Name, err = buffer.ReadString(); err != nil {
			return err
		}
		if groupInfo.Avatar, err = buffer.ReadString(); err != nil {
			return err
		}
		if groupInfo.Description, err = buffer.ReadString(); err != nil {
			return err
		}
		if groupInfo.OwnerAccount, err = buffer.ReadString(); err != nil {
			return err
		}
		if err = buffer.ReadUInt16(&groupInfo.MaxUser); err != nil {
			return err
		}
		if err = buffer.ReadUInt8(&groupInfo.Flag); err != nil {
			return err
		}
		if err = buffer.ReadInt64(&groupInfo.CreatedTimestamp); err != nil {
			return err
		}
		var usersNum uint16
		if err = buffer.ReadUInt16(&usersNum); err != nil {
			return err
		} else {
			groupInfo.Users = make([]*GProtoOtherUser, usersNum, usersNum)
			for i := 0; i < int(usersNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						groupInfo.Users[i] = &GProtoOtherUser{}
						if err = groupInfo.Users[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}
		var marknamesNum uint16
		if err = buffer.ReadUInt16(&marknamesNum); err != nil {
			return err
		} else {
			groupInfo.Marknames = make([]*GProtoGroupMarkname, marknamesNum, marknamesNum)
			for i := 0; i < int(marknamesNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						groupInfo.Marknames[i] = &GProtoGroupMarkname{}
						if err = groupInfo.Marknames[i].Decode(1, bufObj); err != nil {
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

//////////////////////////
// message GroupMarkname {
//     required string account = 1;
//     required string markname = 2;
// }

type GProtoGroupMarkname struct {
	Account  string
	Markname string
}

func (groupMarkname *GProtoGroupMarkname) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		buffer := NewEmptyGBuffer()
		var err error
		if err = buffer.WriteString(groupMarkname.Account); err != nil {
			return nil, err
		}
		if err = buffer.WriteString(groupMarkname.Markname); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return nil, InvalidVersionError
}

func (groupMarkname *GProtoGroupMarkname) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}
	if version == 1 {
		buffer := NewGBuffer(buf)
		var err error
		if groupMarkname.Account, err = buffer.ReadString(); err != nil {
			return err
		}
		if groupMarkname.Markname, err = buffer.ReadString(); err != nil {
			return err
		}
		return nil
	}
	return InvalidVersionError
}
