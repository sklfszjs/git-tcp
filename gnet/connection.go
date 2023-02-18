package gnet

import (
	"errors"
	"fmt"
	"go-tcp/ginterface"
	"io"
	"net"
)

type Connection struct {
	Conn        *net.TCPConn
	ConnID      uint32
	isClosed    bool
	EXITChannel chan bool
	Handler     ginterface.IHandler
	MsgChan     chan []byte
}

func NewConnection(conn *net.TCPConn, connid uint32, handler ginterface.IHandler) ginterface.IConnection {
	c := &Connection{
		Conn:        conn,
		ConnID:      connid,
		isClosed:    false,
		EXITChannel: make(chan bool, 1),
		Handler:     handler,
		MsgChan:     make(chan []byte),
	}
	return c
}

// 专门用来发送消息的模块
func (c *Connection) StartWriter() {
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send error", err)
				return
			}
			//告知写端已经退出
		case <-c.EXITChannel:
			return
		}
	}
}

func (c *Connection) StartReader() {
	defer fmt.Println("connID", c.ConnID, "remote addr ", c.RemoteAddr().String(), "stoped")
	defer c.Stop()
	for {
		//读取客户端数据
		// buf := make([]byte, 512)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("read error", err)
		// 	continue
		// }
		dp := NewDataPackage()
		headData := make([]byte, dp.GetHeadLen())
		//全读取
		if _, err := io.ReadFull(c.GetConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetMsgData(data)

		req := &Request{
			msg:  msg,
			Conn: c,
		}
		go func(req ginterface.IRequest) {
			c.Handler.DoMsgHandler(req)
		}(req)

	}

}

func (c *Connection) Start() {
	//启动读业务
	//TODO启动写业务
	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = false
	c.Conn.Close()
	//告知写routine退出
	c.EXITChannel <- true
	close(c.EXITChannel)
	close(c.MsgChan)
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

func (c *Connection) SendMsg(id uint32, cont []byte) error {
	if c.isClosed {
		return errors.New("connection close when send msg")
	}
	dp := NewDataPackage()
	msg := NewMessage(id, cont)
	fmt.Printf("%#v", msg)
	binarymsg, err := dp.Pack(msg)
	fmt.Println(binarymsg)
	if err != nil {
		return errors.New("data pack error")
	}
	//读写分离
	c.MsgChan <- binarymsg
	// if _, err := c.Conn.Write(binarymsg); err != nil {
	// 	return errors.New("conn write error")
	// }

	return nil
}
