package gnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"go-tcp/ginterface"
	"go-tcp/utils"
)

type DataPackage struct {
}

func NewDataPackage() ginterface.IDataPackage {
	return &DataPackage{}
}

func (d *DataPackage) GetHeadLen() uint32 {
	return 8
}

func (d *DataPackage) Pack(message ginterface.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuf, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuf, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuf, binary.LittleEndian, message.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}

// 这里拆包不取数据部分，只看id和长度。
func (d *DataPackage) Unpack(datapackage []byte) (ginterface.IMessage, error) {
	dataBuf := bytes.NewReader(datapackage)
	msg := &Message{}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data")
	}
	return msg, nil
}
