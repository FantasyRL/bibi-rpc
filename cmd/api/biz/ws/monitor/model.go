package monitor

import (
	"context"
	"github.com/hertz-contrib/websocket"
)

type Client struct {
	ID       int64
	TargetId int64
	Socket   *websocket.Conn
	Send     chan []byte
	Ctx      context.Context
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int64
}

// ClientManager Manager client user
type ClientManager struct {
	Clients    map[int64]*Client //manager
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client //login
	Unregister chan *Client //exit
}

var Manager = ClientManager{
	Clients:    make(map[int64]*Client),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}
