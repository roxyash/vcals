package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gguid "github.com/google/uuid"
	w "github.com/roxyash/pkg/webrtc"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", gguid.New().String()))

}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION"{
		ws := "wss"
	}


	uuid, suuid, _ := createOrGetRoom(uuid)
	return c.Render("peer", fiber.Map{
		"RoomWebSocketAddr": fmt.Sprintf("%s://%s/room%s/websocket/", ws, c.Hostname(), uuid),
		"RoomLink": fmt.Sprintf("%s://%s/room/%s", c.Protocol(), c.Hostname(), uuid),
		"ChatWebSocketAddr":fmt.Sprintf("%s://%s/room/%s/chat/websocket", ws, c.Hostname(), uuid),
		"ViewerWebSocketAddr":fmt.Sprintf("%s://%s/room/%s/viewer/websocket", ws, c.Hostname(), uuid),
		"StreamLink": fmt.Sprintf("%s://%s/stream/%s", c.Protocol(), c.Hostname(), suuid),
		"Type": "room",
	}, "layouts/main")
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)

	w.RoomConn(c, room.Peers)
}

func createOrGetRoom(uuid string) (string, string, Room){

}

func RoomViewerWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid  == "" {
		return
	}

	w.RoomsLock.Lock()
	if peer, ok := w.Rooms[uuid]; ok {
		w.RoomsLock.Unlock()
		roomViewerConn(c, peer.Peers)
		return 
	}

	w.RoomsLock.Unlock()
}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type websocketMessage struct {
	Event string `json:"event`
	Data string `json:"data"`
}