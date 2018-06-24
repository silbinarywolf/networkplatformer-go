package main

import (
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/silbinarywolf/networkplatformer-go/gameclient"
	"github.com/silbinarywolf/networkplatformer-go/netmsg"
)

const (
	maxClients = 256
)

var (
	client               *Client
	isConnected          = false
	lastWorldUpdateTimer time.Time
)

type Client struct {
	*gameclient.Client
	clientSlots []*Char
}

func NewClient() *Client {
	server := &Client{
		Client:      gameclient.NewClient(),
		clientSlots: make([]*Char, maxClients),
	}
	return server
}

func (c *Client) Update() {
RecvMsgLoop:
	for {
		select {
		case buf := <-c.ChRecv():
			kind := netmsg.Kind(buf[0])
			buf = buf[1:]
			switch kind {
			case netmsg.MsgConnectResponse:
				recvMsg := &netmsg.ConnectResponse{}
				err := recvMsg.Unmarshal(buf)
				if err != nil {
					log.Fatal("marshaling error: ", err)
					break
				}

				// Receive starting pos from server and add to chars to simulate
				you.X = recvMsg.X
				you.Y = recvMsg.Y
				chars = append(chars, you)
				isConnected = true

				// Last time we received an update about the world
				lastWorldUpdateTimer = time.Now()

				log.Printf("%s: received login data: %v\n", kind, recvMsg)
			case netmsg.MsgUpdatePlayer:
				recvMsg := &netmsg.UpdatePlayer{}
				err := recvMsg.Unmarshal(buf)
				if err != nil {
					log.Fatal("marshaling error: ", err)
					break
				}
				clientSlot := recvMsg.GetClientSlot()
				char := c.clientSlots[clientSlot]
				if char == nil {
					// Create char if they don't exist
					char = &Char{}
					chars = append(chars, char)
					c.clientSlots[clientSlot] = char
				}
				char.X = recvMsg.X
				char.Y = recvMsg.Y
				char.isKeyLeftPressed = recvMsg.IsKeyLeftPressed
				char.isKeyRightPressed = recvMsg.IsKeyRightPressed
			case netmsg.MsgDisconnectPlayer:
				recvMsg := &netmsg.DisconnectPlayer{}
				err := recvMsg.Unmarshal(buf)
				if err != nil {
					log.Fatal("marshaling error: ", err)
					break
				}
				clientSlot := recvMsg.GetClientSlot()
				char := c.clientSlots[clientSlot]
				if char == nil {
					continue
				}
				char.RemoveFromSimulation()
				c.clientSlots[clientSlot] = nil
			default:
				log.Printf("Unhandled netmsg kind: %s, with data: %v", kind.String(), buf)
			}
		case <-c.ChDisconnected():
			isConnected = false
			you.RemoveFromSimulation()

			log.Println("Lost connection to server")
		default:
			// no more messages
			break RecvMsgLoop
		}
	}

	//
	if you != nil && isConnected {
		elapsed := time.Since(lastWorldUpdateTimer)
		if elapsed > 15*time.Millisecond {
			lastWorldUpdateTimer = time.Now()

			// Send update data to server
			sendMsg := netmsg.UpdatePlayer{
				X:                 you.X,
				Y:                 you.Y,
				IsKeyLeftPressed:  you.isKeyLeftPressed,
				IsKeyRightPressed: you.isKeyRightPressed,
			}
			data, err := proto.Marshal(&sendMsg)
			if err != nil {
				log.Fatal("client update: marshaling error: ", err)
			}

			// send update message to server
			packetData := make([]byte, 1, len(data)+1)
			packetData[0] = netmsg.MsgUpdatePlayer
			packetData = append(packetData, data...)
			client.SendMessage(packetData)
		}
	}
}
