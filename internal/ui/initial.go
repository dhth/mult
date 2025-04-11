package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/mult/internal/types"
)

func InitialModel(cmd []string, config types.Config) Model {
	stackItems := make([]list.Item, 0)

	stackItems = append(stackItems, command{
		IterationNum: 0,
		RunStatus:    running,
	})

	for i := 1; i < config.NumRuns; i++ {
		var rs runStatus
		if config.Sequential {
			rs = scheduled
		} else {
			rs = running
		}
		stackItems = append(stackItems, command{
			IterationNum: i,
			RunStatus:    rs,
		})
	}

	del := newCmdItemDelegate()

	outputTitleStyle := inActivePaneHeaderStyle.
		Background(lipgloss.Color(inactivePaneColor))

	m := Model{
		cmd:              cmd,
		message:          "hello",
		config:           config,
		resultsCache:     make(map[int]string),
		runList:          list.New(stackItems, del, runListWidth, 0),
		outputTitleStyle: outputTitleStyle,
		showHelp:         true,
		firstFetch:       true,
	}

	m.runList.Title = "Runs"
	m.runList.SetStatusBarItemName("run", "runs")
	m.runList.DisableQuitKeybindings()
	m.runList.SetShowHelp(false)
	m.runList.SetFilteringEnabled(false)
	m.runList.Styles.Title = m.runList.Styles.Title.Foreground(lipgloss.Color(defaultBackgroundColor)).
		Background(lipgloss.Color(cmdRunListColor)).
		Bold(true)

	return m
}
