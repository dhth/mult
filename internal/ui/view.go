package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var content string
	var footer string

	var statusBar string
	if m.message != "" {
		statusBar = RightPadTrim(m.message, m.terminalWidth)
	}

	listView := m.runListStyle.Render(m.runList.View())
	outputView := lipgloss.JoinVertical(lipgloss.Left, "\n"+m.outputTitleStyle.Render("Output")+"\n\n"+m.outputVP.View())
	content = lipgloss.JoinHorizontal(lipgloss.Top, listView, outputView)

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("tab: switch focus; j/k/down/up: scroll output up/down")
	}

	footerStr := fmt.Sprintf("%s%s",
		modeStyle.Render("mult"),
		helpMsg,
	)
	footer = footerStyle.Render(footerStr)

	return lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	)
}
