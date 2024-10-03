package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.message = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			switch m.activePane {
			case outputPane:
				m.activePane = cmdRunListPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(inactivePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(activePaneColor))
			default:
				return m, tea.Quit
			}
		case "tab", "shift+tab":
			if m.activePane == cmdRunListPane {
				m.activePane = outputPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(activePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(inactivePaneColor))
			} else if m.activePane == outputPane {
				m.activePane = cmdRunListPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(inactivePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(activePaneColor))
			}
		}
	case HideHelpMsg:
		m.showHelp = false
	case tea.WindowSizeMsg:
		w1, h1 := m.runListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.runList.SetHeight(msg.Height - h1 - 2)
		m.runList.SetWidth(int(float64(msg.Width)*0.25) - w1 - 2)
		m.runListStyle = m.runListStyle.Width(int(float64(msg.Width)*0.25) - w1)

		if !m.outputVPReady {
			m.outputVP = viewport.New(msg.Width-m.runListStyle.GetWidth()-2, msg.Height-8)
			m.outputVP.HighPerformanceRendering = false
			m.outputVPReady = true
		} else {
			m.outputVP.Width = msg.Width - m.runListStyle.GetWidth() - 2
			m.outputVP.Height = msg.Height - 8
		}

	case CmdRanMsg:
		m.numRunsFinished++
		i := msg.iterationNum
		run, ok := m.runList.Items()[i].(command)
		if ok {
			run.Output = msg.output
			run.Err = msg.err
			run.RunStatus = finished
			run.TookMS = msg.tookMS
			cmds = append(cmds, m.runList.SetItem(i, run))

			if msg.err != nil {
				errDetails := cmdErrorDetailsStyle.Render(fmt.Sprintf("---\n%s", msg.err.Error()))
				m.resultsCache[i] = fmt.Sprintf("%s\n%s", run.Output, errDetails)
				m.numErrors++
			} else {
				m.resultsCache[i] = run.Output
			}

			if m.firstFetch {
				selected, ok := m.runList.SelectedItem().(command)
				if ok {
					resultFromCache, ok := m.resultsCache[selected.IterationNum]
					if ok {
						m.outputVP.SetContent(resultFromCache)
						m.firstFetch = false
					}
				}
			}
		}

		if i < m.numRuns-1 && m.sequential {
			if m.stopOnFirstError && msg.err != nil {
				for j := i + 1; j < m.numRuns; j++ {
					nextRun, ok := m.runList.Items()[i+1].(command)
					if ok {
						nextRun.RunStatus = abandoned
						cmds = append(cmds, m.runList.SetItem(j, nextRun))
					}
				}
				m.abandoned = true
			} else {
				if m.delayMS == 0 {
					nextRun, ok := m.runList.Items()[i+1].(command)
					if ok {
						nextRun.RunStatus = running
						cmds = append(cmds, m.runList.SetItem(i+1, nextRun))
						cmds = append(cmds, runCmd(m.cmd, i+1))
					}
				} else {
					nextRun, ok := m.runList.Items()[i+1].(command)
					if ok {
						nextRun.RunStatus = waiting
						cmds = append(cmds, m.runList.SetItem(i+1, nextRun))
						cmds = append(cmds, runAfterDelay(time.Millisecond*time.Duration(m.delayMS), i+1))
					}
				}
			}
		}
	case DelayTimeElapsedMsg:
		run, ok := m.runList.Items()[msg.iterationNum].(command)
		if ok {
			run.RunStatus = running
			cmds = append(cmds, m.runList.SetItem(msg.iterationNum, run))
			cmds = append(cmds, runCmd(m.cmd, msg.iterationNum))
		}

	case CmdRunChosenMsg:
		resultFromCache, ok := m.resultsCache[msg.runNum]
		if ok {
			m.outputVP.SetContent(resultFromCache)
		} else {
			m.outputVP.SetContent("")
		}
	}

	switch m.activePane {
	case cmdRunListPane:
		m.runList, cmd = m.runList.Update(msg)
		cmds = append(cmds, cmd)
	case outputPane:
		m.outputVP, cmd = m.outputVP.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
