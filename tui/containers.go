package tui

import tea "github.com/charmbracelet/bubbletea"

type containersModel struct {
}

func (m containersModel) Init() tea.Cmd {
	return nil
}

func (m containersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m containersModel) View() string {
	return "This is the container tab."
}
