package main

import (
	"fmt"
	"go-tcp/ginterface"
	"go-tcp/gnet"
)

type Testrouter struct {
	gnet.BaseRouter
}

func (this *Testrouter) PreHandler(r ginterface.IRequest) {
	fmt.Println("This is a new request, handle by router1, request ID is", r.GetId())

}
func (this *Testrouter) Handler(r ginterface.IRequest) {
	fmt.Println("handling the new request")
	r.GetConnection().SendMsg(0, []byte(fmt.Sprintf("hello client, you ID is %d", r.GetId())))

}
func (this *Testrouter) PostHandler(r ginterface.IRequest) {
	fmt.Println("bye bye!")
}

type Testrouter_2 struct {
	gnet.BaseRouter
}

func (this *Testrouter_2) PreHandler(r ginterface.IRequest) {
	fmt.Println("This is a new request, handle by router2, request ID is", r.GetId())

}
func (this *Testrouter_2) Handler(r ginterface.IRequest) {
	fmt.Println("handling the new request")
	r.GetConnection().SendMsg(0, []byte(fmt.Sprintf("hello client, you ID is %d", r.GetId())))

}
func (this *Testrouter_2) PostHandler(r ginterface.IRequest) {
	fmt.Println("bye bye!")
}

func main() {

	server := gnet.NewServer()
	server.SetOnConnStart(func(c ginterface.IConnection) {
		c.SendMsg(1, []byte("start!"))
	})
	server.SetOnConnStop(func(c ginterface.IConnection) {
		c.SendMsg(1, []byte("stop!"))
	})
	server.AddRouter(1, &Testrouter{})
	server.AddRouter(2, &Testrouter_2{})
	server.Serve()
}
