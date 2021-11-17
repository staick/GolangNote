package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//1.确定监听地址
	ip := "127.0.0.1"
	port := 4321
	address := fmt.Sprintf("%s:%d", ip, port)

	//2.创建监听
	//func Listen(network, address string) (Listener, error)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	fmt.Println("监听中...")

	//3.接受连接
	//Accept() (Conn, error)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}

	fmt.Println("连接建立成功！")

	//4.创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024)

	//5.接收读取到的数据
	//Read(b []byte) (n int, err error)
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("Client ====> Server，长度：", cnt, "，数据：", string(buf[0:cnt]))
	//6.对数据进行处理
	//func ToUpper(s string) string
	upperData := strings.ToUpper(string(buf[0:cnt]))

	//Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("Client <==== Server，长度：", cnt, "，数据：", upperData)

	//7.关闭连接
	_ = conn.Close()
}
