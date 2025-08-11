package cmd

import (
	"errors"
	"fmt"

	d "github.com/dhth/mult/internal/domain"
	"github.com/dhth/mult/internal/ui"
	"github.com/spf13/cobra"
)

const maxNumRuns = 1000

var (
	errInvalidNumRunsRequested = errors.New("invalid number of runs requested")
	errInvalidFlagsProvided    = errors.New("invalid flags provided")
)

func Execute() error {
	rootCmd := NewRootCommand()

	return rootCmd.Execute()
}

func NewRootCommand() *cobra.Command {
	var (
		numRuns            int
		sequential         bool
		delayMS            int
		stopOnFirstFailure bool
		stopOnFirstSuccess bool
		followRuns         bool
		interactive        bool
	)

	rootCmd := &cobra.Command{
		Use:          "mult [flags] -- <command>",
		Short:        "Run a command multiple times and glance the outputs",
		Example:      `mult -s -n 10 -d 1000 -- curl -sif -m 5 'https://some.url/that?fails=sometimes'`,
		SilenceUsage: true,
		Args:         cobra.MinimumNArgs(1),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			if delayMS > 0 && !sequential {
				return fmt.Errorf("%w: delay can only be specified in sequential mode", errInvalidFlagsProvided)
			}

			if stopOnFirstFailure && !sequential {
				return fmt.Errorf("%w: can only stop on first failure in sequential mode", errInvalidFlagsProvided)
			}

			if stopOnFirstSuccess && !sequential {
				return fmt.Errorf("%w: can only stop on first success in sequential mode", errInvalidFlagsProvided)
			}

			if stopOnFirstSuccess && stopOnFirstFailure {
				return fmt.Errorf("%w: cannot stop on first failure and success at the same time; choose one", errInvalidFlagsProvided)
			}

			if followRuns && !sequential {
				return fmt.Errorf("%w: follow mode can only be used in sequential mode", errInvalidFlagsProvided)
			}

			return nil
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var nRuns int
			if interactive {
				fmt.Printf("number of runs?\n")
				_, err := fmt.Scanf("%d", &nRuns)
				if err != nil {
					return fmt.Errorf("%w: invalid integer provided", errInvalidNumRunsRequested)
				}
			} else {
				nRuns = numRuns
			}

			if nRuns <= 1 || nRuns > maxNumRuns {
				return fmt.Errorf("%w: needs to be between 2 and %d (both inclusive)", errInvalidNumRunsRequested, maxNumRuns)
			}

			config := d.Config{
				NumRuns:            nRuns,
				Sequential:         sequential,
				DelayMS:            delayMS,
				FollowResults:      followRuns,
				StopOnFirstFailure: stopOnFirstFailure,
				StopOnFirstSuccess: stopOnFirstSuccess,
			}

			fmt.Printf("cmd: %v\n", args)
			return ui.RenderUI(args, config)
		},
	}

	rootCmd.Flags().IntVarP(&numRuns, "num-runs", "n", 5, "number of times to run the command")
	rootCmd.Flags().BoolVarP(&sequential, "sequential", "s", false, "whether to invoke the command sequentially")
	rootCmd.Flags().IntVarP(&delayMS, "delay", "d", 0, "time (in ms) to sleep for between runs")
	rootCmd.Flags().BoolVarP(&stopOnFirstFailure, "stop-on-first-failure", "F", false, "whether to stop after first failure")
	rootCmd.Flags().BoolVarP(&stopOnFirstSuccess, "stop-on-first-success", "S", false, "whether to stop after first success")
	rootCmd.Flags().BoolVarP(&followRuns, "follow", "f", false, `start with "follow mode" ON (ie, automatically select the latest command run)`)
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "accept flag values interactively (takes precendence over -n)")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
