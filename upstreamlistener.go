package main

import (
	"github.com/panjf2000/gnet/v2"
	"log"
)

type UpstreamListener struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool

	clients ClientList
}

func (listener *UpstreamListener) Start() {
	log.Fatal(gnet.Run(listener, listener.addr, gnet.WithMulticore(listener.multicore)))
}

func (listener *UpstreamListener) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	c.Write(buf)
	return gnet.None
}

func (listener *UpstreamListener) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	de := NewUpstreamBridge(c)
	listener.clients.addClient(de)
	return
}

func (listener *UpstreamListener) OnBoot(eng gnet.Engine) gnet.Action {
	listener.eng = eng
	return gnet.None
}
