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
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "next tab"),
		),
		prevTab: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "previous tab"),
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

func (k keymaps) disable() {
	k.nextTab.SetEnabled(false)
	k.prevTab.SetEnabled(false)
	k.nextOSWindow.SetEnabled(false)
	k.prevOSWindow.SetEnabled(false)
	k.rename.SetEnabled(false)
	k.sel.SetEnabled(false)
	k.quit.SetEnabled(false)
}

func (k keymaps) enable() {
	k.nextTab.SetEnabled(true)
	k.prevTab.SetEnabled(true)
	k.nextOSWindow.SetEnabled(true)
	k.prevOSWindow.SetEnabled(true)
	k.rename.SetEnabled(true)
	k.sel.SetEnabled(true)
	k.quit.SetEnabled(true)
}
