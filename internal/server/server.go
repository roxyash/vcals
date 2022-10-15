package server

import (
	"flag"
	"os"
	"time"
)

var (
	addr = flag.String("addr",":",os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")

)


func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8000"
	}

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/root/:uuid/websocket", )
	app.Get("/root/:uuid/chat", handlers.RoomChat)
	app.Get("/root/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stream/:ssuid/chat/websocket",)
	app.Get("/stream/:ssuid/viewer/websocket")
}