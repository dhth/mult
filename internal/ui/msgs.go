package ui

type HideHelpMsg struct{}

type CmdRunChosenMsg struct {
	runNum int
}

type CmdRanMsg struct {
	iterationNum int
	output       string
	tookMS       int64
	err          error
}
