package kitty

import (
	"fmt"
	"os/exec"
)

func focusOSWindow(win KittyOSWindow) {
	// TODO: reliable way to changed to the active window within the OS window
	var activeWindow *KittyWindow
	for _, t := range win.Tabs {
		for _, w := range t.Windows {
			if w.IsActive {
				activeWindow = &w
				break
			}
		}
	}
	if activeWindow == nil {
		activeWindow = &win.Tabs[0].Windows[0]
	}

	focusWindow(*activeWindow)
}

func focusTab(tab KittyTab) {
	// focus-tab doesn't work if switching from another window. So we focus to the active window.
	var activeWindow *KittyWindow
	for _, w := range tab.Windows {
		if w.IsActive {
			activeWindow = &w
			break
		}
	}
	if activeWindow == nil {
		activeWindow = &tab.Windows[0]
	}

	focusWindow(*activeWindow)
}

func focusWindow(win KittyWindow) {
	focusWindowID(win.ID)
}

func focusWindowID(id int) {
	err := exec.Command("kitty", "@", "focus-window", "-m", fmt.Sprintf("id:%d", id)).Run()
	if err != nil {
		panic(err)
	}
}

func Focus(item Item) {
	switch item := item.(type) {
	case KittyOSWindow:
		focusOSWindow(item)
	case KittyTab:
		focusTab(item)
	case KittyWindow:
		focusWindow(item)
	}
}
