package test

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestClientSample(t *testing.T) {
	fmt.Println("client test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7077")
	if err != nil {
		fmt.Println("client start error,exit!")
		return
	}

	index := 0
	// 这个单测中客户端一直发，服务端一直收，显然是没有停止标识的，而且客户端服务端都约定了每次
	for true {

		value := "hello " + strconv.Itoa(index)
		_, err = conn.Write([]byte(value))
		index++

		if err != nil {
			fmt.Println("write error,err:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Println("read error, exit, error", err)
		}

		fmt.Printf("server call back now,the content is %s, lenth is %d \n", buf, cnt)
	}
}
