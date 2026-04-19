// Package ui
package ui

import (
	"fmt"
	"sort"
	"strings"

	"port-scanner-visualiser/scanner"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	host        string
	startPort   int
	endPort     int
	concurrency int
	scanned     int
	currentPort int
	openPorts   map[int]string
	done        bool
}

type PortScannedMsg struct {
	port      int
	protocole string
	isOpen    bool
}

func InitialModel(host string, start int, end int, concurrency int) model {
	return model{
		host:        host,
		startPort:   start,
		endPort:     end,
		concurrency: concurrency,
		currentPort: start + concurrency,
		scanned:     0,
		openPorts:   make(map[int]string),
		done:        false,
	}
}

func (m model) totalPorts() int {
	return m.endPort - m.startPort + 1
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, m.concurrency)
	for i := range m.concurrency {
		cmds[i] = scanPortCmd(m.host, m.startPort+i)
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
		if m.currentPort <= m.endPort {
			cmd := scanPortCmd(m.host, m.currentPort)
			m.currentPort++
			return m, cmd
		}
		if m.scanned >= m.totalPorts() {
			m.done = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func header(host string, openCount int) string {
	target := progressStyle.Render("Target: " + host)
	open := openPortStyle.Render(fmt.Sprintf("Open: %d", openCount))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("Port Scanner"),
		lipgloss.JoinHorizontal(lipgloss.Left, target, "  |  ", open),
	)
}

func (m model) View() string {
	percent := float64(m.scanned) / float64(m.totalPorts())
	filled := int(percent * 40)

	barStyle := progressStyle
	if m.done {
		barStyle = doneStyle
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", 40-filled)
	status := "Scanning..."
	if m.done {
		status = "Done ✓"
	}
	progress := barStyle.Render(fmt.Sprintf("%s [%s] %d%%", status, bar, int(percent*100)))

	if len(m.openPorts) == 0 {
		box := boxStyle.Render(emptyStyle.Render("No open ports yet..."))
		return lipgloss.JoinVertical(lipgloss.Left,
			header(m.host, 0),
			progress,
			box,
			emptyStyle.Render("'q' to quit"),
		)
	}

	var b strings.Builder
	keys := make([]int, 0, len(m.openPorts))
	for k := range m.openPorts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Fprintf(&b, "→ %d/tcp (%s)\n", k, m.openPorts[k])
	}

	openPortsBox := boxStyle.Render(openPortStyle.Render(b.String()))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header(m.host, len(m.openPorts)),
		progress,
		openPortsBox,
		emptyStyle.Render("'q' to quit"),
	)
}

func scanPortCmd(host string, port int) tea.Cmd {
	return func() tea.Msg {
		isOpen, protocole := scanner.ScanPort(host, port)
		return PortScannedMsg{port, protocole, isOpen}
	}
}
