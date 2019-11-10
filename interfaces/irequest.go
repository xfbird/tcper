package interfaces

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
