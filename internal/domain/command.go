package domain

type RunStatus uint

const (
	Running RunStatus = iota
	Waiting
	Scheduled
	Finished
	Abandoned
)

type CommandRun struct {
	IterationNum int
	Output       string
	RunStatus    RunStatus
	TookMS       int64
	Err          error
}
