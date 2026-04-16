package main

import (
	"fmt"
	"os"

	"port-scanner-visualiser/config"
	"port-scanner-visualiser/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	host := config.GetFlags()

	p := tea.NewProgram(ui.InitialModel(host, 1024))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
