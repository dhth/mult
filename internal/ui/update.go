package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.message = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "shift+tab":
			if m.activePane == cmdRunListPane {
				m.activePane = outputPane
				m.outputTitleStyle.Background(lipgloss.Color(activePaneColor))
				m.runList.Styles.Title.Background(lipgloss.Color(inactivePaneColor))
			} else if m.activePane == outputPane {
				m.activePane = cmdRunListPane
				m.outputTitleStyle.Background(lipgloss.Color(inactivePaneColor))
				m.runList.Styles.Title.Background(lipgloss.Color(activePaneColor))
			}
		}
	case HideHelpMsg:
		m.showHelp = false
	case tea.WindowSizeMsg:
		w1, h1 := m.runListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.runList.SetHeight(msg.Height - h1 - 2)
		m.runList.SetWidth(int(float64(msg.Width-w1) * 0.3))
		m.runListStyle.Width(int(float64(msg.Width-w1) * 0.3))

		if !m.outputVPReady {
			m.outputVP = viewport.New(int(float64(msg.Width-w1)*0.6), msg.Height-8)
			m.outputVP.HighPerformanceRendering = false
			m.outputVPReady = true
		} else {
			m.outputVP.Width = int(float64(msg.Width-w1) * 0.6)
			m.outputVP.Height = msg.Height - 8
		}

	case CmdRanMsg:
		i := msg.iterationNum
		run, ok := m.runList.Items()[i].(command)
		if ok {
			run.Output = msg.output
			run.Err = msg.err
			run.RunStatus = finished
			run.TookMS = msg.tookMS
			cmds = append(cmds, m.runList.SetItem(i, run))

			if msg.err != nil {
				m.resultsCache[i] = msg.err.Error()
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

		if m.sequential {
			if i < m.numRuns-1 {

				nextRun, ok := m.runList.Items()[i+1].(command)
				if ok {
					nextRun.RunStatus = running
					cmds = append(cmds, m.runList.SetItem(i+1, nextRun))
				}
				cmds = append(cmds, runCmd(m.cmd, i+1))
			}
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
