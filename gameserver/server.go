package gameserver

import (
	"errors"
	"log"
	"net/http"
)

const (
	maxClients = 256
)

var (
	ErrNoMoreClientSlots = errors.New("No more free client slots.")
)

type Server struct {
	addr string

	//
	clientSlots []bool

	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// Create new chat server.
func NewServer(pattern string) *Server {
	return &Server{
		addr:        ":8080",
		clientSlots: make([]bool, maxClients),
		broadcast:   make(chan Message),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.serveWs(w, r)
	})
	println("Listening server...")
	err := http.ListenAndServe(s.addr, nil)
	if err != nil {
		println("Failed to listen:", err.Error())
	}
}

func (s *Server) ListenTLS(sslCert string, sslKey string) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.serveWs(w, r)
	})
	println("Listening server...")
	err := http.ListenAndServeTLS(s.addr, sslCert, sslKey, nil)
	if err != nil {
		println("Failed to listen:", err.Error())
	}
}

func (s *Server) ChRegister() chan *Client { return s.register }

func (s *Server) ChUnregister() chan *Client { return s.unregister }

func (s *Server) ChBroadcast() chan Message { return s.broadcast }

func (s *Server) GetMaxClients() int32 { return int32(len(s.clientSlots)) }

func (s *Server) GetClients() map[*Client]bool { return s.clients }

func (s *Server) HasClient(c *Client) bool {
	_, ok := s.clients[c]
	return ok
}

func (s *Server) RegisterClient(c *Client, data interface{}) {
	c.data = data
	s.clients[c] = true
}

func (s *Server) RemoveClient(c *Client) bool {
	if _, ok := s.clients[c]; ok {
		s.clientSlots[c.clientSlot] = false
		close(c.send)
		delete(s.clients, c)
		return true
	}
	return false
}

// serveWs handles websocket requests from the peer.
func (s *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	clientSlot, err := s.getNextFreeClientSlot()
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		server:     s,
		conn:       conn,
		clientSlot: clientSlot,
		send:       make(chan []byte, 256),
	}
	s.clientSlots[clientSlot] = true
	client.server.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	println("Client connected!")
}

func (s *Server) getNextFreeClientSlot() (int32, error) {
	maxClients := s.GetMaxClients()
	for i := int32(0); i < maxClients; i++ {
		if !s.clientSlots[i] {
			return i, nil
		}
	}
	return 0, ErrNoMoreClientSlots
}
