package gnet

import (
	"fmt"
	"go-tcp/ginterface"
	"net"
)

type Connection struct {
	Conn        *net.TCPConn
	ConnID      uint32
	isClosed    bool
	EXITChannel chan bool
	Router      ginterface.IRouter
}

func NewConnection(conn *net.TCPConn, connid uint32, router ginterface.IRouter) ginterface.IConnection {
	c := &Connection{
		Conn:        conn,
		ConnID:      connid,
		isClosed:    false,
		EXITChannel: make(chan bool, 1),
		Router:      router,
	}
	return c
}

func (c *Connection) StartReader() {
	defer fmt.Println("connID", c.ConnID, "remote addr ", c.RemoteAddr().String(), "stoped")
	defer c.Stop()
	for {
		//读取客户端数据
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			continue
		}
		req := &Request{
			Data: buf,
			Conn: c,
		}
		go func(req ginterface.IRequest) {
			c.Router.PreHandler(req)
			c.Router.Handler(req)
			c.Router.PostHandler(req)
		}(req)

	}

}

func (c *Connection) Start() {
	//启动读业务
	//TODO启动写业务
	go c.StartReader()
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = false
	c.Conn.Close()
	close(c.EXITChannel)
}

func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnectionID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() *net.TCPAddr {
	return nil
}

func (c *Connection) Send(cont []byte) error {
	return nil
}
