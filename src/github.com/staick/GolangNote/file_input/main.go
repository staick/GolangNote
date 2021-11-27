package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b string
	var c float32
	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}