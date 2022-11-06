package tui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Entry to point to start the TUI
func StartTea() {
	tabs := []string{"CONTAINER", "IMAGES", "VOLUMES", "DASHBOARD"}
	// Creation of all the View models?
	containerModelObject := containerModel{}
	imagesModelObject := imagesModel{}
	volumesModelObject := volumesModel{}
	dashboardModelObject := dashboardModel{}
	tabContent := []tea.Model{containerModelObject, imagesModelObject, volumesModelObject, dashboardModelObject}

	m := mainModel{Tabs: tabs, TabContent: tabContent}
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("There was an error: %v\n", err)
		os.Exit(1)
	}
}

type sessionState int

const (
	container sessionState = iota
	images
	volumes
	dashboard
)

type mainModel struct {
	state      sessionState
	Tabs       []string    // Tabs to be shown
	TabContent []tea.Model // What is inside the tabs
	activeTab  int         // Currently chose tab
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// We want to quit and move across tabs
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			m.state++ // changing state so that it know what tab is in focux
			return m, nil
		case "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			m.state--
			return m, nil
		}

		switch m.state {
		case container:
			m.TabContent[m.activeTab].Update(msg)
			break
		case images:
			m.TabContent[m.activeTab].Update(msg)
			break
		case volumes:
			m.TabContent[m.activeTab].Update(msg)
			break
		case dashboard:
			m.TabContent[m.activeTab].Update(msg)
			break
		}

	}

	return m, nil
}

func (m mainModel) View() string {
	doc := strings.Builder{}
	var renderedTabs []string

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "|"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.TabContent[m.activeTab].View()))
	return docStyle.Render(doc.String())
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
