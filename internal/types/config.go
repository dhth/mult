package types

type Config struct {
	NumRuns            int
	Sequential         bool
	DelayMS            int
	StopOnFirstFailure bool
	StopOnFirstSuccess bool
}
