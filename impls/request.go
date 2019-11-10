package impls

import "github.com/xiangrui2019/tcper/interfaces"

type Request struct {
	conn interfaces.IConnection
	data []byte
}

func (request *Request) GetConnection() interfaces.IConnection {
	return request.conn
}

func (request *Request) GetData() []byte {
	return request.data
}
