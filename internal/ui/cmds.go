package ui

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/dhth/mult/internal/executor"
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

func runCmd(runner *executor.Runner, cmd []string, iterationNum int) tea.Cmd {
	return func() tea.Msg {
		env := append(os.Environ(), fmt.Sprintf("MULT_RUN_NUM=%d", iterationNum+1))
		startTime := time.Now()
		out, err := runner.Run(cmd, env)
		endTime := time.Now()
		return CmdRanMsg{
			iterationNum: iterationNum,
			output:       string(out),
			tookMS:       endTime.Sub(startTime).Milliseconds(),
			err:          err,
		}
	}
}
