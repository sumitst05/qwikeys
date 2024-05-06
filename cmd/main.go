package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	words, err := readWordsFromFile("./cmd/words.txt")
	if err != nil {
		panic(err)
	}

	wordCount := 10
	text := generateRandomString(words, wordCount)
	printGray(text)

	err = keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	interrupt := false

	startTime := time.Now()
	inaccuracy := 0
	input := ""

	for {
		charInput, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
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

		colorize(text, input)

		if key == keyboard.KeyCtrlC {
			interrupt = true
			break
		}

		if len(input) == len(text) {
			break
		}
	}

	duration := time.Since(startTime).Seconds()
	speed := float64(wordCount) / (duration / 60)
	accuracy := 100.0 - (float64(inaccuracy) / float64(len(text)) * 100)

	if !interrupt {
		fmt.Printf("\n\nwpm: %v\n", int(speed))
		fmt.Printf("accuracy: %v%%\n", int(accuracy))
		fmt.Printf("time: %vs\n", int(duration))
	} else {
		fmt.Println("\nInterrupt recieved... exiting")
	}
}

func generateRandomString(words []string, wordCount int) string {
	text := ""
	for i := 0; i < wordCount; i++ {
		s := rand.Intn(len(words))
		text += words[s]
		if i != wordCount-1 {
			text += " "
		}
	}
	return text
}

func printGray(text string) {
	fmt.Print("\033[90m") // Set text color to gray
	fmt.Print(text)
}

func colorize(text, input string) {
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
	fmt.Print("\r", output) // Rewrite the entire string on same place
}

func readWordsFromFile(filename string) ([]string, error) {
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
