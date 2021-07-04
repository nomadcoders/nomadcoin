package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/utils"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleErr(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}
}
