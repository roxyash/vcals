package chat

import "github.com/fasthttp/websocket"

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}
