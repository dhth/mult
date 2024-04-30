package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhth/mult/internal/ui"
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var (
	numRuns       = flag.Int("n", 5, "number of times to run the command")
	sequential    = flag.Bool("s", false, "whether to invoke the command sequentially")
	delayMS       = flag.Int("delay-ms", 0, "time to sleep for between runs")
	stopOnFailure = flag.Bool("ff", false, "whether to stop after first failure")
	interactive   = flag.Bool("i", false, "accept flag values interactively (takes precendence over -n)")
)

func Execute() {
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
			die("provide a valid integer")
		}
	} else {
		nRuns = *numRuns
	}

	if nRuns <= 1 {
		die("number of runs needs to greater than 1")
	}

	cmdToRun := flag.Args()
	if len(cmdToRun) == 0 {
		die("Provide a command to run")
	}
	ui.RenderUI(cmdToRun, nRuns, *sequential, *delayMS, *stopOnFailure)
}
