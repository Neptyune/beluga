package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/neptyune/beluga/utils"
)

var myStr = "This is the volumes tab"

type volumesModel struct {
}

func (m volumesModel) Init() tea.Cmd {
	return nil
}

func (m volumesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	myStr = utils.VolumePrune()
	return m, nil
}

func (m volumesModel) View() string {
	return myStr
}
