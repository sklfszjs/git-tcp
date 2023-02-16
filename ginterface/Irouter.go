package ginterface

type IRouter interface {
	PreHandler(IRequest)
	Handler(IRequest)
	PostHandler(IRequest)
}
