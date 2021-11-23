package main

import (
	"fmt"
)

func main() {
	chip := make([]int, 4)
	for i:=0; i<4; i++ {
		chip[i] = i + 1
	}
	fmt.Println(chip)

	chip2 := make([]int, 2)
	copy(chip2, chip[2:4])
	fmt.Println(chip2)
}