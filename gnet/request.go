package gnet

import (
	"go-tcp/ginterface"
)

type Request struct {
	Conn *Connection
	Data []byte
}

func (r *Request) GetConnection() ginterface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}
