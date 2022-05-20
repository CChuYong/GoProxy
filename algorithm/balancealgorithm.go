package algorithm

import (
	"awesomeProject/backend"
	"net"
)

type BalanceAlgorithm interface {
	GetNextServer(list backend.ServerList, con net.Conn)
}
