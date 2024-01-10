package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := createItems()
	l := list.New(items, itemDelegate{}, 0, 0)
	// l.Title = ""
	l.SetShowTitle(false)
	l.SetHeight(-1)
	l.SetShowStatusBar(false)
	l.InfiniteScrolling = true
	// l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := newModel(l)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
