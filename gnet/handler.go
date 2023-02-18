package gnet

import (
	"fmt"
	"go-tcp/ginterface"
	"strconv"
)

type Handler struct {
	Apis map[uint32]ginterface.IRouter
}

func NewHandler() ginterface.IHandler {
	return &Handler{
		Apis: make(map[uint32]ginterface.IRouter),
	}
}

func (h *Handler) AddRouter(msgID uint32, router ginterface.IRouter) {
	if _, ok := h.Apis[msgID]; ok {
		panic("repeat api, msgId=" + strconv.Itoa(int(msgID)))
	}
	h.Apis[msgID] = router
	fmt.Println("router add success!")
}

func (h *Handler) DoMsgHandler(request ginterface.IRequest) {
	//1.找到msgID
	//2.找到handler
	//3.调度
	handler, ok := h.Apis[request.GetId()]
	if !ok {
		fmt.Println("handler of", request.GetId(), "not found!")
		return
	}
	handler.PreHandler(request)
	handler.Handler(request)
	handler.PostHandler(request)
}
