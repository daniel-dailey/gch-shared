package packets

import (
	"github.com/daniel-dailey/gch-shared/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommPacketType int

const (
	CommPacketTypeMessage CommPacketType = iota
	CommPacketTypeUserList
)

type CommPacket struct {
	MessageType   CommPacketType
	MessagePacket *MessagePacket
	UserList      map[string]*entities.User
}

type MessagePacket struct {
	UUID      string
	UserUUID  string
	RoomUUID  string
	UserName  string
	Timestamp int64
	Content   *MessageContent
}

type MessageContent struct {
	StringContent string
}

type UUIDNamePacket struct {
	UUID string
	Name string
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
	case CommPacketTypeUserList:
		cp.UserList = payload.(map[string]*entities.User)
	}
	return cp
}
