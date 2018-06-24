package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/silbinarywolf/networkplatformer-go/gameserver"
	"github.com/silbinarywolf/networkplatformer-go/netmsg"
)

var (
	server *Server
)

type Server struct {
	*gameserver.Server
}

func NewServer() *Server {
	server := &Server{
		Server: gameserver.NewServer("/abgame"),
	}
	return server
}

func (s *Server) Update() {
RecvMsgLoop:
	for {
		select {
		case client := <-s.ChRegister():
			clientSlot := int32(client.ClientSlot())

			// Create player instance
			char := &Char{
				X: float64(rand.Int63n(90) + 130),
				Y: float64(40 - rand.Int63n(30)),
			}

			// Create client
			s.RegisterClient(client, char)

			// Add client to simulation
			chars = append(chars, char)

			// Send connection response
			{
				sendMsg := netmsg.ConnectResponse{
					ClientSlot: clientSlot,
					X:          char.X,
					Y:          char.Y,
				}
				data, err := proto.Marshal(&sendMsg)
				if err != nil {
					log.Fatal("client connect: marshaling error: ", err)
				}

				// Send to connecting player their information
				packetData := make([]byte, 1, len(data)+1)
				packetData[0] = netmsg.MsgConnectResponse
				packetData = append(packetData, data...)
				client.SendMessage(packetData)
			}
		case client := <-s.ChUnregister():
			if s.RemoveClient(client) {
				char := client.Data().(*Char)
				char.RemoveFromSimulation()

				log.Printf("client #%d disconnected", client.ClientSlot())

				// Tell clients player disconnected
				sendMsg := netmsg.DisconnectPlayer{
					ClientSlot: int32(client.ClientSlot()),
				}
				data, err := proto.Marshal(&sendMsg)
				if err != nil {
					log.Fatal("client connect: marshaling error: ", err)
				}

				// Send disconnected player
				packetData := make([]byte, 1, len(data)+1)
				packetData[0] = netmsg.MsgDisconnectPlayer
				packetData = append(packetData, data...)
				for otherClient := range s.GetClients() {
					if otherClient == client {
						continue
					}
					otherClient.SendMessage(packetData)
				}
			}
		case message := <-s.ChBroadcast():
			var (
				client = message.Client()
				buf    = message.Data()
			)
			kind := netmsg.Kind(buf[0])
			buf = buf[1:]
			switch kind {
			case netmsg.MsgUpdatePlayer:
				// Receive update
				recvMsg := &netmsg.UpdatePlayer{}
				err := recvMsg.Unmarshal(buf)
				if err != nil {
					log.Fatal("unmarshaling error: ", err)
					break
				}
				char := client.Data().(*Char)
				char.X = recvMsg.X
				char.Y = recvMsg.Y
				char.isKeyLeftPressed = recvMsg.IsKeyLeftPressed
				char.isKeyRightPressed = recvMsg.IsKeyRightPressed
			default:
				log.Printf("Unhandled netmsg kind: %s, with data: %v\n", kind.String(), buf)
			}
		default:
			// no-op
			break RecvMsgLoop
		}
	}

	// Send updates to clients
	for client := range s.GetClients() {
		char := client.Data().(*Char)
		elapsed := time.Since(char.lastUpdatedTimer)
		if elapsed > 15*time.Millisecond {
			// Reset countdown till next update
			char.lastUpdatedTimer = time.Now()

			// Send update to other players
			sendMsg := netmsg.UpdatePlayer{
				ClientSlot:        int32(client.ClientSlot()),
				X:                 char.X,
				Y:                 char.Y,
				IsKeyLeftPressed:  char.isKeyLeftPressed,
				IsKeyRightPressed: char.isKeyRightPressed,
			}
			data, err := proto.Marshal(&sendMsg)
			if err != nil {
				log.Fatal("client update: marshaling error: ", err)
			}
			packetData := make([]byte, 1, len(data)+1)
			packetData[0] = netmsg.MsgUpdatePlayer
			packetData = append(packetData, data...)
			for otherClient := range s.GetClients() {
				if otherClient == client {
					continue
				}
				otherClient.SendMessage(packetData)
			}
		}
	}
}
