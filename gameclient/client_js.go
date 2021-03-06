// +build js

package gameclient

import (
	"net"

	"github.com/gopherjs/websocket"
)

type Client struct {
	clientShared

	conn net.Conn

	// Outbound messages to the server.
	send chan []byte
}

func NewClient() *Client {
	return &Client{
		clientShared: newClientShared(),
		conn:         nil,
	}
}

func (c *Client) Dial(addr string) error {
	conn, err := websocket.Dial("ws://" + addr + "/ws") // Blocks until connection is established.
	if err != nil {
		// handle error
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) DialTLS(addr string) error {
	conn, err := websocket.Dial("wss://" + addr + "/ws") // Blocks until connection is established.
	if err != nil {
		// handle error
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Listen() error {
	err := c.readPump() // this is blocking
	c.disconnect <- true
	c.conn.Close()
	return err
}

func (c *Client) SendMessage(message []byte) {
	// NOTE(Jake): 2018-05-27
	//
	// This is not blocking, at least for JavaScript impl. of Websocket.
	//
	// We also don't need to add a write deadline because its non-blocking.
	// (this function returns nil)
	//
	//c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	c.conn.Write(message)
}

func (c *Client) readPump() error {
	conn := c.conn
	for {
		// NOTE(Jake): 2018-06-20
		//
		// Disabling this in hopes that it fixes timeout issues
		// on the server. JavaScript impl. might not need this
		// and might just close automatically after X amount of time.
		//
		//conn.SetReadDeadline(time.Now().Add(pongWait))

		// todo(Jake): 2018-05-27
		//
		// Perhaps profile / figure out how keep allocations here low?
		// Maybe this isnt even a big deal?
		//
		buf := make([]byte, 1024)
		size, err := conn.Read(buf) // Blocks until a WebSocket frame is received.
		if err != nil {
			// handle error
			return err
		}
		buf = buf[:size]
		c.recv <- buf
	}
}
