package utils

import (
	"encoding/json"
	"io/ioutil"
	"tomdog/tdface"
)

type GlobalObj struct {

	// 当前 tomdog 的全局 server
	TcpServer tdface.IServer

	// 当前服务器的主机IP
	Host string

	// 当前服务器监听的端口
	TcpPort int

	// 当前服务器名称
	Name string

	// 当前 tomdog 的版本号
	Version string

	// 读取的数据报的最大值
	MaxPacketSize uint32

	// 当前主机允许的最大连接个数
	MaxConn int

	// 是否开启日志
	Log bool
}

// GlobalObject 定义一个全局变量
var GlobalObject = new(GlobalObj)

func (g *GlobalObj) Init() {
	// 设置一些初始值
	g.Host = "0.0.0.0"
	g.TcpPort = 7477
	g.Name = "tomdog"
	g.Version = "V0.4"
	g.MaxPacketSize = 4096
	g.MaxConn = 1200
	g.Log = true

	g.Reload()
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("configs/tomdog.json")
	if err != nil {
		panic(err)
	}

	// 传指针有用么？有用
	err = json.Unmarshal(data, GlobalObject)

	if err != nil {
		panic(err)
	}
}
