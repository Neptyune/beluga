package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type dashboardModel struct {
	memPercent  float64
	memProgress progress.Model
}

func (m dashboardModel) Init() tea.Cmd {
	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	m.memProgress = prog
	return tickCmd()
}

type tickMsg time.Time

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, nil
		}
	case tickMsg:
		m.memPercent += 0.25
		if m.memPercent > 1.0 {
			m.memPercent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	}
	return m, nil
}

func (m dashboardModel) View() string {
	pad := strings.Repeat(" ", 2)
	return "\n Viewing me rn" +
		pad + m.memProgress.ViewAs(m.memPercent)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
