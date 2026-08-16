package ui

import (
	"context"
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	d "github.com/dhth/mult/internal/domain"
	"github.com/dhth/mult/internal/executor"
)

var errFailedToConfigureDebugging = errors.New("failed to configure debugging")

func RenderUI(ctx context.Context, cmd []string, config d.Config, runner *executor.Runner) error {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			return fmt.Errorf("%w: %s", errFailedToConfigureDebugging, err.Error())
		}
		defer f.Close()
	}

	p := tea.NewProgram(InitialModel(cmd, config, runner), tea.WithContext(ctx))
	_, err := p.Run()

	return err
}
