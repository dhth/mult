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

type userMsgKind uint

const (
	userMsgInfo userMsgKind = iota
	userMsgErr
)

type userMsg struct {
	value         string
	kind          userMsgKind
	numFramesLeft uint
}

type Model struct {
	cmd               []string
	config            d.Config
	runList           list.Model
	lastRunIndex      int
	outputVP          viewport.Model
	outputVPReady     bool
	resultsCache      map[int]string
	msg               userMsg
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

func (m *Model) clearRunList() tea.Cmd {
	restart := func() tea.Msg {
		return CmdListClearedMsg{}
	}

	numRuns := len(m.runList.Items())
	if numRuns == 0 {
		return restart
	}

	stackItems := make([]list.Item, numRuns)

	for i := range numRuns {
		var rs runStatus
		if i == 0 || !m.config.Sequential {
			rs = running
		} else {
			rs = scheduled
		}
		stackItems[i] = command{
			IterationNum: i,
			RunStatus:    rs,
		}
	}

	m.resultsCache = make(map[int]string)
	m.lastRunIndex = -1
	m.firstFetch = true
	m.averageMS = 0
	m.totalMS = 0
	m.numRunsFinished = 0
	m.numSuccessfulRuns = 0
	m.numErrors = 0
	m.abandoned = false
	m.runList.Select(0)

	return tea.Sequence(m.runList.SetItems(stackItems), restart)
}

func (m Model) restartRuns() tea.Cmd {
	if m.config.Sequential {
		return runCmd(m.cmd, 0)
	}

	var cmds []tea.Cmd
	for i := 0; i < m.config.NumRuns; i++ {
		cmds = append(cmds, runCmd(m.cmd, i))
	}

	return tea.Batch(cmds...)
}
