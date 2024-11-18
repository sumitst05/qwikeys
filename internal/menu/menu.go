package menu

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

type MenuItem struct {
	Label string
}

type Menu struct {
	Themes              []MenuItem
	WordCount           []MenuItem
	TimeControl         []MenuItem
	SelectedTheme       int
	SelectedWordCount   int
	SelectedTimeControl int
}

type GameOptions struct {
	Theme       int
	WordCount   int
	TimeControl int
}

func NewMenu() *Menu {
	return &Menu{
		Themes: []MenuItem{
			{Label: "Time Rush"},
			{Label: "Word Sprint"},
		},
		WordCount: []MenuItem{
			{Label: "10"},
			{Label: "20"},
			{Label: "40"},
			{Label: "80"},
			{Label: "100"},
		},
		TimeControl: []MenuItem{
			{Label: "15 seconds"},
			{Label: "30 seconds"},
			{Label: "45 seconds"},
			{Label: "1 minute"},
			{Label: "2 minutes"},
		},
		SelectedTheme:       0,
		SelectedWordCount:   0,
		SelectedTimeControl: 0,
	}
}

func (m *Menu) clearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[94m")
	fmt.Print("QwiKeys\n\n")
	fmt.Print("\033[94m")
	fmt.Println("↑ / Ctrl+W   ↓ / Ctrl+S")
}

func (m *Menu) Print(items []MenuItem, selectedIndex int, prompt string) {
	fmt.Print("\033[94m")
	fmt.Print("QwiKeys\n\n")

	fmt.Print("\033[94m")
	fmt.Println("↑ / Ctrl+W   ↓ / Ctrl+S")

	m.clearScreen()
	fmt.Println(prompt)

	fmt.Print("\033[39m")
	for i, item := range items {
		if i == selectedIndex {
			fmt.Printf("> %s\n", item.Label)
		} else {
			fmt.Printf(" %s\n", item.Label)
		}
	}
}

func (m *Menu) Select(items []MenuItem, selectedIndex int, prompt string) int {
	m.Print(items, selectedIndex, prompt)

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	interrupt := false
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyCtrlW:
			selectedIndex = (selectedIndex - 1 + len(items)) % len(items)
		case keyboard.KeyArrowUp:
			selectedIndex = (selectedIndex - 1 + len(items)) % len(items)
		case keyboard.KeyCtrlS:
			selectedIndex = (selectedIndex + 1) % len(items)
		case keyboard.KeyArrowDown:
			selectedIndex = (selectedIndex + 1) % len(items)
		case keyboard.KeyEnter:
			return selectedIndex // Return the updated selectedIndex
		case keyboard.KeyCtrlC:
			interrupt = true
			break
		}

		// Clear the screen and display the updated menu
		fmt.Print("\033[H\033[2J")
		m.Print(items, selectedIndex, prompt)

		// Reset foreground color
		fmt.Print("\033[39m")

		// Move cursor to beginning of the line
		fmt.Print("\033[G")

		if interrupt {
			break
		}
	}

	if interrupt {
		fmt.Println("\nInterrupt received. Exiting...")
		os.Exit(130)
	}

	return selectedIndex
}

func (m *Menu) SelectTheme() {
	m.SelectedTheme = m.Select(m.Themes, m.SelectedTheme, "Select Theme:")
	if m.SelectedTheme == 0 {
		m.SelectTimeControl()
		switch m.SelectedTimeControl {
		case 0:
			m.SelectedWordCount = 10
			m.SelectedTimeControl = 15
		case 1:
			m.SelectedWordCount = 20
			m.SelectedTimeControl = 30
		case 2:
			m.SelectedWordCount = 30
			m.SelectedTimeControl = 45
		case 3:
			m.SelectedWordCount = 40
			m.SelectedTimeControl = 60
		case 4:
			m.SelectedWordCount = 80
			m.SelectedTimeControl = 120
		}
	} else if m.SelectedTheme == 1 {
		m.SelectWordCount()
		switch m.SelectedWordCount {
		case 0:
			m.SelectedWordCount = 10
		case 1:
			m.SelectedWordCount = 20
		case 2:
			m.SelectedWordCount = 40
		case 3:
			m.SelectedWordCount = 80
		case 4:
			m.SelectedWordCount = 100
		}
	}
}

func (m *Menu) SelectTimeControl() {
	m.SelectedTimeControl = m.Select(m.TimeControl, m.SelectedTimeControl, "Select Time Control:")
}

func (m *Menu) SelectWordCount() {
	m.SelectedWordCount = m.Select(m.WordCount, m.SelectedWordCount, "Select Word Count:")
}

func (m *Menu) Display() *GameOptions {
	fmt.Print("\033[H\033[2J")

	m.SelectTheme()

	return &GameOptions{
		TimeControl: m.SelectedTimeControl,
		WordCount:   m.SelectedWordCount,
	}
}
