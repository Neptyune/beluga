package tui

import tea "github.com/charmbracelet/bubbletea"

type imagesModel struct {
}

func (m imagesModel) Init() tea.Cmd {
	return nil
}

func (m imagesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m imagesModel) View() string {
	return "This is the images tab"
}
