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
	GameModes     []MenuItem
	Themes        []MenuItem
	Players       int
	SelectedMode  int
	SelectedTheme int
}

func NewMenu() *Menu {
	return &Menu{
		GameModes: []MenuItem{
			{Label: "Single Player"},
			{Label: "Multi Player"},
		},
		Themes: []MenuItem{
			{Label: "Time Rush"},
			{Label: "Word Sprint"},
		},
		SelectedMode:  0,
		SelectedTheme: 0,
	}
}

func (m *Menu) Print(items []MenuItem, selectedIndex int) {
	fmt.Print("\033[94m")
	fmt.Println("QwiKeys")

	fmt.Print("\033[39m")
	for i, item := range items {
		if i == selectedIndex {
			fmt.Printf("> %s\n", item.Label)
		} else {
			fmt.Printf(" %s\n", item.Label)
		}
	}
}

func (m *Menu) Select(items []MenuItem, selectedIndex int) {
	m.Print(items, selectedIndex)

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			selectedIndex = (selectedIndex - 1 + len(items)) % len(items)
		case keyboard.KeyArrowDown:
			selectedIndex = (selectedIndex + 1) % len(items)
		case keyboard.KeyEnter:
			return
		case keyboard.KeyCtrlC:
			fmt.Println("\nInterrupt received. Exiting...")
			os.Exit(130)
		}

		// Clear the screen and display the updated menu
		fmt.Print("\033[H\033[2J")
		m.Print(items, selectedIndex)

		// Reset foreground color
		fmt.Print("\033[39m")

		// Move cursor to beginning of the line
		fmt.Print("\033[G")
	}
}

func (m *Menu) SelectMode() {
	m.Select(m.GameModes, m.SelectedMode)
}

func (m *Menu) SelectTheme() {
	m.Select(m.Themes, m.SelectedTheme)
}

func (m *Menu) Display() (int, int) {
	m.SelectMode()
	fmt.Print("\033[H\033[2J")
	m.SelectTheme()
	return m.SelectedMode, m.SelectedTheme
}
