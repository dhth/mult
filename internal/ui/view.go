package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var runListWidth = 32

func (m Model) View() tea.View {
	var content string
	var footer string

	var statusBar string
	if m.msg.value != "" {
		switch m.msg.kind {
		case userMsgErr:
			statusBar = userMsgErrStyle.Render(m.msg.value)
		case userMsgInfo:
			statusBar = userMsgInfoStyle.Render(m.msg.value)
		}
	}

	switch m.activePane {
	case helpPane:
		var helpVP string
		if !m.helpVPReady {
			helpVP = "\n  Initializing..."
		} else {
			helpVP = helpVPStyle.Render(fmt.Sprintf("%s\n\n%s\n", helpVPTitleStyle.Render("Help"), m.helpVP.View()))
		}
		content = helpVP
	default:
		listView := runListStyle.Render(m.runList.View())
		outputView := lipgloss.JoinVertical(lipgloss.Left, "\n"+m.outputTitleStyle.Render("Output")+"\n\n"+m.outputVP.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top, listView, outputView)
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	var helpMsg string
	if m.showHelpIndicator {
		helpMsg = helpMsgStyle.Render("Press ? for help")
	}

	numRunsMsg := numRunsStyle.Render(fmt.Sprintf("%d/%d", m.numRunsFinished, m.config.NumRuns))

	var averageTimeMsg string
	if m.numSuccessfulRuns > 0 {
		averageTimeMsg = averageTimeMsgStyle.Render(fmt.Sprintf("average time: %d ms", m.averageMS))
	}

	var numErrorsMsg string
	if m.numErrors > 0 {
		numErrorsMsg = numErrorsStyle.Render(fmt.Sprintf("%d errors", m.numErrors))
	}

	var abandonedMsg string
	if m.abandoned {
		abandonedMsg = abandonedMsgStyle.Render("abandoned")
	}

	var followingMsg string
	if m.config.FollowResults {
		followingMsg = followingStyle.Render("[following]")
	}

	footerStr := fmt.Sprintf("%s%s%s%s%s%s%s",
		modeStyle.Render("mult"),
		helpMsg,
		numRunsMsg,
		averageTimeMsg,
		numErrorsMsg,
		abandonedMsg,
		followingMsg,
	)
	footer = footerStyle.Render(footerStr)

	v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	))
	v.AltScreen = true

	return v
}
