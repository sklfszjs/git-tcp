package gnet

import (
	"errors"
	"fmt"
	"go-tcp/ginterface"
	"go-tcp/utils"
	"net"
)

type Server struct {
	Name        string
	Ip          string
	Port        int
	IpVersion   string
	Handler     ginterface.IHandler
	ConnManager ginterface.IConnManager
	OnConnStart func(conn ginterface.IConnection)
	OnConnStop  func(conn ginterface.IConnection)
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
	s.Handler.StartWorkerPool()
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
			//这里判断一下链接是否过多
			if s.ConnManager.Size() > utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}
			dealconn := NewConnection(s, conn, cid, s.Handler)
			cid++
			go dealconn.Start()

		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("server stop")
	s.ConnManager.ClearConnection()
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
		Name:        utils.GlobalObject.Name,
		Ip:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		IpVersion:   "tcp4",
		Handler:     NewHandler(),
		ConnManager: NewConnManager(),
	}
	return res

}

func (s *Server) GetConnManager() ginterface.IConnManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(f func(ginterface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(ginterface.IConnection)) {
	s.OnConnStop = f
}

func (s *Server) CallOnConnStart(conn ginterface.IConnection) {
	if s.OnConnStart == nil {
		fmt.Println("no onconnstart function")
		return
	}
	s.OnConnStart(conn)
}

func (s *Server) CallOnConnStop(conn ginterface.IConnection) {
	if s.OnConnStop == nil {
		fmt.Println("no onconnstop function")
		return
	}
	s.OnConnStop(conn)
}
