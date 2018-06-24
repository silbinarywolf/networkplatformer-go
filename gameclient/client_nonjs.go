// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package gameclient

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	clientShared
	conn *websocket.Conn

	// Outbound messages to the server.
	send chan []byte
}

func NewClient() *Client {
	return &Client{
		clientShared: newClientShared(),
		conn:         nil,
		send:         make(chan []byte, 256),
	}
}

func (c *Client) Dial(addr string) error {
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+addr+"/ws", nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) DialTLS(addr string) error {
	conn, _, err := websocket.DefaultDialer.Dial("wss://"+addr+"/ws", nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Listen() {
	go c.writePump()
	go c.readPump()
}

func (c *Client) ChRecv() chan []byte { return c.recv }

func (c *Client) SendMessage(message []byte) {
	c.send <- message
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		c.disconnect <- true
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, buf, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.recv <- buf
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
		c.disconnect <- true
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
