package gnet

import (
	"errors"
	"go-tcp/ginterface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ginterface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() ginterface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]ginterface.IConnection),
	}
}

func (m *ConnManager) AddConnection(conn ginterface.IConnection) {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	m.connections[conn.GetConnectionID()] = conn
}

func (m *ConnManager) RemoveConnection(conn ginterface.IConnection) {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	delete(m.connections, conn.GetConnectionID())
}

func (m *ConnManager) ClearConnection() {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	for _, conn := range m.connections {
		conn.Stop()
		m.RemoveConnection(conn)
	}
}

func (m *ConnManager) Size() int {
	return len(m.connections)
}

func (m *ConnManager) GetConnection(id uint32) (ginterface.IConnection, error) {
	m.connLock.RLock()
	defer m.connLock.RUnlock()
	if v, ok := m.connections[id]; ok {
		return v, nil
	}
	return nil, errors.New("connection doesnt exist")
}
