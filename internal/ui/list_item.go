package ui

import (
	"fmt"

	d "github.com/dhth/mult/internal/domain"
)

type cmdRunItem struct {
	d.CommandRun
}

func (c cmdRunItem) Title() string {
	return cmdRunNumStyle.Render(fmt.Sprintf("run #%d", c.IterationNum+1))
}

func (c cmdRunItem) Description() string {
	var runIndicator string
	switch c.RunStatus {
	case d.Waiting:
		runIndicator = cmdWaitingStyle.Render("waiting")
	case d.Scheduled:
		runIndicator = cmdScheduledStyle.Render("scheduled")
	case d.Running:
		runIndicator = cmdRunningStyle.Render("running")
	case d.Abandoned:
		runIndicator = cmdAbandonedStyle.Render("abandoned")
	default:
		var exitIndicator string
		if c.Err != nil {
			exitIndicator = cmdErrorStyle.Render("errored")
		} else {
			exitIndicator = cmdSuccessStyle.Render("finished")
		}
		runIndicator = fmt.Sprintf("%s %s",
			exitIndicator,
			cmdDurationStyle.Render(fmt.Sprintf("after %d ms",
				c.TookMS)))
	}
	return runIndicator
}

func (c cmdRunItem) FilterValue() string {
	return fmt.Sprintf("%d", c.IterationNum)
}

func toListItem(cmdRun d.CommandRun) cmdRunItem {
	return cmdRunItem{CommandRun: cmdRun}
}
