package tdnet

// HeadLen = len(Message.id) + len(Message.DataLen)
var HeadLen uint32 = 8

// Message 采取最经典的 TLU (Type-Len-Value) 封包格式来解决粘包问题
type Message struct {

	// Head
	// 消息 ID
	RouterId uint32

	// 消息长度
	DataLen uint32

	// Body
	// 消息内容
	Data []byte
}

// NewMsgPackage 创建一个消息体
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		RouterId: id,
		DataLen:  uint32(len(data)),
		Data:     data,
	}
}

// GetDataLen
// 在 Go 语言中，当你实现一个接口的方法时，可以选择将接收者声明为指针类型或非指针类型。这两者之间有一些重要的区别：
// 如果方法的接收者声明为指针类型，那么只有指向该类型的指针才能够调用该方法。这意味着你必须使用指针来调用方法，否则会导致编译错误。
// 如果方法的接收者声明为非指针类型，那么该类型的值和指针都可以调用该方法。这意味着你可以使用类型的值或指针来调用方法。
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetRouterId() uint32 {
	return m.RouterId
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
