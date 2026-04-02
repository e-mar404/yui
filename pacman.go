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

	return pacmanMsg{
		pkgType: all,
		pkgs:    parsePkgs(out),
		err:     err,
	}
}

func loadExplicitlyInstalledPkgs() tea.Msg {
	out, err := exec.Command("pacman", "-Qe").Output()

	return pacmanMsg{
		pkgType: explicit,
		pkgs:    parsePkgs(out),
		err:     err,
	}
}

func loadAURInstalledPkgs() tea.Msg {
	out, err := exec.Command("pacman", "-Qm").Output()

	return pacmanMsg{
		pkgType: aur,
		pkgs:    parsePkgs(out),
		err:     err,
	}
}

func parsePkgs(in []byte) []pkg {
	pkgs := []pkg{}
	for line := range strings.SplitSeq(string(in), "\n") {
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

	return pkgs
}
