package kitty

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
)

type Item interface {
	Value() string
	FilterValue() string
}

var _ list.Item = Item(nil)

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
	Title            string     `json:"title,omitempty"`
	IsFocused        bool       `json:"is_focused,omitempty"`
	IsActive         bool       `json:"is_active,omitempty"`
	PlatformWindowID int        `json:"platform_window_id,omitempty"`
	Tabs             []KittyTab `json:"tabs,omitempty"`
}

func (k KittyOSWindow) Value() string {
	// TODO: assign a name to the OS window similar to Tmux Session name
	return fmt.Sprintf("%v", k.Title)
}

func (k KittyOSWindow) FilterValue() string {
	return ""
}

func CreateItems() ([]list.Item, error) {
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
	titles, err := osWindowTitles()
	if err != nil {
		return nil, err
	}
	for _, w := range v {
		w.Title = titles[w.ID]
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
	return res, nil
}

func osWindowTitles() (map[int]string, error) {
	script := `
import json
windows = boss.list_os_windows()
res = dict([(w['id'], get_os_window_title(w['id'])) for w in windows])
answer = json.dumps(res)
`
	res, err := RunKitten(script)
	if err != nil {
		return nil, err
	}
	var v map[int]string
	err = json.Unmarshal([]byte(res), &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}
