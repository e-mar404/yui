package main

import (
	"os/exec"
	"strings"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type pkgList []pkg

type pkg struct {
	name    string
	version string
}

type pacmanMsg struct {
	pkgType menu
	list    pkgList
	err     error
}

func loadAllInstalledPkgs() tea.Msg {
	out, err := exec.Command("pacman", "-Q").Output()

	var list pkgList
	for _, line := range strings.Split(string(out), "\n") {
		if line == "" {
			continue
		}

		s := strings.Split(line, " ")

		list = append(list, pkg{
			name:    s[0],
			version: s[1],
		})
	}

	return pacmanMsg{
		pkgType: all,
		list:    list,
		err:     err,
	}
}

func pacmanColumns() []table.Column {
	return []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Version", Width: 20},
	}
}
