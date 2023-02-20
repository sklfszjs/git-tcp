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
	fmt.Println("before\n")

}
func (this *Testrouter) Handler(r ginterface.IRequest) {
	fmt.Println("during\n")
	fmt.Println("id", r.GetId(), "data", string(r.GetData()))
	r.GetConnection().SendMsg(5, []byte("1"))

}
func (this *Testrouter) PostHandler(r ginterface.IRequest) {
	fmt.Println("after\n")
}

type Testrouter_2 struct {
	gnet.BaseRouter
}

func (this *Testrouter_2) PreHandler(r ginterface.IRequest) {
	fmt.Println("before_2\n")

}
func (this *Testrouter_2) Handler(r ginterface.IRequest) {
	fmt.Println("during_2\n")
	fmt.Println("id", r.GetId(), "data", string(r.GetData()))
	r.GetConnection().SendMsg(5, []byte("1"))

}
func (this *Testrouter_2) PostHandler(r ginterface.IRequest) {
	fmt.Println("after_2\n")
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
