package main

import (
	"os/exec"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type pkg struct {
	name    string
	version string
}

type pacmanMsg struct {
	pkgType  menu
	pkgItems []list.Item
	err      error
}

func (p pkg) Title() string {
	return p.name
}

func (p pkg) Description() string {
	return ""
}

func (p pkg) FilterValue() string {
	return p.name
}

func loadAllInstalledPkgs() tea.Msg {
	out, err := exec.Command("pacman", "-Q").Output()

	pkgItems := []list.Item{}
	for line := range strings.SplitSeq(string(out), "\n") {
		if line == "" {
			continue
		}

		s := strings.Split(line, " ")

		pkgItems = append(pkgItems, list.Item(
			pkg{
				name:    s[0],
				version: s[1],
			},
		))
	}

	return pacmanMsg{
		pkgType:  all,
		pkgItems: pkgItems,
		err:      err,
	}
}

func newPkgList() list.Model {
	// TODO: create custom delegate for my use case
	// it should resemble a table as much as possible, i just want it to be a list for the built-in filtering
	return list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
}
