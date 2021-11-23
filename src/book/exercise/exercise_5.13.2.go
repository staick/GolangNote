package main

import "fmt"

func chargeNum(num int) bool {
	if num > 5 && num < 10 {
		return true
	}
	return false
}

func main() {
	fmt.Println("这个数是否大于5小于10:", chargeNum(8))
}