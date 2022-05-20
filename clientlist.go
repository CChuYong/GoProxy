package main

import (
	"net"
	"sync/atomic"
)

type ClientList struct {
	listMap   map[net.Conn]*UpstreamBridge
	connected int32
}

func (c *ClientList) addClient(con *UpstreamBridge) {
	c.listMap[con.sourceConn] = con
	atomic.AddInt32(&c.connected, 1)
}

func (c *ClientList) removeClient() {

}
