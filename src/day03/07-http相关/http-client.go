package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//1.设置请求URL
	client := http.Client{}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("clent.Get err:", err)
		return
	}

	//2.获取请求头的相关信息
	//使用beego,gin等web框架，可以使用更加简单的方式获取下面的信息
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("header : ", resp.Header)

	fmt.Println("content-type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)

	body := resp.Body
	fmt.Println("body 111:", body)
	//body是接口类型，需要使用下面的方法转换成字节类型
	//func ReadAll(r io.Reader) ([]byte, error)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return
	}

	fmt.Println("body string:", string(readBodyStr))
}
