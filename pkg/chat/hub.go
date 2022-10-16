package chat

type Hub struct {
	clients     map[*Client]bool
	broadcast   chan []byte
	register    chan *Client
	ungregister chan *Client
}
