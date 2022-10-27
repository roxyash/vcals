package webrtc

import (
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/roxyash/vcals/pkg/chat"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConnection
	websocket      *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

func (t *ThreadSafeWriter) WriteJSON(v interface{}) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)

}

func (p *Peers) AddTrack(t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP {

}

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP) {

}

func (p *Peers) SignalPeerConnection() {

}

func (p *Peers) DispatchKeyFrame() {

}

type websocketMessage struct {
	Event string `json: "event"`
	Data  string `json:"data"`
}
