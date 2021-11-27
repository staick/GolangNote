package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ioutil, err := ioutil.ReadFile("./ioutil/main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(ioutil))
}