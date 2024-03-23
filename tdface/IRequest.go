package tdface

// IRequest 接口实际上是将客户端请求的连接信息和请求的数据封装在一起
type IRequest interface {

	// GetConnection 获取请求连接信息，返回接口而不是实现类，这体现了里氏代换原则
	GetConnection() IConnection

	// GetData 获取请求消息数据
	GetData() IMessage
}
