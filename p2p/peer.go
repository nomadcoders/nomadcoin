package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()
	var keys []string
	for key := range p.v {
		keys = append(keys, key)
	}
	return keys
}

func (p *peer) close() {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	p.conn.Close()
	delete(Peers.v, p.key)
}

func (p *peer) read() {
	defer p.close()
	for {
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			break
		}
		handleMsg(&m, p)
	}
}

func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}
	go p.read()
	go p.write()
	Peers.v[key] = p
	return p
}
