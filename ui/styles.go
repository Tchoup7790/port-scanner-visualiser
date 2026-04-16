// Package ui
package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#0000FF")).
			Padding(0, 1)

	progressStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A1A1A1"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			MarginTop(1)

	openPortStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	emptyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))
)
