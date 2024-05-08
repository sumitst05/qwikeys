package main

import (
	"fmt"
	"os"
	"qwikeys/internal/game"
)

func main() {
	game := game.NewGame("./pkg/words.txt", 15)

	result, err := game.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
