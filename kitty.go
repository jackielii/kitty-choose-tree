package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
)

type KittyWindow struct {
	ID      int      `json:"id,omitempty"`
	Cmdline []string `json:"cmdline,omitempty"`
	Cwd     string   `json:"cwd,omitempty"`
	// ForegroundProcess
	IsActive  bool   `json:"is_active,omitempty"`
	IsFocused bool   `json:"is_focused,omitempty"`
	Title     string `json:"title,omitempty"`
}

func (k KittyWindow) Value() string {
	return fmt.Sprintf("\t\t%s", k.Title)
}

func (k KittyWindow) FilterValue() string {
	return k.Title
}

type KittyTab struct {
	ID        int           `json:"id,omitempty"`
	IsActive  bool          `json:"is_active,omitempty"`
	IsFocused bool          `json:"is_focused,omitempty"`
	Title     string        `json:"title,omitempty"`
	Windows   []KittyWindow `json:"windows,omitempty"`
}

func (k KittyTab) Value() string {
	return fmt.Sprintf("\t%s", k.Title)
}

func (k KittyTab) FilterValue() string {
	return k.Title
}

type KittyOSWindow struct {
	ID               int        `json:"id,omitempty"`
	IsFocused        bool       `json:"is_focused,omitempty"`
	IsActive         bool       `json:"is_active,omitempty"`
	PlatformWindowID int        `json:"platform_window_id,omitempty"`
	Tabs             []KittyTab `json:"tabs,omitempty"`
}

func (k KittyOSWindow) Value() string {
	// TODO: assign a name to the OS window similar to Tmux Session name
	return fmt.Sprintf("%v", k.ID)
}

func (k KittyOSWindow) FilterValue() string {
	return ""
}

func createItems() []list.Item {
	out, err := exec.Command("kitty", "@", "ls").CombinedOutput()
	if err != nil {
		panic(err)
	}
	var v []KittyOSWindow
	err = json.Unmarshal(out, &v)
	if err != nil {
		panic(err)
	}
	res := make([]list.Item, 0)
	for _, w := range v {
		res = append(res, w)
		for _, t := range w.Tabs {
			if len(t.Windows) == 1 && t.Windows[0].Title == "kitty-choose-tree" {
				continue
			}
			res = append(res, t)
			for _, wi := range t.Windows {
				res = append(res, wi)
			}
		}
	}
	return res
}

func focusTab(tab KittyTab) {
	var activeWindow *KittyWindow
	for _, w := range tab.Windows {
		if w.IsActive {
			activeWindow = &w
		}
	}
	if activeWindow == nil {
		panic("no active window")
	}

	focusWindow(*activeWindow)
}

func focusWindow(win KittyWindow) {
	err := exec.Command("kitty", "@", "focus-window", "-m", fmt.Sprintf("id:%d", win.ID)).Run()
	if err != nil {
		panic(err)
	}
}
