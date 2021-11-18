package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id		int
	Name	string
	Age		int
	gender	string
}

func main() {
	lily := Student{
		Id:	1,
		Name:	"Lily",
		Age:	18,
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

	var lily2 Student
	//func Unmarshal(data []byte, v interface{}) error
	if err = json.Unmarshal(encodeInfo, &lily2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	fmt.Println("name:", lily2.Name)
	fmt.Println("id:", lily2.Id)
	fmt.Println("Age:", lily2.Id)
	fmt.Println("gender:", lily2.gender)
}
