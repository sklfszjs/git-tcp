package ginterface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnection() *net.TCPConn
	GetConnectionID() uint32
	RemoteAddr() *net.TCPAddr
	SendMsg(uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
