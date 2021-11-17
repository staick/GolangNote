package main

import (
	"fmt"
	"net"
)

func main() {
	//1.建立连接
	ip := "127.0.0.1"
	port := 4321
	address := fmt.Sprintf("%s:%d", ip, port)

	//func Dial(network, address string) (Conn, error)
	//ip地址默认为127.0.0.1
	//net.Dial("tcp",":4321")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	fmt.Println("client与server连接建立成功！")

	//2.向Server发送数据
	sendData := []byte("helloworld")
	//Write(b []byte) (n int, err error)
	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("Client ====> Server cnt:", cnt, ", data:", string(sendData))

	//3.接收服务器返回的数据
	buf := make([]byte, 1024)
	//Read(b []byte) (n int, err error)
	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("Client <==== Server cnt:", cnt, ", data:", string(buf[0:cnt]))

	//4.关闭连接
	conn.Close()
}
