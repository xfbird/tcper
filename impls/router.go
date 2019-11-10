package impls

import "github.com/xiangrui2019/tcper/interfaces"

type BaseRouter struct {
}

func (router *BaseRouter) BeforeMiddleware(request interfaces.IRequest) {

}

func (router *BaseRouter) API(request interfaces.IRequest) {

}

func (router *BaseRouter) AfterMiddleware(request interfaces.IRequest) {

}
