package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//要求，服务端在注册rpc对象时，能让编译器检测出注册对象是否合法

//创建接口，在接口中定义方法的原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

//调用该方法时，需要给i传参，参数应该是实现了HelloWorld方法的类对象
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

/*------------------------------------------------------------*/

//像调用本地函数一样，调用远程函数
//定义类
type Myclient struct {
	c *rpc.Client
}

//由于使用了c调用Call，因此需要初始化c
func InitClient (addr string) {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{c:conn}
}

//实现函数，原型参照上面的interface来实现
func (this *Myclient) HelloWorld (a string, b *string) error {
	return this.c.Call("hello.HelloWorld", a, b)
}