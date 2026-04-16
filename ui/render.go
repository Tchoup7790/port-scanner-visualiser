// Package ui
package ui

import (
	"fmt"
	"sort"

	"port-scanner-visualiser/scanner"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	host        string
	totalPorts  int
	scanned     int
	currentPort int
	openPorts   map[int]string
}

type PortScannedMsg struct {
	port      int
	protocole string
	isOpen    bool
}

func InitialModel(host string, totalPorts int) model {
	return model{
		host:        host,
		totalPorts:  totalPorts,
		currentPort: 101,
		scanned:     0,
		openPorts:   make(map[int]string),
	}
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 100)
	for i := range 100 {
		cmds[i] = scanPortCmd(m.host, i+1)
	}
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case PortScannedMsg:
		m.scanned++
		if msg.isOpen {
			m.openPorts[msg.port] = msg.protocole
		}
		if m.currentPort <= m.totalPorts {
			cmd := scanPortCmd(m.host, m.currentPort)
			m.currentPort++
			return m, cmd
		}
		if m.scanned >= m.totalPorts {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	progress := fmt.Sprintf("Scanning... %d/%d\n\n", m.scanned, m.totalPorts)

	if len(m.openPorts) == 0 {
		return progress + "No open ports yet...\n'q' to quit"
	}

	var result strings.Builder
	result.WriteString("Open ports:\n")

	keys := make([]int, 0, len(m.openPorts))
	for k := range m.openPorts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Fprintf(&result, " - %d/tcp (%s)\n", k, m.openPorts[k])
	}

	return progress + result.String() + "\n'q' to quit"
}

func scanPortCmd(host string, port int) tea.Cmd {
	return func() tea.Msg {
		isOpen, protocole := scanner.ScanPort(host, port)
		return PortScannedMsg{port, protocole, isOpen}
	}
}
