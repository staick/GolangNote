package main

import "fmt"

//编写一个函数，它调用自己10次再退出

func selfLoop(n int) {
	if n <= 0 {
		fmt.Println("循环结束，退出！")
		return
	}
	fmt.Printf("循环第%d\n", 11-n)
	n -= 1
	selfLoop(n)
}

func main() {
	selfLoop(10)
}