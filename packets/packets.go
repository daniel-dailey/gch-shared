package packets

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommPacketType int

const (
	CommPacketTypeMessage CommPacketType = iota
	CommPacketTypeOnline
	CommPacketTypeUserList
)

func (cpt CommPacketType) String() string {
	switch cpt {
	case CommPacketTypeMessage:
		return "Message"
	case CommPacketTypeOnline:
		return "Online"
	default:
		return ""
	}
}

type CommPacket struct {
	MessageType   CommPacketType
	MessagePacket *MessagePacket
	OnlinePacket  *OnlinePacket
}

type MessagePacket struct {
	UUID      string
	UserUUID  string
	RoomUUID  string
	UserName  string
	Timestamp int64
	Content   *MessageContent
}

type OnlinePacket struct {
	UUIDNamePacket *UUIDNamePacket
	Online         bool
}

type MessageContent struct {
	StringContent string
}

type UUIDNamePacket struct {
	UUID string `bson:"uuid,omitempty"`
	Name string `bson:"name,omitempty"`
}

type AuthPacket struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username,omitempty"`
	Hash     string             `bson:"hash,omitempty"`
}

func BuildCommPacket(typ CommPacketType, payload interface{}) *CommPacket {
	cp := &CommPacket{
		MessageType: typ,
	}
	switch typ {
	case CommPacketTypeMessage:
		cp.MessagePacket = payload.(*MessagePacket)
	case CommPacketTypeOnline:
		cp.OnlinePacket = payload.(*OnlinePacket)
	}
	return cp
}
