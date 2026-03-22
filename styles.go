package main

import (
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
)

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

func tableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	return s
}
