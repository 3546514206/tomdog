package tdface

// IRouter 路由接口负责路由使用框架整合给当前连接自定的业务处理方法
type IRouter interface {

	// PreHandle 在处理业务之前的钩子方法
	PreHandle(request IRequest)

	// Handle 处理业务的方法
	Handle(request IRequest)

	// AfterHandle 处理完业务之后的钩子方法
	AfterHandle(request IRequest)

	GetRouterId() uint32
}
