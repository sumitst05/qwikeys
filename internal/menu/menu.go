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
	GameModes       []MenuItem
	Themes          []MenuItem
	Players         []MenuItem
	SelectedMode    int
	SelectedTheme   int
	SelectedPlayers int
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
		Players: []MenuItem{
			{Label: "2"},
			{Label: "3"},
			{Label: "4"},
			{Label: "5"},
		},
		SelectedMode:    0,
		SelectedTheme:   0,
		SelectedPlayers: 0,
	}
}

func (m *Menu) Print(items []MenuItem, selectedIndex int) {
	fmt.Print("\033[94m")
	fmt.Print("QwiKeys\n\n")

	fmt.Print("\033[94m")
	fmt.Println("↑ / Ctrl+W   ↓ / Ctrl+S")

	fmt.Println("Select Mode:")
	if m.SelectedMode != 0 && m.SelectedPlayers == 0 && m.SelectedTheme == 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("\033[94m")
		fmt.Print("QwiKeys\n\n")
		fmt.Print("\033[94m")
		fmt.Println("↑ / Ctrl+W   ↓ / Ctrl+S")
		fmt.Println("Select Number of Players:")
	}

	if m.SelectedPlayers != 0 && m.SelectedTheme == 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("\033[94m")
		fmt.Print("QwiKeys\n\n")
		fmt.Print("\033[94m")
		fmt.Println("↑ / Ctrl+W   ↓ / Ctrl+S")
		fmt.Println("Select Theme:")
	}

	fmt.Print("\033[39m")
	for i, item := range items {
		if i == selectedIndex {
			fmt.Printf("> %s\n", item.Label)
		} else {
			fmt.Printf(" %s\n", item.Label)
		}
	}
}

func (m *Menu) Select(items []MenuItem, selectedIndex int) int {
	m.Print(items, selectedIndex)

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
		m.Print(items, selectedIndex)

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

func (m *Menu) SelectPlayers() {
	m.SelectedPlayers = m.Select(m.Players, m.SelectedPlayers)
}

func (m *Menu) SelectMode() {
	m.SelectedMode = m.Select(m.GameModes, m.SelectedMode)

	if m.SelectedMode == 1 {
		m.SelectPlayers()
	}
}

func (m *Menu) SelectTheme() {
	m.SelectedTheme = m.Select(m.Themes, m.SelectedTheme)
}

func (m *Menu) Display() (int, int, int) {
	m.SelectMode()

	fmt.Print("\033[H\033[2J")

	m.SelectTheme()

	return m.SelectedMode, m.SelectedTheme, m.SelectedPlayers
}
