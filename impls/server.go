package impls

import (
	"fmt"
	"github.com/xiangrui2019/tcper/utils"
	"log"
	"net"
)

var connectionId uint32

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) listenTCP() (*net.TCPListener, error) {
	addr, err := net.ResolveTCPAddr(utils.GetTCPVersion(),
		fmt.Sprintf("%s:%s", utils.GetIpAddr(), utils.GetPort()))

	if err != nil {
		return nil, err
	}

	listenner, err := net.ListenTCP(utils.GetTCPVersion(), addr)

	if err != nil {
		return nil, err
	}

	return listenner, nil
}

func (server *Server) Start() error {
	log.Printf("[Tcper] Server Listenner at IP Addr: %s, Port: %s, TCP Version: %s, is starting\n",
		utils.GetIpAddr(), utils.GetPort(), utils.GetTCPVersion())

	listenner, err := server.listenTCP()

	if err != nil {
		log.Println(err)

		return err
	}

	log.Println("[Tcper] Server started success.")

	for {
		conn, err := listenner.AcceptTCP()

		if err != nil {
			log.Println(err)
			continue
		}

		dealConn := NewConnection(
			conn, connectionId, nil)

		connectionId++

		dealConn.Start()
	}
}

func (server *Server) Stop() {

}

func (server *Server) Serve() {
	server.Start()

	select {}
}
