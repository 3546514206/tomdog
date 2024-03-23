package tdnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"tomdog/tdface"
	"tomdog/utils"
)

// DataPack 封包拆包实现类
type DataPack struct {
}

// NewDataPack 构造函数
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度
func (d *DataPack) GetHeadLen() uint32 {
	// 包头长度=len(Message.id)+len(Message.DataLen)
	return HeadLen
}

func (d *DataPack) Pack(response tdface.IMessage) ([]byte, error) {
	reponseDataBuf := bytes.NewBuffer([]byte{})

	// 这行代码用于将一个uint32类型的DataLen按照小端字节序写入到指定的缓冲区(reponseDataBuf)中，以便后续进行网络传输或其他操作。
	if err := binary.Write(reponseDataBuf, binary.LittleEndian, response.GetDataLen()); err != nil {
		return nil, err
	}

	// 写消息ID
	if err := binary.Write(reponseDataBuf, binary.LittleEndian, response.GetRouterId()); err != nil {
		return nil, err
	}

	// 写data数据
	if err := binary.Write(reponseDataBuf, binary.LittleEndian, response.GetData()); err != nil {
		return nil, err
	}

	// 返回封包之后的数据
	return reponseDataBuf.Bytes(), nil
}

func (d *DataPack) Unpack(dataByte []byte) (tdface.IMessage, error) {
	dataBuf := bytes.NewBuffer(dataByte)

	msg := &Message{}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.RouterId); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	return msg, nil
}
