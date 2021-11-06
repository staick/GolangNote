package main

import (
	"fmt"
	"os"
)

func main() {
	//1.延迟，关键字，可以用于修饰语句，函数，确保这条语句可以在当前栈推出的时候执行
	//2.一般用于做资源清理工作
	//3.解锁，关闭文件
	//4.在同一个函数多次调用defer，执行时类似于栈的机制：先入后出
	filename := "01-switch.go"
	readFile(filename)
}

func readFile(filename string) {
	//func Open(name String) (*File, error) {
	//1.go语言一般会将错误码作为最后一个参数返回
	//2.err一般nil代表没有错误，执行成功，非nil表示执行失败
	f1, err := os.Open(filename)

	//匿名函数，没有名字，属于一次性逻辑 ==>lamada表达式
	defer func () {
		fmt.Println("准备关闭文件！")
		_=f1.Close()
	}()	//创建一个匿名函数，同时调用

	//defer f1.Close()

	if err != nil {
		fmt.Println("os.Open(\"01-switch.go\") ==> 打开文件失败，err:", err)
		return
	}

	buf := make([]byte, 1024)
	//func (f *File) Read(b []byte) (n int, err error)
	n, _ := f1.Read(buf)
	fmt.Println("读取文件的实际长度：", n)
	fmt.Println("读取文件内容：", string(buf))
}
