package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//只读的方式打开文件
	file, err := os.Open("./file_open/file.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	//读取文件内容
	var content []byte
	var buf[128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file:", err)
			return
		}

		content = append(content, buf[:n]...)
	}
	fmt.Println("data:", string(content))
}