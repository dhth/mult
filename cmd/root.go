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
	numRuns    = flag.Int("n", 5, "number of times to run the command")
	sequential = flag.Bool("s", false, "whether to invoke the command sequentially")
)

func Execute() {
	flag.Parse()

	cmdToRun := flag.Args()
	if len(cmdToRun) == 0 {
		die("Provide a command to run")
	}
	if *numRuns <= 1 {
		die("num-runs needs to be atleast 2")
	}
	ui.RenderUI(cmdToRun, *numRuns, *sequential)
}
