package tui

import (
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/guptarohit/asciigraph"
)

type graphModel struct {
	stopwatch stopwatch.Model
	cpu       []float64
	memory    []float64
}

type dashboardModel struct {
	graph graphModel
}

func (m dashboardModel) Init() tea.Cmd {
	return nil
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			updateGraph()
			return m, nil
		}
	}
	return m, nil
}

func (m dashboardModel) View() string {

}

func updateGraph() string {
	return lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).
		Render(asciigraph.Plot(data))
}
