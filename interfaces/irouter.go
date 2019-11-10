package interfaces

type IRouter interface {
	BeforeMiddleware(request IRequest)
	API(request IRequest)
	AfterMiddleware(request IRequest)
}
