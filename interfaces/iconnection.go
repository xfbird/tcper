package interfaces

import (
	"net"
)

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnectionId() uint32
	GetRemoteAddr() net.Addr
	Send(data []byte) error
}
