package main

import "tomdog/tdnet"

func main() {
	server := tdnet.NewServer()

	server.AddRouter(&tdnet.PingRouter{
		BaseRouter: tdnet.BaseRouter{RouterId: 10001},
	})
	server.AddRouter(&tdnet.TestRouter{
		BaseRouter: tdnet.BaseRouter{RouterId: 10002},
	})

	server.Serve()
}
