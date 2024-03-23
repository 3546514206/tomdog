package tdface

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Serve 开启业务服务方法
	Serve()

	// AddRouter 向服务器中添加路由，供客户端连接使用
	AddRouter(router IRouter)
}
