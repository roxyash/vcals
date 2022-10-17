package chat

type Hub struct {
	clients     map[*Client]bool
	broadcast   chan []byte
	register    chan *Client
	ungregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		ungregister: make(chan *Client),
		clients:     make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.ungergister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <- h.broadcast:
			for client := range h.clients{
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
