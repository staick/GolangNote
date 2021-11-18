package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1.建立连接
	//func Dial(network, address string) (Conn, error)
	//默认ip为127.0.0.1
	conn, err := net.Dial("tcp", ":4321")
	if err != nil {
		fmt.Println("net.Dail err:", err)
		return
	}
	fmt.Println("client与server连接建立成功！")

	for {
		//2.向Server发送数据
		sendData := []byte("helloworld")
		cnt ,err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err:", err)
		}

		fmt.Println("Client ====> Server cnt:", cnt, ", data: ", string(sendData))

		//3.接收服务器返回的数据
		//创建buf，用于接收服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Client <==== Server cnt:", cnt, ", data: ", string(buf[0:cnt]))
		time.Sleep(1 * time.Second)
	}


}
