package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.rpc连接服务器
	conn, err := rpc.Dial("tcp", ":8800")
	if err != nil {
		fmt.Println("rpc.Dial err:", err)
		return
	}
	defer conn.Close()

	//2.调用远程函数
	var reply string	//接收返回值
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("conn.Call err:", err)
		return
	}

	fmt.Println(reply)
}