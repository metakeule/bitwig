package bitwig

import (
	"fmt"
	"net"
	"sync"

	"github.com/scgolang/osc"
)

// Handler handles a bitwig connection
type Handler interface {

	// Handle handles a message from bitwig
	Handle(m osc.Message)
}

// HandlerFunc is a function that serves as Handler
type HandlerFunc func(m osc.Message)

func (h HandlerFunc) Handle(m osc.Message) {
	h(m)
}

// Connection is a connection to a bitwig instance
type Connection interface {
	// Send sends a message to bitwig
	Send(msg osc.Message)

	// Close closes the connection
	Close()

	// IsClosed returns wether the connection is closed
	IsClosed() bool
}

type connection struct {
	handler    map[string]osc.MessageHandler
	sendChan   chan osc.Message
	client     *osc.UDPConn
	server     *osc.UDPConn
	errChan    chan error
	closeChan  chan bool
	sendStr    string
	receiveStr string
	closed     bool
	sync.RWMutex
}

type Option func(c *connection)

func SendAddress(host string, port int) Option {
	return func(c *connection) {
		c.sendStr = fmt.Sprintf("%s:%d", host, port)
	}
}

func ListenAddress(host string, port int) Option {
	return func(c *connection) {
		c.receiveStr = fmt.Sprintf("%s:%d", host, port)
	}
}

func Connect(h map[string]osc.MessageHandler, opts ...Option) (Connection, error) {
	var c = &connection{}
	c.handler = h
	c.errChan = make(chan error)
	c.closeChan = make(chan bool)
	c.sendChan = make(chan osc.Message, 1)

	// defaults as of OSC4Bitwig
	c.sendStr = "127.0.0.1:8000"
	c.receiveStr = "127.0.0.1:9000"
	//	c.receiveStr = "127.0.0.1:10000"

	for _, opt := range opts {
		opt(c)
	}

	if h != nil {

		// Setup the server.
		laddr2, err := net.ResolveUDPAddr("udp", c.receiveStr)
		if err != nil {
			return nil, fmt.Errorf("can't listen on udp %s: %s", c.receiveStr, err.Error())
		}
		c.server, err = osc.ListenUDP("udp", laddr2)

		if err != nil {
			return nil, fmt.Errorf("can't listen on udp %s: %s", c.receiveStr, err.Error())
		}

		go c.listen()
	}

	raddr2, err := net.ResolveUDPAddr("udp", c.sendStr)
	if err != nil {
		return nil, fmt.Errorf("can't send on udp %s: %s", c.sendStr, err.Error())
	}
	c.client, err = osc.DialUDP("udp", nil, raddr2)

	if err != nil {
		return nil, fmt.Errorf("can't send on udp %s: %s", c.sendStr, err.Error())
	}

	go func() {

		for {

			select {
			case <-c.closeChan:
				c.server.Close()
				c.client.Close()
				close(c.sendChan)
				close(c.errChan)
				return
			case err := <-c.errChan:
				c.Lock()
				c.closed = true
				c.Unlock()
				c.server.Close()
				c.client.Close()
				close(c.sendChan)
				close(c.errChan)
				fmt.Printf("error: %s\n", err.Error())
				_ = err
				return
			case msg := <-c.sendChan:
				//fmt.Printf("trying to send to bitwig: %#v\n", msg)
				if err := c.client.Send(msg); err != nil {
					fmt.Printf("error while sending %v: %s\n", msg, err.Error())
					c.errChan <- fmt.Errorf("error while sending %v: %s\n", msg, err.Error())
				}
			}
		}
	}()

	return c, nil
}

func (c *connection) listen() {
	err := c.server.Serve(1, osc.Dispatcher(c.handler))
	if err != nil {
		c.errChan <- fmt.Errorf("error while starting to listen (local %v, remote %v): %v", c.server.LocalAddr(), c.server.RemoteAddr(), err)
	}
}

func (c *connection) Send(msg osc.Message) {
	var cl bool
	c.RLock()
	cl = c.closed
	c.RUnlock()
	if cl {
		// fmt.Println("bitwig connection has been closed")
		return
	}
	c.sendChan <- msg
}

func (c *connection) Close() {
	var cl bool
	c.RLock()
	cl = c.closed
	c.RUnlock()
	if cl {
		return
	}
	c.Lock()
	c.closeChan <- true
	c.closed = true
	c.Unlock()
}

func (c *connection) IsClosed() bool {
	var cl bool
	c.RLock()
	cl = c.closed
	c.RUnlock()
	return cl
}
