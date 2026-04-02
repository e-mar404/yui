package main

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/paginator"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type pkg struct {
	name    string
	version string
}

type pkgList struct {
	pkgs           []pkg
	cursorPosition int // position in []pkgs (0 indexed)
	paginator      paginator.Model
}

func (p pkgList) Init() tea.Cmd {
	return nil
}

func (p pkgList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if p.cursorPosition < (p.paginator.ItemsOnPage(len(p.pkgs)) - 1) {
				p.cursorPosition += 1
			} else if !p.paginator.OnLastPage() { // if last page is not taken into account an undesired loop will happen
				p.paginator.NextPage()
				p.cursorPosition = 0
			}

		case "k":
			if p.cursorPosition > 0 {
				p.cursorPosition -= 1
			} else if !p.paginator.OnFirstPage() { // if first page is not taken into account an undesired loop will happen
				p.paginator.PrevPage()
				p.cursorPosition = p.paginator.ItemsOnPage(len(p.pkgs)) - 1
			}
		}
	}

	p.paginator, cmd = p.paginator.Update(msg)
	return p, cmd
}

func (p pkgList) View() tea.View {
	var v tea.View
	var b strings.Builder

	start, end := p.paginator.GetSliceBounds(len(p.pkgs))
	for idx, pkg := range p.pkgs[start:end] {
		line := fmt.Sprintf("%s | %s", pkg.name, pkg.version)
		if idx == p.cursorPosition {
			line = lipgloss.NewStyle().Background(lipgloss.Color("180")).Foreground(lipgloss.Black).Render(line)
		}
		b.WriteString(line + "\n")
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
		cursorPosition: 0,
		paginator:      paginator.New(),
	}
}
