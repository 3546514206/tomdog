package tdface

// IDataPack 封包拆包接口
type IDataPack interface {

	// GetHeadLen 获取包头长度
	GetHeadLen() uint32

	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)

	// Unpack 拆包方法
	Unpack([]byte) (IMessage, error)
}
