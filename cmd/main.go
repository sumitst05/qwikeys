package main

import (
	"fmt"
	"qwikeys/internal/game"
	"qwikeys/internal/menu"
)

func main() {
	menu := menu.NewMenu()
	mode, theme, players := menu.Display()

	fmt.Println(mode, theme, players)

	// Clear screen
	fmt.Print("\033[H\033[2J")

	game := game.NewGame("./pkg/words.txt", 15)

	result, err := game.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
