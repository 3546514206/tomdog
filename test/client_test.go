package test

import (
	"fmt"
	"net"
	"strconv"
	"syscall"
	"testing"
	"time"
)

// TCP 系统调用的缓冲区大小
func TestTcpBufSize(t *testing.T) {
	// 创建一个临时的 TCP 连接
	conn, err := net.Dial("tcp", "127.0.0.1:7077")
	if err != nil {
		fmt.Printf("Dial error: %v\n", err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// 获取底层文件描述符
	fd, _ := conn.(*net.TCPConn).File()

	// 获取 TCP 缓冲区大小
	rBufferSize, err := syscall.GetsockoptInt(syscall.Handle(int(fd.Fd())), syscall.SOL_SOCKET, syscall.SO_RCVBUF)
	if err != nil {
		fmt.Printf("GetsockoptInt error: %v\n", err)
		return
	}

	sBufferSize, err := syscall.GetsockoptInt(syscall.Handle(int(fd.Fd())), syscall.SOL_SOCKET, syscall.SO_SNDBUF)
	if err != nil {
		fmt.Printf("GetsockoptInt error: %v\n", err)
		return
	}

	fmt.Printf("默认接收缓冲区大小：%d\n", rBufferSize)
	fmt.Printf("默认发送缓冲区大小：%d\n", sBufferSize)

}

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
