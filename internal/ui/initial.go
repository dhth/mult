package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	d "github.com/dhth/mult/internal/domain"
)

func InitialModel(cmd []string, config d.Config) Model {
	stackItems := make([]list.Item, config.NumRuns)

	for i := range config.NumRuns {
		var rs runStatus
		if i == 0 || !config.Sequential {
			rs = running
		} else {
			rs = scheduled
		}
		stackItems[i] = command{
			IterationNum: i,
			RunStatus:    rs,
		}
	}

	del := newCmdItemDelegate()

	outputTitleStyle := inActivePaneHeaderStyle.
		Background(lipgloss.Color(inactivePaneColor))

	m := Model{
		cmd:              cmd,
		msg:              userMsg{},
		config:           config,
		lastRunIndex:     -1,
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
