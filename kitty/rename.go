package kitty

import (
	"fmt"
	"os/exec"
)

func renameOSWindow(win KittyOSWindow, name string) {
	err := exec.Command("kitty", "@", "rename-window", "-m", fmt.Sprintf("id:%d", win.ID), name).Run()
	if err != nil {
		panic(err)
	}
}

func renameTab(tab KittyTab, name string) {
	err := exec.Command("kitty", "@", "rename-tab", "-m", fmt.Sprintf("id:%d", tab.ID), name).Run()
	if err != nil {
		panic(err)
	}
}

func renameWindow(win KittyWindow, name string) {
	err := exec.Command("kitty", "@", "rename-window", "-m", fmt.Sprintf("id:%d", win.ID), name).Run()
	if err != nil {
		panic(err)
	}
}

func Rename(item Item, name string) {
	switch item := item.(type) {
	case KittyOSWindow:
		renameOSWindow(item, name)
	case KittyTab:
		renameTab(item, name)
	case KittyWindow:
		renameWindow(item, name)
	}
}
