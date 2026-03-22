package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type yui struct {
	title      string
	menu       []menu
	activeMenu int
	styles     styles
}

func (y yui) Init() tea.Cmd {
	return nil
}

func (y yui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width := msg.Width

		_, rightHeaderPadding, _, leftHeaderPadding := y.styles.header.GetPadding()
		_, rigthTitlePadding, _, leftTitlePadding := y.styles.title.GetPadding()

		menuWidth := width -
			(leftHeaderPadding + rightHeaderPadding) -
			(leftTitlePadding + rigthTitlePadding + len(y.title))

		y.styles.header = y.styles.header.Width(width)
		y.styles.menu = y.styles.menu.Width(menuWidth)

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

	return y, cmd
}

func (y yui) View() string {
	return fmt.Sprintf("%s\n", y.headerView())
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

func NewYui() yui {
	return yui{
		title:      "yui",
		styles:     defaultStyles(),
		menu:       []menu{all, explicit, aur},
		activeMenu: 0,
	}
}
