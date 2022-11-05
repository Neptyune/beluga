package tui

import tea "github.com/charmbracelet/bubbletea"

type dashboardModel struct {
}

func (m dashboardModel) Init() tea.Cmd {
	return nil
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m dashboardModel) View() string {
	return "This is the dashboard tab"
}
