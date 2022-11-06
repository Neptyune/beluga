package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var myStr = "This is the volumes tab"

type volumesInputModel struct {
	commandOptions     []string
	Cursor             *int
	volumesOutputModel tea.Model
}

func (m volumesInputModel) Init() tea.Cmd {
	return nil
}

func (m volumesInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newMainModel, _ := m.volumesOutputModel.Update(msg)
	m.volumesOutputModel = newMainModel
	return m, nil
}

func (m volumesInputModel) View() string {
	myStr = ""
	for i, commandOption := range m.commandOptions {
		if *m.Cursor == i {
			myStr += "[x]" + commandOption + "\n"
		} else {
			myStr += "[ ]" + commandOption + "\n"
		}
	}
	return myStr
}

type volumesOutputModel struct {
	commandOutputs []string
	cursor         *int
}

func (m volumesOutputModel) Init() tea.Cmd {
	return nil
}

func (m volumesOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "up":
			*m.cursor--
			return m, nil
		case "down":
			*m.cursor++
			return m, nil
		}
	}

	return m, nil
}

func (m volumesOutputModel) View() string {
	return "HELLO STRING 2"
}
