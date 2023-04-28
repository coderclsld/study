package main

import (
	. "clsld.com/test-wire/data"
)

func main() {
	message := NewMessage()
	channel := NewChannel(message)
	broadCast := NewBroadCast(channel)
	broadCast.Start()
}
