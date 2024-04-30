package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane uint

const (
	cmdRunListPane Pane = iota
	outputPane
)

type model struct {
	cmd              []string
	numRuns          int
	runList          list.Model
	outputVP         viewport.Model
	outputVPReady    bool
	resultsCache     map[int]string
	message          string
	runListStyle     lipgloss.Style
	outputTitleStyle lipgloss.Style
	terminalHeight   int
	terminalWidth    int
	showHelp         bool
	activePane       Pane
	lastPane         Pane
	firstFetch       bool
	sequential       bool
	delayMS          int
	numRunsFinished  int
	numErrors        int
	stopOnFirstError bool
	abandoned        bool
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, hideHelp(time.Minute*1))
	cmds = append(cmds, runCmd(m.cmd, 0))
	if m.sequential {
		return tea.Batch(cmds...)
	} else {
		for i := 1; i < m.numRuns; i++ {
			cmds = append(cmds, runCmd(m.cmd, i))
		}
		return tea.Batch(cmds...)
	}
}
