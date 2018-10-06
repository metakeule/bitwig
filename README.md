# monome
Go library to remote control bitwig studio

[![Documentation](http://godoc.org/github.com/metakeule/bitwig?status.png)](http://godoc.org/github.com/metakeule/bitwig)

## Installation

It is recommended to use Go 1.11 with module support (`$GO111MODULE=on`).

```
go get -d github.com/metakeule/bitwig/...
```

For bitwig studio to be controllable via OSC, you need to install ad enable the OSC control by Jürgen Moßgraber: 

https://github.com/git-moss/DrivenByMoss

Leave the default ports and set the MIDI port to some virtual port and you are ready to go.

## Example

```go
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
```
