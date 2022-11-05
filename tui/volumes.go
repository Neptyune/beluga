package tui

import tea "github.com/charmbracelet/bubbletea"

type volumesModel struct {
}

func (m volumesModel) Init() tea.Cmd {
	return nil
}

func (m volumesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m volumesModel) View() string {
	return "This is the volumes tab"
}
