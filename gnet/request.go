package gnet

import (
	"go-tcp/ginterface"
)

type Request struct {
	Conn *Connection
	msg  ginterface.IMessage
}

func (r *Request) GetConnection() ginterface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetId() uint32 {
	return r.msg.GetMsgId()
}
