package domain

type Config struct {
	NumRuns            int
	Sequential         bool
	DelayMS            int
	FollowResults      bool
	StopOnFirstFailure bool
	StopOnFirstSuccess bool
}
