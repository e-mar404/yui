package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type pkg struct {
	name    string
	version string
}

type pkgList struct {
	pkgs []pkg
}

func (p pkgList) Init() tea.Cmd {
	return nil
}

func (p pkgList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p pkgList) View() tea.View {
	var v tea.View
	var b strings.Builder

	for _, pkg := range p.pkgs {
		line := fmt.Sprintf("%s | %s\n", pkg.name, pkg.version)
		b.WriteString(line)
	}

	v.SetContent(b.String())

	return v
}

func (p *pkgList) SetPkgs(pkgs []pkg) {
	p.pkgs = pkgs
}

func newPkgList() pkgList {
	return pkgList{}
}
