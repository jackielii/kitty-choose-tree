package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackielii/kitty-choose-tree/kitty"
)

func main() {
	items, err := kitty.CreateItems()
	if err != nil {
		log.Fatal(err)
	}
	m := newModel(items)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
