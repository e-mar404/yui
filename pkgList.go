package main

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/paginator"
	tea "charm.land/bubbletea/v2"
)

type pkg struct {
	name    string
	version string
}

type pkgList struct {
	pkgs      []pkg
	paginator paginator.Model
}

func (p pkgList) Init() tea.Cmd {
	return nil
}

func (p pkgList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.paginator, cmd = p.paginator.Update(msg)
	return p, cmd
}

func (p pkgList) View() tea.View {
	var v tea.View
	var b strings.Builder

	start, end := p.paginator.GetSliceBounds(len(p.pkgs))
	for _, pkg := range p.pkgs[start:end] {
		line := fmt.Sprintf("%s | %s\n", pkg.name, pkg.version)
		b.WriteString(line)
	}

	b.WriteString(p.paginator.View())

	v.SetContent(b.String())

	return v
}

func (p *pkgList) SetPkgs(pkgs []pkg) {
	p.pkgs = pkgs
	p.paginator.SetTotalPages(len(p.pkgs))
}

func (p *pkgList) SetHeight(height int) {
	// TODO: number of items per page will be height - title line - paginator line (assume 2 for now)
	p.paginator.PerPage = height - 2
}

func newPkgList() pkgList {
	return pkgList{
		paginator: paginator.New(),
	}
}
