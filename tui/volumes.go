package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/neptyune/beluga/customTable"
)

var myStr1 = "This is the volumes tab"

type volumesComponent struct {
	commandOptions []string
	commandOutputs []string
	Cursor         int
	customTable    customTable.Model
}

func (m volumesComponent) Init() tea.Cmd {
	return nil
}

func (m volumesComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil

}

func (m volumesComponent) View() string {
	// myStr1 = ""
	// for i, commandOption := range m.commandOptions {
	// 	if m.Cursor == i {
	// 		myStr1 += "[x]" + commandOption + "\n"
	// 	} else {
	// 		myStr1 += "[ ]" + commandOption + "\n"
	// 	}
	// }

	// myString2 := m.commandOutputs[m.Cursor]
	// return myString2

	columns := []customTable.Column{
		{Title: "Id", Width: 12},
		{Title: "Image", Width: 16},
		{Title: "Created", Width: 14},
		{Title: "Status", Width: 18},
		{Title: "Ports", Width: 18},
		{Title: "Names", Width: 16},
	}

	rows := []customTable.Row{
		{"64fd6eb0e26d", "docker101tutorial", "10 minutes ago", parseStatusString("Up 5 minutes"), "0.0.0.0:80->80/tcp", "docker-tutorial"},
		{"64fd6eb0e26d", "docker101tutorial", "10 minutes ago", parseStatusString("Up 5 minutes"), "0.0.0.0:80->80/tcp", "docker-tutorial"},
		{"7952bfa6b435", "alpine/git", "12 minutes ago", parseStatusString("Exited (0) 12 minutes ago"), "", "repo"},
		{"7952bfa6b435", "alpine/git", "12 minutes ago", parseStatusString("Exited (0) 12 minutes ago"), "", "repo"},
		{"7952bfa6b435", "alpine/git", "12 minutes ago", parseStatusString("Exited (0) 12 minutes ago"), "", "repo"},
	}

	t := customTable.New(
		customTable.WithColumns(columns),
		customTable.WithRows(rows),
		customTable.WithFocused(true),
		customTable.WithHeight(7),
	)

	s := customTable.DefaultStyles()
	s.Header = s.Header.
		Foreground(themeColours[1])
	// BorderStyle(lipgloss.NormalBorder()).
	// BorderForeground(lipgloss.Color("240")).
	// BorderBottom(true).
	// Bold(false)
	s.Selected = s.Selected.
		Foreground(themeColours[0])
		// Foreground(themeColours[4])
		// UnsetFaint()
		// Background(lipgloss.Color("57"))
	s.Cell = s.Cell.
		Foreground(themeColours[0]).
		Faint(true)

	t.SetStyles(s)

	// tableContent := m.customTable.View()

	// shutterTableStyle := lipgloss.NewStyle().
	// 	Foreground(themeColours[1])

	// table := lipgloss.JoinVertical(
	// 	lipgloss.Left,
	// 	shutterTableStyle.Render("╭"+strings.Repeat("─", 100-2)+"╮"),
	// 	tableContent,
	// 	shutterTableStyle.Render("╰"+strings.Repeat("─", 100-2)+"╯"),
	// )

	// return table + "\n"
	return `╭────────────────────────────────────────────────────────────────────────────────────────────────────────╮
	Id            Image           Created            Status              Ports              Names       
64fd6eb0e26d  docker101tutori…  10 minutes ago  Up 5 minutes        0.0.0.0:80->80/tcp  docker-tutorial  
64fd6eb0e26d  docker101tutori…  10 minutes ago  Up 5 minutes        0.0.0.0:80->80/tcp  docker-tutorial  
7952bfa6b435  alpine/git        12 minutes ago  Exited 12 minutes…                      repo             
7952bfa6b435  alpine/git        12 minutes ago  Exited 12 minutes…                      repo             
7952bfa6b435  alpine/git        12 minutes ago  Exited 12 minutes…                      repo             
																										
																										
╰────────────────────────────────────────────────────────────────────────────────────────────────────────╯`

}

func parseStatusString(status string) string {
	return strings.Replace(status, " (0) ", " ", 1)
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
