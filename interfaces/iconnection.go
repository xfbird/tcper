package interfaces

import (
	"github.com/xiangrui2019/tcper/impls"
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

type HandlerFunc func(*impls.Connection, []byte, int) error
