package main

import "fmt"

func testSscanf() {
	var a int
	var b string
	var c float32
	var str string = "88 hello 8.8\n"
	fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	testSscanf()
}