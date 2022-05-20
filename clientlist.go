package main

import "net"

type ClientList struct {
	listMap map[net.Conn]*UpstreamBridge
}

func (c *ClientList) addClient(con *UpstreamBridge) {
	c.listMap[con.sourceConn] = con
}

func (c *ClientList) removeClient() {

}
