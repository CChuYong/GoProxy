package main

import "github.com/panjf2000/gnet/v2"

type DownstreamBridge struct {
	Bridge
	targetServer BackendServer
	frontend     *UpstreamBridge
}

func (dsb *DownstreamBridge) init() {
	gnet.NewClient()
}

func NewDownstreamBridge(usb *UpstreamBridge) *DownstreamBridge {
	db := &DownstreamBridge{frontend: usb}
	go db.init()
	return db
}
