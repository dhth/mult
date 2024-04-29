package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func InitialModel(cmd []string, numRuns int, sequential bool, delayMS int) model {

	stackItems := make([]list.Item, 0)

	stackItems = append(stackItems, command{
		IterationNum: 0,
		RunStatus:    running,
	})

	for i := 1; i < numRuns; i++ {
		var rs runStatus
		if sequential {
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

	baseStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(defaultBackgroundColor))

	tableListStyle := baseStyle.Copy().
		PaddingTop(1).
		PaddingRight(2).
		PaddingLeft(1).
		PaddingBottom(1)

	outputTitleStyle := inActivePaneHeaderStyle.Copy().
		Background(lipgloss.Color(inactivePaneColor))

	m := model{
		cmd:              cmd,
		message:          "hello",
		numRuns:          numRuns,
		resultsCache:     make(map[int]string),
		runList:          list.New(stackItems, del, 0, 0),
		runListStyle:     tableListStyle,
		outputTitleStyle: outputTitleStyle,
		showHelp:         true,
		firstFetch:       true,
		sequential:       sequential,
		delayMS:          delayMS,
	}

	m.runList.Title = "Runs"
	m.runList.SetStatusBarItemName("run", "runs")
	m.runList.DisableQuitKeybindings()
	m.runList.SetShowHelp(false)
	m.runList.Styles.Title.Foreground(lipgloss.Color(defaultBackgroundColor))
	m.runList.Styles.Title.Background(lipgloss.Color(cmdRunListColor))
	m.runList.Styles.Title.Bold(true)

	return m
}
