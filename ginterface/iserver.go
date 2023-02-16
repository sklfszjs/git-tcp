package ginterface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(IRouter)
}
