package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	//http://127.0.0.1:8080/user
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/user返回的数据")
	})

	//http://127.0.0.1:8080/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/name请求返回的数据")
	})

	//http://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/id请求返回的数据")
	})

	fmt.Println("http server start...")
	//func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}
