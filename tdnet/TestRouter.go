package tdnet

import (
	"fmt"
	"strconv"
	"tomdog/tdface"
	"tomdog/utils"
)

type TestRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *PingRouter) Handle(request tdface.IRequest) {
	utils.Logging("call test router after handle")
	utils.Logging("recv from client: msgid = " + strconv.Itoa(int(request.GetData().GetRouterId())) + ", data=" + string(request.GetData().GetData()))

	err := request.GetConnection().SendMsg(20002, []byte("tested...tested..tested..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *PingRouter) AfterHandle(request tdface.IRequest) {
	utils.Logging("call test router after handle")
}
