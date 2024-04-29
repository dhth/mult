package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultBackgroundColor = "#282828"
	cmdWaitingColor        = "#fabd2f"
	cmdScheduledColor      = "#928374"
	cmdRunningColor        = "#83a598"
	cmdRanColor            = "#b8bb26"
	cmdErrorColor          = "#fb4934"
	cmdRunListColor        = "#b8bb26"
	cmdDurationColor       = "#928374"
	cmdRunColor            = "#928374"
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

	modeStyle = baseStyle.Copy().
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#b8bb26"))

	cmdIndicatorStyle = lipgloss.NewStyle().
				Bold(true)

	cmdWaitingStyle = cmdIndicatorStyle.Copy().
			Foreground(lipgloss.Color(cmdWaitingColor))

	cmdScheduledStyle = cmdIndicatorStyle.Copy().
				Foreground(lipgloss.Color(cmdScheduledColor))

	cmdRunningStyle = cmdIndicatorStyle.Copy().
			Foreground(lipgloss.Color(cmdRunningColor))

	cmdRanStyle = cmdIndicatorStyle.Copy().
			Foreground(lipgloss.Color(cmdRanColor))

	cmdErrorStyle = cmdIndicatorStyle.Copy().
			Foreground(lipgloss.Color(cmdErrorColor))

	helpMsgStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color("#83a598"))

	numRunsStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color(cmdScheduledColor))

	numErrorsStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true).
			Foreground(lipgloss.Color(numErrorsColor))

	inActivePaneHeaderStyle = baseStyle.Copy().
				Align(lipgloss.Left).
				Bold(true).
				Background(lipgloss.Color(inactivePaneColor))

	cmdRunNumStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(cmdRunColor))

	cmdDurationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(cmdDurationColor))
)
