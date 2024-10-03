package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/dhth/mult/internal/ui"
)

var (
	errInvalidNumRunsRequested = errors.New("invalid number of runs requested")
	errNoCommandProvided       = errors.New("no command provided")
)

var (
	numRuns       = flag.Int("n", 5, "number of times to run the command")
	sequential    = flag.Bool("s", false, "whether to invoke the command sequentially")
	delayMS       = flag.Int("delay-ms", 0, "time to sleep for between runs")
	stopOnFailure = flag.Bool("ff", false, "whether to stop after first failure")
	interactive   = flag.Bool("i", false, "accept flag values interactively (takes precendence over -n)")
)

func Execute() error {
	flag.Usage = func() {
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

	if nRuns <= 1 {
		return fmt.Errorf("%w: needs to greater than 1", errInvalidNumRunsRequested)
	}

	cmdToRun := flag.Args()
	if len(cmdToRun) == 0 {
		return errNoCommandProvided
	}
	return ui.RenderUI(cmdToRun, nRuns, *sequential, *delayMS, *stopOnFailure)
}
