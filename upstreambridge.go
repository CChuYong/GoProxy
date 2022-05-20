package main

import (
	"net"
)

type UpstreamBridge struct {
	Bridge
	sourceConn net.Conn
}

func (usb *UpstreamBridge) init() {

}

func NewUpstreamBridge(con net.Conn) *UpstreamBridge {
	usb := &UpstreamBridge{sourceConn: con}
	go usb.init()
	return usb
}

func (usb *UpstreamBridge) OnTraffic() {

}