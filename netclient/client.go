package netclient

import "time"

const (
	// Time allowed to write a message to the peer.
	writeWait = 5000 * time.Millisecond

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type clientShared struct {
	// Inbound messages from the server.
	recv chan []byte

	// Disconnect
	disconnect chan bool
}

func newClientShared() clientShared {
	return clientShared{
		recv:       make(chan []byte, 256),
		disconnect: make(chan bool),
	}
}

func (c *clientShared) ChRecv() chan []byte { return c.recv }

func (c *clientShared) ChDisconnected() chan bool { return c.disconnect }
