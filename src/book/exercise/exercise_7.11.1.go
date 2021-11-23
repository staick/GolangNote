package main

import "fmt"

type Game struct {
	name     string
	date     string
	platform string
}

func main() {
	game := Game{
		name:     "Tetris",
		date:     "2021-11-23",
		platform: "dos",
	}

	fmt.Printf("name:%s, date:%s, platform:%s\n", game.name, game.date, game.platform)
}