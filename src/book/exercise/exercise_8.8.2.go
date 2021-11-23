package main

import (
	"fmt"
)

type Robot interface {
	talk()
}

type T805 struct {
	Name string
}

func (a *T805) talk() {
	fmt.Println("T805 is starting!")
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) talk() {
	fmt.Println("R2D2 is starting!")
}

func Talk(r Robot) {
	r.talk()
}

func main() {
	t := T805 {
		Name: "t805",
	}
	Talk(&t)

	r := R2D2 {
		Broken: false,
	}
	Talk(&r)
}