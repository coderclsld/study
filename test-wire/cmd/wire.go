package main

import (
	. "clsld.com/test-wire/data"
	"github.com/google/wire"
)

func InitializeBroadCast() BroadCast {
	wire.Build(NewBroadCast, NewChannel, NewMessage)
	return BroadCast{}
}
