package main

import (
	"fmt"
	"qwikeys/internal/game"
	"qwikeys/internal/menu"
)

func main() {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	menu := menu.NewMenu()
	options := menu.Display()

	// Clear screen
	fmt.Print("\033[H\033[2J")

	game := game.NewGame("./pkg/words.txt", options.WordCount, options.TimeControl, options.Players)

	result, err := game.Run()
	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}
