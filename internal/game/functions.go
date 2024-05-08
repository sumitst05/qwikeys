package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

func Colorize(text, input string) {
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
	fmt.Print("\r", output)
}
