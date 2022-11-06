package tui

import (
	"fmt"
	"os"
	"strings"

	commandExecuter "github.com/neptyune/beluga/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Entry to point to start the TUI
func StartTea() {
	tabs := []string{"CONTAINERS", "IMAGES", "VOLUMES", "DASHBOARD"}
	// Creation of all the View models?
	containersModelObject := containersModel{}
	imagesModelObject := imagesModel{}
	//volumesModelOutput := CreateVolumeOutputModel()
	//volumesModelInput := CreateVolumeInputModel()
	volumeModel := createVolumesComponent()

	dashboardModelObject := dashboardModel{}
	tabContent := []tea.Model{containersModelObject, imagesModelObject, volumeModel, dashboardModelObject}
	//tabContent := []component{createVolumesComponent(), createVolumesComponent(), createVolumesComponent(), createVolumesComponent()}
	m := mainModel{Tabs: tabs, TabContent: tabContent}
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("There was an error: %v\n", err)
		os.Exit(1)
	}
}

func createVolumesComponent() tea.Model {
	return volumesComponent{
		commandOptions: []string{"List", "Prune"},
		commandOutputs: []string{commandExecuter.VolumeList(), commandExecuter.VolumePrune()},
		Cursor:         0,
	}
}

//func CreateVolumeOutputModel() volumesComponent {
//	var cursorPos int = 0
//	return volumesOutputModel{
//		commandOutputs: []string{commandExecuter.VolumeList(), commandExecuter.VolumePrune()},
//		cursor:         &cursorPos,
//	}
//}
//
//func CreateVolumeInputModel() volumesInputModel {
//	volumesModelOutput := CreateVolumeOutputModel()
//	return volumesInputModel{
//		commandOptions:     []string{"List", "Prune"},
//		Cursor:             volumesModelOutput.cursor,
//		volumesOutputModel: volumesModelOutput,
//	}
//}

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

		newMainModel, _ := m.TabContent[m.activeTab].Update(msg)
		m.TabContent[m.activeTab] = newMainModel

	}

	return m, nil
}

func (m mainModel) View() string {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Tabs [LAYER: 1]
	{
		var renderedTabs []string

		for i, tab := range m.Tabs {
			if i == m.activeTab {
				renderedTabs = append(renderedTabs, activeTabStyle.Render(tab))
			} else {
				renderedTabs = append(renderedTabs, tabStyle.Render(tab))
			}
		}

		var renderedTabBlock = lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

		gap := strings.Repeat(" ", max(0, physicalWidth-lipgloss.Width(renderedTabBlock)-sum(layerSpacing[:1])*2))
		doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabBlock, gap))
	}

	// Tab Content [LAYER: 2]
	{
		doc.WriteString("\n" + tabContentStyle.Width(physicalWidth-sum(layerSpacing[:2])*2).Render(m.TabContent[m.activeTab].View()))
	}

	// Page Drawing (the outer border) [LAYER: 0]
	{
		footer := lipgloss.JoinHorizontal( //Hard coded and doesn't change
			lipgloss.Bottom,
			footerEdgeStyle.Margin(0, 0, 0, 1).Render("└"),
			lipgloss.PlaceHorizontal(
				physicalWidth-4,
				lipgloss.Center,
				lipgloss.JoinHorizontal(
					lipgloss.Bottom,
					footerEdgeStyle.Render("〈"),
					lipgloss.NewStyle().Foreground(themeColours[0]).Italic(true).Render(" beluga "),
					footerEdgeStyle.Render("〉"),
				),
				lipgloss.WithWhitespaceChars("─"),
				lipgloss.WithWhitespaceForeground(themeColours[3]),
			),
			footerEdgeStyle.Margin(0, 1, 0, 0).Render("┘"),
		)

		page := pageStyle.Render(doc.String())
		doc.Reset()
		doc.WriteString(lipgloss.JoinVertical(lipgloss.Left, page, footerMarginAdder.Render(footer)))
	}
	// doc.WriteString("\n\n" + fmt.Sprintln(layerSpacing[0], layerSpacing[1], sum(layerSpacing[:2]), layerSpacing[:2], "\n", physicalWidth, physicalWidth-sum(layerSpacing[:1])*2))
	return doc.String()
}

// STYLING //

// A list to keep track of each of our horisontal layers spacing (padding/margins)
// -> Each element is a sum of the previous with the new spacing for that layer added
// [LAYER: x] references are given next to items that are on that specified layer (indexed from 0)
var layerSpacing = []int{
	(pageStyle.GetHorizontalFrameSize()) / 2,
	(tabContentStyle.GetHorizontalFrameSize()) / 2,
}

// Our theme colours light to dark
var themeColours = map[int]lipgloss.Color{
	0: lipgloss.Color("#B8D1EB"),
	1: lipgloss.Color("#89B9D9"),
	2: lipgloss.Color("#68A4CA"),
	3: lipgloss.Color("#4489B2"),
	4: lipgloss.Color("#255688"),
	5: lipgloss.Color("#183C69"),
}

var (
	// Page (the border of the app) [LAYER: 0]
	pageStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, true, false).
			BorderForeground(themeColours[4]).
			Margin(1, 1, 0).
			Padding(0, 1, 1)

	// Footer (bottom of border) [LAYER: 0]
	footerEdgeStyle = lipgloss.NewStyle().
			Foreground(themeColours[4])

	footerMarginAdder = lipgloss.NewStyle().
				Margin(0, 0, 1)

	// Tabs [LAYER: 1]
	tabBaseStyle = lipgloss.NewStyle().
			Margin(1, 2, 0).
			Padding(0, 1, 0).
			Foreground(themeColours[0])

	tabStyle = tabBaseStyle.Copy().
			Faint(true)

	activeTabStyle = tabBaseStyle.Copy().
			Border(lipgloss.NormalBorder(), false, false, true).
			BorderForeground(themeColours[1])

	// Tab Content [LAYER: 2]
	tabContentStyle = lipgloss.NewStyle().
			Margin(0, 4).
			Padding(2, 0).
			Align(lipgloss.Center).
			Border(lipgloss.NormalBorder())
)

// END OF STYLING //

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

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
