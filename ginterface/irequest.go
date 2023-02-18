package ginterface

type IRequest interface {
	GetConnection() IConnection
	GetId() uint32
	GetData() []byte
}
