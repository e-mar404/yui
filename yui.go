package main

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type yui struct {
	title      string
	menu       []menu
	pkgList    list.Model
	activeMenu int
	styles     styles
}

func (y yui) Init() tea.Cmd {
	return loadAllInstalledPkgs
}

func (y yui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height := msg.Width, msg.Height

		_, rightHeaderPadding, _, leftHeaderPadding := y.styles.header.GetPadding()
		_, rigthTitlePadding, _, leftTitlePadding := y.styles.title.GetPadding()

		menuWidth := width -
			(leftHeaderPadding + rightHeaderPadding) -
			(leftTitlePadding + rigthTitlePadding + len(y.title))

		y.styles.header = y.styles.header.Width(width)
		y.styles.menu = y.styles.menu.Width(menuWidth)

		// TODO: instead of a hard coded value find a way to calculate it
		y.pkgList.SetHeight(height - 5)
		y.pkgList.SetWidth(width)

	case pacmanMsg:
		if msg.err != nil {
			return y, tea.Quit
		}

		switch msg.pkgType {
		case all:
			cmds = append(cmds, y.pkgList.SetItems(msg.pkgItems))
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			y.activeMenu = int(all)

		case "e":
			y.activeMenu = int(explicit)

		case "u":
			y.activeMenu = int(aur)

		case "q", "ctrl+c":
			return y, tea.Quit
		}
	}

	y.pkgList, cmd = y.pkgList.Update(msg)
	cmds = append(cmds, cmd)

	return y, tea.Batch(cmds...)
}

func (y yui) View() tea.View {
	var v tea.View
	v.AltScreen = true

	header := y.headerView()
	content := y.contentView()

	v.SetContent(fmt.Sprintf("%s\n%s", header, content))

	return v
}

func (y yui) headerView() string {
	var styledMenuItems []string
	for i, item := range y.menu {
		style := y.styles.inActiveMenuItem
		if y.activeMenu == i {
			style = y.styles.activeMenuItem
		}

		styledMenuItems = append(styledMenuItems, style.Render(item.String()))
	}

	title := y.styles.title.Render(y.title)
	menu := y.styles.menu.Render(
		lipgloss.JoinHorizontal(lipgloss.Left, styledMenuItems...),
	)
	header := fmt.Sprintf("%s%s", title, menu)

	return y.styles.header.Render(header)
}

func (y yui) contentView() string {
	return y.pkgList.View()
}

func NewYui() yui {
	menu := newMenu()

	return yui{
		title:      "yui",
		styles:     defaultStyles(),
		pkgList:    newPkgList(),
		menu:       menu,
		activeMenu: 0,
	}
}
