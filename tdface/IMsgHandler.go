package tdface

// IMsgHandler 消息管理抽象层
type IMsgHandler interface {

	// DoMsgHandler 处理消息
	DoMsgHandler(request IRequest)

	// AddRouter 添加路由
	AddRouter(router IRouter)
}
