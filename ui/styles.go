// Package ui
package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("4")).
			Padding(0, 1)

	progressStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4"))

	doneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			MarginTop(1)

	openPortStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2"))

	emptyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))
)
