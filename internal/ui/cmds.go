package ui

import (
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func chooseRunEntry(runNum int) tea.Cmd {
	return func() tea.Msg {
		return CmdRunChosenMsg{runNum}
	}
}

func hideHelp(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return HideHelpMsg{}
	})
}

func runCmd(cmd []string, iterationNum int) tea.Cmd {
	return func() tea.Msg {
		var c *exec.Cmd

		if len(cmd) == 1 {
			c = exec.Command(cmd[0])
		} else {
			c = exec.Command(cmd[0], cmd[1:]...)

		}
		startTime := time.Now()
		out, err := c.CombinedOutput()
		endTime := time.Now()
		return CmdRanMsg{
			iterationNum: iterationNum,
			output:       string(out),
			tookMS:       endTime.Sub(startTime).Milliseconds(),
			err:          err,
		}
	}
}
