package gnet

import (
	"errors"
	"fmt"
	"go-tcp/ginterface"
	"go-tcp/utils"
	"net"
)

type Server struct {
	Name      string
	Ip        string
	Port      int
	IpVersion string
	Handler   ginterface.IHandler
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back error", err)
		return errors.New("CallBack error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("%s is working!\n", s.Name)
	//1.获取TCP地址
	go func() {
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr error", err)
			return
		}
		//2.监听地址
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP error", err)
			return
		}
		fmt.Println("start success")
		var cid uint32
		cid = 0
		//3.业务处理
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error", err)
				continue
			}
			dealconn := NewConnection(conn, cid, s.Handler)
			cid++
			go dealconn.Start()

		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(msgID uint32, router ginterface.IRouter) {
	s.Handler.AddRouter(msgID, router)
}

func NewServer() ginterface.IServer {
	res := &Server{
		Name:      utils.GlobalObject.Name,
		Ip:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		IpVersion: "tcp4",
		Handler:   NewHandler(),
	}
	return res

}
