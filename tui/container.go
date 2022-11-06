package tui

import tea "github.com/charmbracelet/bubbletea"

type containerModel struct {
}

func (m containerModel) Init() tea.Cmd {
	return nil
}

func (m containerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m containerModel) View() string {
	return "This is the container tab."
}
