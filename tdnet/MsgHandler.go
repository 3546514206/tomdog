package tdnet

import (
	"strconv"
	"tomdog/tdface"
	"tomdog/utils"
)

type MsgHandler struct {

	// 接口 map
	apis map[uint32]tdface.IRouter
}

func (m *MsgHandler) DoMsgHandler(request tdface.IRequest) {
	msgid := request.GetData().GetRouterId()
	if _, ok := m.apis[msgid]; !ok {
		panic("no suitable optional route selector")
	}

	m.apis[msgid].PreHandle(request)
	m.apis[msgid].Handle(request)
	m.apis[msgid].AfterHandle(request)
}

func (m *MsgHandler) AddRouter(router tdface.IRouter) {
	// 重复添加路由
	if _, ok := m.apis[router.GetRouterId()]; ok {
		panic("repeat router, routerId =" + strconv.Itoa(int(router.GetRouterId())))
	}

	m.apis[router.GetRouterId()] = router
	utils.Logging("router add success, routerId = " + strconv.Itoa(int(router.GetRouterId())))
}

// NewMsgHandler 构造函数
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		apis: make(map[uint32]tdface.IRouter),
	}
}
