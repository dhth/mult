package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultBackgroundColor = "#282828"
	cmdWaitingColor        = "#fabd2f"
	cmdScheduledColor      = "#928374"
	averageTimeColor       = "#fabd2f"
	cmdRunningColor        = "#83a598"
	cmdRanColor            = "#b8bb26"
	cmdErrorColor          = "#fb4934"
	cmdRunListColor        = "#b8bb26"
	cmdDurationColor       = "#928374"
	cmdRunColor            = "#928374"
	cmdAbandonedColor      = "#bdae93"
	cmdErrorDetailsColor   = "#928374"
	activePaneColor        = "#b8bb26"
	inactivePaneColor      = "#928374"
	modeColor              = "#b8bb26"
	helpMsgColor           = "#83a598"
	helpViewTitleColor     = "#83a598"
	helpHeaderColor        = "#83a598"
	helpSectionColor       = "#fabd2f"
	numErrorsColor         = "#fb4934"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color("#282828"))

	runListStyle = baseStyle.
			PaddingTop(1).
			PaddingRight(2).
			PaddingBottom(1).
			Width(runListWidth + 4)

	modeStyle = baseStyle.
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#b8bb26"))

	cmdIndicatorStyle = lipgloss.NewStyle().
				Bold(true)

	cmdWaitingStyle = cmdIndicatorStyle.
			Foreground(lipgloss.Color(cmdWaitingColor))

	cmdScheduledStyle = cmdIndicatorStyle.
				Foreground(lipgloss.Color(cmdScheduledColor))

	cmdRunningStyle = cmdIndicatorStyle.
			Foreground(lipgloss.Color(cmdRunningColor))

	cmdSuccessStyle = cmdIndicatorStyle.
			Foreground(lipgloss.Color(cmdRanColor))

	cmdErrorStyle = cmdIndicatorStyle.
			Foreground(lipgloss.Color(cmdErrorColor))

	cmdAbandonedStyle = cmdIndicatorStyle.
				Foreground(lipgloss.Color(cmdAbandonedColor))

	cmdErrorDetailsStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(cmdErrorDetailsColor))

	helpMsgStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color("#83a598"))

	numRunsStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color(cmdScheduledColor))

	averageTimeMsgStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Bold(true).
				Foreground(lipgloss.Color(averageTimeColor))

	numErrorsStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color(numErrorsColor))

	abandonedMsgStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Bold(true).
				Foreground(lipgloss.Color(cmdAbandonedColor))

	inActivePaneHeaderStyle = baseStyle.
				Align(lipgloss.Left).
				Bold(true).
				Background(lipgloss.Color(inactivePaneColor))

	cmdRunNumStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(cmdRunColor))

	cmdDurationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(cmdDurationColor))
)
