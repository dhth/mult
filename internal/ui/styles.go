package ui

import (
	"charm.land/lipgloss/v2"
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
	followingColor         = "#d3869b"
	cmdErrorDetailsColor   = "#928374"
	activePaneColor        = "#b8bb26"
	inactivePaneColor      = "#928374"
	numErrorsColor         = "#fb4934"
	modeColor              = "#b8bb26"
	helpColor              = "#83a598"
	userMsgInfoColor       = "#8ec07c"
	userMsgErrColor        = "#fb4934"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(defaultBackgroundColor))

	runListStyle = lipgloss.NewStyle().
			PaddingTop(1).
			PaddingRight(2).
			PaddingBottom(1).
			Width(runListWidth + 4)

	modeStyle = baseStyle.
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color(modeColor))

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

	statusBarElementStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Bold(true)

	helpMsgStyle = statusBarElementStyle.
			Foreground(lipgloss.Color(helpColor))

	numRunsStyle = statusBarElementStyle.
			Foreground(lipgloss.Color(cmdScheduledColor))

	averageTimeMsgStyle = statusBarElementStyle.
				Foreground(lipgloss.Color(averageTimeColor))

	numErrorsStyle = statusBarElementStyle.
			Foreground(lipgloss.Color(numErrorsColor))

	abandonedMsgStyle = statusBarElementStyle.
				Foreground(lipgloss.Color(cmdAbandonedColor))

	followingStyle = statusBarElementStyle.
			Foreground(lipgloss.Color(followingColor))

	inActivePaneHeaderStyle = baseStyle.
				Align(lipgloss.Left).
				Bold(true).
				Background(lipgloss.Color(inactivePaneColor))

	cmdRunNumStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(cmdRunColor))

	cmdDurationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(cmdDurationColor))

	userMsgInfoStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(userMsgInfoColor))

	userMsgErrStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(userMsgErrColor))

	helpVPStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingTop(1).
			PaddingRight(2).
			PaddingBottom(1)

	helpVPTitleStyle = baseStyle.
				Bold(true).
				Background(lipgloss.Color(helpColor))
)
