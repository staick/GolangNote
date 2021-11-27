package main

import (
	"fmt"
)

func testInput() {
	var a int
	var b string
	var c float32

	// fmt.Scanf("%d%s%f", &a, &b, &c)
	// fmt.Scan(&a, &b, &c)
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	testInput()
}