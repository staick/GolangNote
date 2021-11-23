package main

import "fmt"

//编写一个函数，它接受2个参数并返回3个值
func in2Out3(a, b int) (c, d, e int) {
	c = a + b
	d = a - b
	e = c + d
	return
}
 
func main() {
	c, d, e := in2Out3(20, 10)
	fmt.Printf("%d, %d, %d\n", c, d, e)
}