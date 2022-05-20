package main

import (
	"net"
)

type UpstreamBridge struct {
	Bridge
	sourceConn net.Conn
	backend    *DownstreamBridge
}

func (usb *UpstreamBridge) init() {
	dsb := NewDownstreamBridge(usb)
	usb.backend = dsb

}

func NewUpstreamBridge(con net.Conn) *UpstreamBridge {
	usb := &UpstreamBridge{sourceConn: con}
	go usb.init()
	return usb
}

func (usb *UpstreamBridge) OnTraffic() {

}
