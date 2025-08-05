package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	d "github.com/dhth/mult/internal/domain"
)

type Pane uint

const (
	cmdRunListPane Pane = iota
	outputPane
)

type Model struct {
	cmd               []string
	config            d.Config
	runList           list.Model
	outputVP          viewport.Model
	outputVPReady     bool
	resultsCache      map[int]string
	message           string
	runListStyle      lipgloss.Style
	outputTitleStyle  lipgloss.Style
	terminalHeight    int
	terminalWidth     int
	showHelp          bool
	activePane        Pane
	firstFetch        bool
	averageMS         int64
	totalMS           int64
	numRunsFinished   int
	numSuccessfulRuns int
	numErrors         int
	abandoned         bool
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, hideHelp(time.Second*30))
	cmds = append(cmds, runCmd(m.cmd, 0))

	if m.config.Sequential {
		return tea.Batch(cmds...)
	}

	for i := 1; i < m.config.NumRuns; i++ {
		cmds = append(cmds, runCmd(m.cmd, i))
	}

	return tea.Batch(cmds...)
}
