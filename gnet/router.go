package gnet

import "go-tcp/ginterface"

type BaseRouter struct {
}

func (b *BaseRouter) PreHandler(ginterface.IRequest)  {}
func (b *BaseRouter) Handler(ginterface.IRequest)     {}
func (b *BaseRouter) PostHandler(ginterface.IRequest) {}
