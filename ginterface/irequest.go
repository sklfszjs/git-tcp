package ginterface

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
