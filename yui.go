package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type yui struct {
	title      string
	menus      []menu
	pkgLists   []pkgList
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

		// TODO: calculate height of header and footer and subtract that from height (assuming 5 for now)
		for _, menu := range y.menus {
			y.pkgLists[menu].SetHeight(height - 5)
		}

	case pacmanMsg:
		if msg.err != nil {
			return y, tea.Quit
		}

		switch msg.pkgType {
		case all:
			y.pkgLists[all].SetPkgs(msg.pkgs)

		case explicit:
			y.pkgLists[explicit].SetPkgs(msg.pkgs)

		case aur:
			y.pkgLists[aur].SetPkgs(msg.pkgs)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			y.activeMenu = int(all)
			return y, loadAllInstalledPkgs

		case "e":
			y.activeMenu = int(explicit)
			return y, loadExplicitlyInstalledPkgs

		case "u":
			y.activeMenu = int(aur)
			return y, loadAURInstalledPkgs

		case "q", "ctrl+c":
			return y, tea.Quit
		}
	}

	model, cmd := y.pkgLists[y.activeMenu].Update(msg)
	y.pkgLists[y.activeMenu] = model.(pkgList)

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
	for i, item := range y.menus {
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
	return y.pkgLists[y.activeMenu].View().Content
}

func NewYui() yui {
	menu := newMenu()

	var pkgLists []pkgList
	for range len(menu) {
		pkgLists = append(pkgLists, newPkgList())
	}

	return yui{
		title:      "yui",
		styles:     defaultStyles(),
		pkgLists:   pkgLists,
		menus:      menu,
		activeMenu: 0,
	}
}
