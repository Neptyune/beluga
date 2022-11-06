package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var myStr = "This is the volumes tab"

type volumesModel struct {
	commandOptions []string
	cursor         int
	//volumesOutputModel tea.Model
}

func (m volumesModel) Init() tea.Cmd {
	return nil
}

func (m volumesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// We want to quit and move across tabs
		switch keypress := msg.String(); keypress {
		case "up":
			m.cursor--
			return m, nil
		case "down":
			m.cursor++
			return m, nil
		case "enter":

		}
	}

	return m, nil
}

func (m volumesModel) View() string {
	myStr = ""
	//fmt.Println(m.cursor)
	for i, commandOption := range m.commandOptions {
		if m.cursor == i {
			myStr += "[x]" + commandOption + "\n"
		} else {
			myStr += "[]" + commandOption + "\n"
		}
	}
	return myStr
}

type volumesOutputModel struct {
	commandPrompts []interface{}
}

func (m volumesOutputModel) Init() tea.Cmd {
	return nil
}

func (m volumesOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// We want to quit and move across tabs
		switch keypress := msg.String(); keypress {

		}
	}

	return m, nil
}

func (m volumesOutputModel) View() string {
	//myStr = ""
	////fmt.Println(m.cursor)
	//for i, commandOption := range m.commandOptions {
	//	if m.cursor == i {
	//		myStr += "[x]" + commandOption + "\n"
	//	} else {
	//		myStr += "[ ]" + commandOption + "\n"
	//	}
	//}
	//return myStr
}
