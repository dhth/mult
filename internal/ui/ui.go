package ui

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	d "github.com/dhth/mult/internal/domain"
)

var errFailedToConfigureDebugging = errors.New("failed to configure debugging")

func RenderUI(cmd []string, config d.Config) error {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			return fmt.Errorf("%w: %s", errFailedToConfigureDebugging, err.Error())
		}
		defer f.Close()
	}

	p := tea.NewProgram(InitialModel(cmd, config))
	_, err := p.Run()

	return err
}
