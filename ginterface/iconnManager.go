package ginterface

type IConnManager interface {
	AddConnection(IConnection)
	RemoveConnection(IConnection)
	GetConnection(uint32) (IConnection, error)
	Size() int
	ClearConnection()
}
