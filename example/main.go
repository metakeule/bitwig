package main

import (
	"fmt"
	"time"

	"github.com/metakeule/bitwig"
	"github.com/scgolang/osc"
)

func receivedPlay(msg osc.Message) error {
	fmt.Printf("bitwig said: %v\n", msg)
	return nil
}

func main() {
	handler := map[string]osc.MessageHandler{
		"/play": osc.Method(receivedPlay),
	}

	conn, err := bitwig.Connect(handler)
	if err != nil {
		fmt.Printf("can't connect to bitwig: %#v\n", err.Error())
		return
	}

	fmt.Println("bitwig.Play")
	conn.Send(bitwig.Play())
	time.Sleep(time.Second * 3)
	fmt.Println("bitwig.Stop")
	conn.Send(bitwig.Stop())
	time.Sleep(time.Second * 2)
}
