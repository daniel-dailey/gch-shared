package packets

type Message struct {
	UUID     string
	UserUUID string
	UserName string
	Content  *MessageContent
}

type MessageContent struct {
	StringContent string
}
