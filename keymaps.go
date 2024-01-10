package main

import "github.com/charmbracelet/bubbles/key"

type keymaps struct {
	nextTab      key.Binding
	prevTab      key.Binding
	nextOSWindow key.Binding
	prevOSWindow key.Binding

	rename key.Binding
	sel    key.Binding
	quit   key.Binding
}

func newKeymaps() keymaps {
	return keymaps{
		nextTab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next tab"),
		),
		prevTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "previous tab"),
		),
		nextOSWindow: key.NewBinding(
			key.WithKeys("J"),
			key.WithHelp("J", "next session"),
		),
		prevOSWindow: key.NewBinding(
			key.WithKeys("K"),
			key.WithHelp("K", "previous session"),
		),
		rename: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "rename"),
		),
		sel: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("ctrl+c, q", "quit"),
		),
	}
}
