package ui

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var errFailedToConfigureDebugging = errors.New("failed to configure debugging")

func RenderUI(cmd []string, numRuns int, sequential bool, delayMS int, stopOnFailure bool) error {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			return fmt.Errorf("%w: %s", errFailedToConfigureDebugging, err.Error())
		}
		defer f.Close()
	}

	p := tea.NewProgram(InitialModel(cmd, numRuns, sequential, delayMS, stopOnFailure), tea.WithAltScreen())
	_, err := p.Run()

	return err
}
