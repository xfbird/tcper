package impls

import (
	"github.com/xiangrui2019/tcper/utils"
	"log"
	"net"
)

type HandlerFunc func(*Connection, []byte, int) error

type Connection struct {
	Connection   *net.TCPConn
	ConnectionId uint32
	isClosed     bool
	Handler      HandlerFunc
	ExitChan     chan bool
}

func NewConnection(conn *net.TCPConn,
	connectionId uint32,
	handler_func HandlerFunc) *Connection {
	return &Connection{
		Connection:   conn,
		ConnectionId: connectionId,
		isClosed:     false,
		Handler:      handler_func,
		ExitChan:     make(chan bool, 1),
	}
}

func (connection *Connection) StartReader() {
	log.Println("Reader Goroutine is running....")
	defer log.Printf("ConnectionId = %d RemoteAddr = %s Reader is exit.",
		connection.ConnectionId, connection.GetRemoteAddr())
	defer connection.Stop()

	for {
		buffer := make([]byte, utils.GetMaxBuffer())

		cnt, err := connection.Connection.Read(buffer)

		if err != nil {
			log.Println("recvice error.", err)
			continue
		}

		if err := connection.Handler(connection,
			buffer, cnt); err != nil {
			log.Println("handler execute error.")
			continue
		}
	}
}

func (connection *Connection) Start() {
	log.Println("[Tcper Connection] Connection Start().... ConnectionId = ", connection.ConnectionId)

	go connection.StartReader()
}
func (connection *Connection) Stop() {
	log.Println("[Tcper Connection] Connection Stop().... ConnectionId = ", connection.ConnectionId)

	if connection.isClosed == true {
		return
	}

	connection.isClosed = true

	connection.Connection.Close()

	close(connection.ExitChan)
}
func (connection *Connection) GetTCPConnection() *net.TCPConn {
	return connection.Connection
}

func (connection *Connection) GetConnectionId() uint32 {
	return connection.ConnectionId
}
func (connection *Connection) GetRemoteAddr() net.Addr {
	return connection.Connection.RemoteAddr()
}
