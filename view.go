package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jackielii/kitty-choose-tree/kitty"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(kitty.Item)
	if !ok {
		return
	}

	str := i.Value()
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

type model struct {
	list list.Model
	keys keymaps

	typing bool
}

func newModel(items []list.Item) model {
	l := list.New(items, itemDelegate{}, 0, 0)
	// l.Title = ""
	l.SetShowTitle(false)
	l.SetHeight(-1)
	l.SetShowStatusBar(false)
	// l.SetShowPagination(false)
	l.InfiniteScrolling = true
	// l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	keys := newKeymaps()
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.rename,
			keys.sel,
			keys.nextTab,
			keys.prevTab,
			keys.nextOSWindow,
			keys.prevOSWindow,
			// keys.quit,
		}
	}
	return model{list: l, keys: newKeymaps()}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch {
		// case key.Matches(msg, m.keys.quit):
		// 	return m, tea.Quit
		case key.Matches(msg, m.keys.rename):
			// return m, tea.

		case key.Matches(msg, m.keys.sel):
			kitty.Focus(m.list.SelectedItem().(kitty.Item))
			return m, tea.Quit
		case key.Matches(msg, m.keys.nextTab):
			items := m.list.VisibleItems()
			for i := m.list.Cursor() + 1; i < len(items); i++ {
				if _, ok := items[i].(kitty.KittyTab); ok {
					m.list.Select(i)
					break
				}
			}
		case key.Matches(msg, m.keys.prevTab):
			items := m.list.VisibleItems()
			for i := m.list.Cursor() - 1; i >= 0; i-- {
				if _, ok := items[i].(kitty.KittyTab); ok {
					m.list.Select(i)
					break
				}
			}
		case key.Matches(msg, m.keys.nextOSWindow):
			items := m.list.VisibleItems()
			for i := m.list.Cursor() + 1; i < len(items); i++ {
				if _, ok := items[i].(kitty.KittyOSWindow); ok {
					m.list.Select(i)
					break
				}
			}
		case key.Matches(msg, m.keys.prevOSWindow):
			items := m.list.VisibleItems()
			for i := m.list.Cursor() - 1; i >= 0; i-- {
				if _, ok := items[i].(kitty.KittyOSWindow); ok {
					m.list.Select(i)
					break
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	listView := m.list.View()

	return listView
}
