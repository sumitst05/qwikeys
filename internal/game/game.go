package game

import (
	"fmt"
	"math"
	"time"

	"github.com/eiannone/keyboard"
)

type Game struct {
	WordListPath string
	WordCount    int
	Time         int
	Players      int
}

func NewGame(wordListPath string, wordCount int, time int, players int) *Game {
	return &Game{
		WordListPath: wordListPath,
		WordCount:    wordCount,
		Time:         time,
		Players:      players,
	}
}

func (g *Game) Run() (string, error) {
	fmt.Print("\033[94m")
	fmt.Print("QwiKeys\n\n")

	words, err := ReadWordsFromFile(g.WordListPath)
	if err != nil {
		return "", err
	}

	text := GenerateRandomString(words, g.WordCount)
	fmt.Print("\033[90m") // Set text color to gray
	fmt.Println(text)

	err = keyboard.Open()
	if err != nil {
		return "", err
	}
	defer keyboard.Close()

	startTime := time.Now()
	correctChars := 0
	totalChars := len(text)
	input := ""

	timer := time.NewTimer(time.Duration(g.Time) * time.Second)

	done := make(chan bool)

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				done <- true
				return
			}
			if key == keyboard.KeyCtrlC {
				fmt.Println("\nInterrupt received. Exiting...")
				done <- true
				return
			}
			if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else {
				if char == rune(text[len(input)]) {
					correctChars++
				}
				input += string(char)
			}
			if len(input) == len(text) {
				Colorize(text, input, g.Time, startTime)
				done <- true
				return
			}
		}
	}()

	for {
		select {
		case <-timer.C:
			duration := time.Since(startTime).Seconds()
			wpm := float64(correctChars) / 5.0 / (duration / 60)
			rawWpm := float64(totalChars) / 5.0 / (duration / 60)
			accuracy := (float64(correctChars) / float64(totalChars)) * 100

			result := fmt.Sprintf("\n\nwpm: %v\n", int(wpm)) +
				fmt.Sprintf("raw: %v\n", int(rawWpm)) +
				fmt.Sprintf("accuracy: %.2f%%\n", accuracy) +
				fmt.Sprintf("time: %vs\n", int(duration))

			return result, nil

		case <-done:
			duration := time.Since(startTime).Seconds()
			wpm := float64(correctChars) / 5.0 / (duration / 60)
			rawWpm := float64(totalChars) / 5.0 / (duration / 60)
			accuracy := (float64(correctChars) / float64(totalChars)) * 100

			result := fmt.Sprintf("\n\nwpm: %v\n", math.Round(wpm)) +
				fmt.Sprintf("raw: %v\n", math.Round(rawWpm)) +
				fmt.Sprintf("accuracy: %v%%\n", math.Round(accuracy)) +
				fmt.Sprintf("time: %vs\n", math.Round(duration))

			return result, nil

		default:
			Colorize(text, input, g.Time, startTime)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
