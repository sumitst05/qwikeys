package game

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type Game struct {
	WordListPath string
	WordCount    int
}

func NewGame(wordListPath string, wordCount int) *Game {
	return &Game{
		WordListPath: wordListPath,
		WordCount:    wordCount,
	}
}

func (game *Game) Run() (string, error) {
	words, err := ReadWordsFromFile(game.WordListPath)
	if err != nil {
		return "", err
	}

	text := GenerateRandomString(words, game.WordCount)
	fmt.Print("\033[90m") // Set text color to gray
	fmt.Print(text)

	err = keyboard.Open()
	if err != nil {
		return "", err
	}
	defer keyboard.Close()

	startTime := time.Now()
	inaccuracy := 0
	input := ""

	interrupt := false
	for {
		charInput, key, err := keyboard.GetKey()
		if err != nil {
			return "", err
		}

		if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		} else {
			if len(input) < len(text) && charInput != rune(text[len(input)]) {
				inaccuracy++
			}
			input += string(charInput)
		}

		Colorize(text, input)

		if key == keyboard.KeyCtrlC {
			interrupt = true
			break
		}

		if len(input) == len(text) {
			break
		}
	}

	if interrupt {
		fmt.Println("Interrupt recieved. Exiting...")
		os.Exit(1)
	}

	duration := time.Since(startTime).Seconds()
	speed := float64(game.WordCount) / (duration / 60)
	accuracy := 100.0 - (float64(inaccuracy) / float64(len(text)) * 100)

	result := fmt.Sprintf("\n\nwpm: %v\n", int(speed)) +
		fmt.Sprintf("accuracy: %v%%\n", int(accuracy)) +
		fmt.Sprintf("time: %vs\n", int(duration))

	return result, nil
}
