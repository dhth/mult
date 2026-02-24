package ui

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "charm.land/bubbletea/v2"
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

func runAfterDelay(interval time.Duration, iterationNum int) tea.Cmd {
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return DelayTimeElapsedMsg{iterationNum}
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

		c.Env = append(os.Environ(), fmt.Sprintf("MULT_RUN_NUM=%d", iterationNum+1))
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
