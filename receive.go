package bitwig

import (
	"github.com/scgolang/osc"
)

func ReceiveSwitch(msg osc.Message) bool {
	i, _ := msg.Arguments[0].ReadInt32()

	return i == 1
}

func ReceiveValue(msg osc.Message) int8 {
	i, _ := msg.Arguments[0].ReadInt32()

	return int8(i)
}
