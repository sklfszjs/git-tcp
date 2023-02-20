package gnet

import (
	"fmt"
	"go-tcp/ginterface"
	"go-tcp/utils"
	"strconv"
)

type Handler struct {
	Apis map[uint32]ginterface.IRouter
	//消息队列
	TaskQueue []chan ginterface.IRequest
	//工作池的worker数量
	WorkerPoolSize uint32
}

func NewHandler() ginterface.IHandler {
	return &Handler{
		Apis:           make(map[uint32]ginterface.IRouter),
		TaskQueue:      make([]chan ginterface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
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

// 启动一个工作池,单例模式
func (h *Handler) StartWorkerPool() {
	for i := 0; i < int(h.WorkerPoolSize); i++ {
		h.TaskQueue[i] = make(chan ginterface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go h.startOneWorker(uint32(i), h.TaskQueue[i])
	}
}

// 启动一个工作流程,不对外暴露
func (h *Handler) startOneWorker(i uint32, reqqueue chan ginterface.IRequest) {
	fmt.Println("the worker num is ", i)
	for {
		select {
		case req := <-reqqueue:
			h.DoMsgHandler(req)
		}
	}

}

func (h *Handler) SendMsgToTaskQueue(request ginterface.IRequest) {
	//消息平均分配给不同worker
	i := request.GetConnection().GetConnectionID() % h.WorkerPoolSize

	//将消息发送给对应的worker的TaskQueue
	h.TaskQueue[i] <- request
}
