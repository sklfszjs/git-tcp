package ginterface

type IMessage interface {
	GetMsgId() uint32
	SetMsgId(uint32)
	GetMsgLen() uint32
	SetMsgLen(uint32)
	GetMsgData() []byte
	SetMsgData([]byte)
}
