package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//定义类对象
type World struct {

}

//绑定类方法
func (this *World) HelloWorld (name string, resp *string) error {
	*resp = name + ",你好！"
	return nil
}

func main() {
	//1.注册RPC服务，绑定对象
	// func rpc.RegisterName(name string, rcvr interface{}) error
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc服务失败！", err)
		return
	}
	//2.设置监听
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听...")
	//3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("lintener.Accept err", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接建立成功...")

	//4.绑定服务
	rpc.ServeConn(conn)
}