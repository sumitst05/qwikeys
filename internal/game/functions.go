package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func GenerateRandomString(words []string, wordCount int) string {
	text := ""
	for i := 0; i < wordCount; i++ {
		r := rand.Intn(len(words))
		text += words[r]
		if i != wordCount-1 {
			text += " "
		}
	}

	return text
}

func Colorize(text, input string, timeLimit int, startTime time.Time) {
	output := ""
	for i, char := range text {
		if i < len(input) {
			if char == rune(input[i]) {
				output += fmt.Sprintf("\033[32m%s\033[0m", string(char)) // Green color for correct input
			} else {
				output += fmt.Sprintf("\033[31m%s\033[0m", string(char)) // Red color for incorrect input
			}
		} else {
			output += "\033[90m" + string(char) + "\033[0m" // Gray for remaining characters
		}
	}

	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Print("\033[94m")      // Blue for title
	fmt.Print("QwiKeys\n\n")

	fmt.Println(output)

	if timeLimit > 0 {
		go func() {
			for {
				remainingTime := timeLimit - int(time.Since(startTime).Seconds())
				fmt.Printf("\rTime Left: %ds", remainingTime)
				time.Sleep(time.Second)
				if remainingTime == 0 {
					return
				}
			}
		}()
	}
}
