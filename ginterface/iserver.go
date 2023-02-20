package ginterface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(uint32, IRouter)
	GetConnManager() IConnManager
	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	CallOnConnStart(IConnection)
	CallOnConnStop(IConnection)
}
