package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

var myStr1 = "This is the volumes tab"

type volumesComponent struct {
	commandOptions []string
	commandOutputs []string
	Cursor         int
}

func (m volumesComponent) Init() tea.Cmd {
	return nil
}

func (m volumesComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "up":
			m.Cursor--
			return m, nil
		case "down":
			m.Cursor++
			return m, nil
		}
	}

	return m, nil

}

func (m volumesComponent) View() string {
	myStr1 = ""
	for i, commandOption := range m.commandOptions {
		if m.Cursor == i {
			myStr1 += "[x]" + commandOption + "\n"
		} else {
			myStr1 += "[ ]" + commandOption + "\n"
		}
	}

	myString2 := m.commandOutputs[m.Cursor]
	return myString2
}

//type volumesInputModel struct {
//	commandOptions     []string
//	Cursor             *int
//	volumesOutputModel tea.Model
//}
//
//func (m volumesInputModel) Init() tea.Cmd {
//	return nil
//}
//
//func (m volumesInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	newMainModel, _ := m.volumesOutputModel.Update(msg)
//	m.volumesOutputModel = newMainModel
//	return m, nil
//}
//
//func (m volumesInputModel) View() string {
//	myStr1 = ""
//	for i, commandOption := range m.commandOptions {
//		if *m.Cursor == i {
//			myStr1 += "[x]" + commandOption + "\n"
//		} else {
//			myStr1 += "[ ]" + commandOption + "\n"
//		}
//	}
//	return ""
//}
//
//type volumesOutputModel struct {
//	commandOutputs []string
//	cursor         *int
//}
//
//func (m volumesOutputModel) Init() tea.Cmd {
//	return nil
//}
//
//func (m volumesOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch keypress := msg.String(); keypress {
//		case "up":
//			*m.cursor--
//			return m, nil
//		case "down":
//			*m.cursor++
//			return m, nil
//		}
//	}
//
//	return m, nil
//}
//
//func (m volumesOutputModel) View() string {
//	return "HELLO STRING 2"
//	//return m.commandOutputs[*m.cursor]
//}
