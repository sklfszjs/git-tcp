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
	r.GetConnection().GetConnection().Write([]byte("before\n"))

}
func (this *Testrouter) Handler(r ginterface.IRequest) {
	fmt.Println("during\n")
	r.GetConnection().GetConnection().Write([]byte("during\n"))

}
func (this *Testrouter) PostHandler(r ginterface.IRequest) {
	r.GetConnection().GetConnection().Write([]byte("after\n"))
	fmt.Println("after\n")
}
func main() {

	server := gnet.NewServer()
	server.AddRouter(&Testrouter{})
	server.Serve()
}
