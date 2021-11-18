package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//1.注册路由 router
	//xxxx/user ===> func1
	//xxxx/name ===> func2
	//xxxx/id ===> func3

	//https://127.0.0.1:8080/user，func是回调函数，用于路由的响应，这个回调函数原型是固定的
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request：===> 包含客户端发来的数据
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		//这里是具体处理业务逻辑

		//write：===> 通过writer将数据返回给客户端
		//func WriteString(w Writer, s string) (n int, err error)
		_, _ = io.WriteString(writer, "这是/user请求返回的数据！")
	})

	//https://127.0.0.1:8080/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/name请求返回的数据！")
	})

	//https://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/id请求返回的数据！")
	})

	fmt.Println("Http Server start ...")
	//func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}
