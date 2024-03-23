package test

import (
	"fmt"
	"net"
	"testing"
	"time"
	"tomdog/tdnet"
)

func TestServer(t *testing.T) {
	s := tdnet.NewServer()

	go ClientTest()
	s.Serve()
}

func ClientTest() {

	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7077")

	if err != nil {
		fmt.Println("client start err，exit！")
		return
	}

	for {
		_, err := conn.Write([]byte("hello tomdog"))
		if err != nil {
			fmt.Println("write err,exit!")
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server call back %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}

}
