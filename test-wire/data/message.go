package data

import "fmt"

type Message string
type Channel struct {
	Message Message
}
type BroadCast struct {
	Channel Channel
}

func NewMessage() Message {
	return Message("Hello Wire!")
}
func NewChannel(m Message) Channel {
	return Channel{Message: m}
}
func NewBroadCast(c Channel) BroadCast {
	return BroadCast{Channel: c}
}
func (c Channel) GetMsg() Message {
	return c.Message
}
func (b BroadCast) Start() {
	msg := b.Channel.GetMsg()
	fmt.Println(msg)
}
