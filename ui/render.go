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
		currentPort: 1,
		scanned:     0,
		openPorts:   make(map[int]string),
	}
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 100)

	for i := 0; i < 100; i++ {
		cmds[i] = scanPortCmd(m.host, i+1)
		m.currentPort++
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
		m.currentPort++
		if m.currentPort <= m.totalPorts {
			return m, scanPortCmd(m.host, m.currentPort)
		}
		if m.scanned >= m.totalPorts {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	progress := fmt.Sprintf("Scanning...%d/%d", m.scanned, m.totalPorts)
	totalPorts := printResult(m.openPorts)

	if m.scanned > m.totalPorts {
		return progress
	}
	return totalPorts
}

func scanPortCmd(host string, port int) tea.Cmd {
	return func() tea.Msg {
		isOpen, protocole := scanner.ScanPort(host, port)
		return PortScannedMsg{port, protocole, isOpen}
	}
}

func printResult(portsMap map[int]string) [][]string {
	keys := make([]int, 0, len(portsMap))
	for k := range portsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := make([][]string, 0, len(keys))

	for _, k := range keys {
		result = append(result, []string{
			fmt.Sprintf("%d", k),
			portsMap[k],
		})
	}

	return result
}
