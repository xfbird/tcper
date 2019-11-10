package interfaces

type IServer interface {
	Start()
	Stop()
	Serve()
}
