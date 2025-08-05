package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/dhth/mult/internal/types"
	"github.com/dhth/mult/internal/ui"
)

const maxNumRuns = 1000

var (
	errInvalidNumRunsRequested = errors.New("invalid number of runs requested")
	errNoCommandProvided       = errors.New("no command provided")
	errInvalidFlagsProvided    = errors.New("invalid flags provided")
)

var (
	numRuns            = flag.Int("n", 5, "number of times to run the command")
	sequential         = flag.Bool("s", false, "whether to invoke the command sequentially")
	delayMS            = flag.Int("d", 0, "time (in ms) to sleep for between runs")
	stopOnFirstFailure = flag.Bool("F", false, "whether to stop after first failure")
	stopOnFirstSuccess = flag.Bool("S", false, "whether to stop after first success")
	interactive        = flag.Bool("i", false, "accept flag values interactively (takes precendence over -n)")
)

func Execute() error {
	flag.Usage = func() {
		helpText := `Run a command multiple times and glance the outputs.

Usage: mult [flags]
`
		fmt.Fprintf(os.Stderr, "%s\nFlags:\n", helpText)
		flag.PrintDefaults()
	}
	flag.Parse()

	var nRuns int
	if *interactive {
		fmt.Printf("number of runs?\n")
		_, err := fmt.Scanf("%d", &nRuns)
		if err != nil {
			return fmt.Errorf("%w: invalid integer provided", errInvalidNumRunsRequested)
		}
	} else {
		nRuns = *numRuns
	}

	if nRuns <= 1 || nRuns > maxNumRuns {
		return fmt.Errorf("%w: needs to be between 2 and %d (both inclusive)", errInvalidNumRunsRequested, maxNumRuns)
	}

	if *delayMS > 0 && !*sequential {
		return fmt.Errorf("%w: -d can only be used in sequential mode", errInvalidFlagsProvided)
	}

	if *stopOnFirstFailure && !*sequential {
		return fmt.Errorf("%w: -F can only be used in sequential mode", errInvalidFlagsProvided)
	}

	if *stopOnFirstSuccess && !*sequential {
		return fmt.Errorf("%w: -S can only be used in sequential mode", errInvalidFlagsProvided)
	}

	if *stopOnFirstSuccess && *stopOnFirstFailure {
		return fmt.Errorf("%w: -F and -S cannot be used at the same time", errInvalidFlagsProvided)
	}

	cmdToRun := flag.Args()
	if len(cmdToRun) == 0 {
		return errNoCommandProvided
	}

	config := types.Config{
		NumRuns:            nRuns,
		Sequential:         *sequential,
		DelayMS:            *delayMS,
		StopOnFirstFailure: *stopOnFirstFailure,
		StopOnFirstSuccess: *stopOnFirstSuccess,
	}
	return ui.RenderUI(cmdToRun, config)
}
