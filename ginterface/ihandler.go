package ginterface

type IHandler interface {
	//调度某一方法
	DoMsgHandler(IRequest)
	//添加处理逻辑
	AddRouter(msgID uint32, router IRouter)
}
