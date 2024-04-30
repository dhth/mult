package ui

import "fmt"

type runStatus uint

const (
	running runStatus = iota
	waiting
	scheduled
	finished
	abandoned
)

type command struct {
	IterationNum int
	Output       string
	RunStatus    runStatus
	TookMS       int64
	Err          error
}

func (c command) Title() string {
	return cmdRunNumStyle.Render(fmt.Sprintf("run #%d", c.IterationNum+1))
}

func (c command) Description() string {
	var runIndicator string
	switch c.RunStatus {
	case waiting:
		runIndicator = cmdWaitingStyle.Render("waiting")
	case scheduled:
		runIndicator = cmdScheduledStyle.Render("scheduled")
	case running:
		runIndicator = cmdRunningStyle.Render("running")
	case abandoned:
		runIndicator = cmdAbandonedStyle.Render("abandoned")
	default:
		if c.Err != nil {
			runIndicator = cmdErrorStyle.Render("error")
		} else {
			runIndicator = fmt.Sprintf("%s %s",
				cmdRanStyle.Render("finished"),
				cmdDurationStyle.Render(fmt.Sprintf("in %d ms",
					c.TookMS)))
		}
	}
	return runIndicator
}

func (c command) FilterValue() string {
	return fmt.Sprintf("%d", c.IterationNum)
}
