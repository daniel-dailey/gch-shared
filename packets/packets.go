package packets

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
