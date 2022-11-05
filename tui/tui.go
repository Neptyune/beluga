package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Entry to point to start the TUI
func StartTea() {
	tabs := []string{"CONTAINER", "IMAGES", "VOLUMES", "DASHBOARD"}

	m := mainModel{Tabs: tabs}
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("There was an error: %v\n", err)
		os.Exit(1)
	}
}

type sessionState int

const (
	container sessionState = iota
	images
	volumes
	dashboard
)

type mainModel struct {
	state      sessionState
	Tabs       []string  // Tabs to be shown
	TabContent tea.Model // What is inside the tabs
	activeTab  int       // Currently chose tab
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m mainModel) View() string {
	return "Hello world!"
}
