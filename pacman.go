package main

import (
	"os/exec"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type pacmanMsg struct {
	pkgType menu
	pkgs    []pkg
	err     error
}

func loadAllInstalledPkgs() tea.Msg {
	out, err := exec.Command("pacman", "-Q").Output()

	pkgs := []pkg{}
	for line := range strings.SplitSeq(string(out), "\n") {
		if line == "" {
			continue
		}

		s := strings.Split(line, " ")

		pkgs = append(pkgs,
			pkg{
				name:    s[0],
				version: s[1],
			},
		)
	}

	return pacmanMsg{
		pkgType: all,
		pkgs:    pkgs,
		err:     err,
	}
}
