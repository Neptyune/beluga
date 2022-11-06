package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

//var myStr1 = "This is the volumes tab"

type componentInputModel struct {
	commandOptions     []string
	Cursor             *int
	volumesOutputModel tea.Model
}

func (m componentInputModel) Init() tea.Cmd {
	return nil
}

func (m componentInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newMainModel, _ := m.volumesOutputModel.Update(msg)
	m.volumesOutputModel = newMainModel
	return m, nil
}

func (m componentInputModel) View() string {
	myStr1 = ""
	for i, commandOption := range m.commandOptions {
		if *m.Cursor == i {
			myStr1 += "[x]" + commandOption + "\n"
		} else {
			myStr1 += "[ ]" + commandOption + "\n"
		}
	}
	return ""
}

type componentOutputModel struct {
	commandOutputs []string
	cursor         *int
}

func (m componentOutputModel) Init() tea.Cmd {
	return nil
}

func (m componentOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m componentOutputModel) View() string {
	return "HELLO STRING 2"
	//return m.commandOutputs[*m.cursor]
}
