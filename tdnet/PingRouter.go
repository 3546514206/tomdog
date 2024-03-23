package tdnet

import (
	"fmt"
	"strconv"
	"tomdog/tdface"
	"tomdog/utils"
)

type PingRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *TestRouter) Handle(request tdface.IRequest) {
	utils.Logging("recv from client: msgid = " + strconv.Itoa(int(request.GetData().GetRouterId())) + ", data=" + string(request.GetData().GetData()))

	err := request.GetConnection().SendMsg(20001, []byte("pong...pong..pong..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *TestRouter) AfterHandle(request tdface.IRequest) {
	utils.Logging("call ping router after handle")

	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping router\n"))
	//
	//if err != nil {
	//	fmt.Println("after handle  ping router error")
	//}
}
