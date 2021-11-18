package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id		int
	Name	string
	Age		int

	//注意，gender是小写的，小写字母开头的，再json编码时会忽略掉
	gender	string
}

func main() {
	//在网络传输的时候，把Student结构体，编码成json字符串，传输===>结构体===>字符串===>编码
	//在接收字符串的时候，需要将字符串转换成结构体，然后操作===>字符串===>结构体===>解码

	lily := Student {
		Id:	1,
		Name:	"Lily",
		Age:	20,
		gender:	"女",
	}

	//编码
	//func Marshal(v interface{}) ([]byte, error)
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	fmt.Println("encodeInfo:", string(encodeInfo))

	//解码
	var lily2 Student
	//func Unmarshal(data []byte, v interface{}) error
	if err = json.Unmarshal(encodeInfo, &lily2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	fmt.Println("name:", lily2.Name)
	fmt.Println("gender:", lily2.gender)
	fmt.Println("age:", lily2.Age)
	fmt.Println("id:", lily2.Id)
}
