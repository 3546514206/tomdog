package tdface

import "net"

type IConnection interface {

	// Start 启动连接，当前连接开始工作
	Start()

	// Stop 停止连接，结束当前连接状态
	Stop()

	// GetTCPConnection 从当前的连接中获取原始的 TCPConn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前连接 ID
	GetConnID() uint32

	// RemoteAddr 获取远端客户端地址
	RemoteAddr() net.Addr

	SendMsg(msgId uint32, data []byte) error
}

// HandFunc 定义一个统一处理连接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
