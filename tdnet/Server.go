package tdnet

import (
	"fmt"
	"net"
	"strconv"
	"tomdog/tdface"
	"tomdog/utils"
)

type Server struct {
	// 服务器的名称
	Name string
	// IPV4或者其他
	IPVersion string
	// 服务器绑定的 IP 地址，点分十进制表示
	IP string
	// 服务器绑定的端口
	Port int
	// 消息处理器
	msgHandler tdface.IMsgHandler
}

func (s *Server) AddRouter(router tdface.IRouter) {
	s.msgHandler.AddRouter(router)
}

func (s *Server) Start() {
	utils.Logging("[starting] server listenner at IP:PORT " + s.IP + "" + strconv.Itoa(s.Port) + ", is starting")

	go func() {
		// 1、获取一个 ip 地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		// 2、监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, "err:", err)
			return
		}
		// 已经监听成功
		utils.Logging("[listening] start tomdog server " + s.Name + " success, now listening")

		// sever.go 应该有一个自动生成 connID 的方法，并且生成的ID应该满足要求
		var cid uint32
		cid = 0

		// 3、启动网络连接
		for {
			// 3.1 阻塞等待客户端建立连接请求
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				// 获取连接失败
				fmt.Println("accept error")
				continue
			}

			// 3.2 todo Server.Start() 设置服务器最大连接控制，如果超过最大连接，则关闭最新的链接

			// 3.3 Server.Start() 处理该信链接请求的业务方法
			dealConn := NewConnection(tcpConn, cid, s.msgHandler)
			cid++

			go dealConn.Start()
		}
		//	end for
	}()

}

func (s *Server) Stop() {
	utils.Logging("[stop] tomdog server ,name " + s.Name)
	// todo 关闭资源
}

func (s *Server) Serve() {
	s.Start()

	// todo Server.Serve
	select {}
}

func NewServer() tdface.IServer {

	utils.GlobalObject.Init()

	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandler(),
	}

	return s
}
