package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	header           lipgloss.Style
	title            lipgloss.Style
	menu             lipgloss.Style
	activeMenuItem   lipgloss.Style
	inActiveMenuItem lipgloss.Style
}

func defaultStyles() styles {
	return styles{
		header: lipgloss.NewStyle().
			Padding(1),

		title: lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("38")).
			Padding(0, 1),

		menu: lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Right),

		activeMenuItem: lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("183")),

		inActiveMenuItem: lipgloss.NewStyle().
			Padding(0, 1),
	}
}
