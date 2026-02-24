package ui

import (
	_ "embed"
	"fmt"
	"time"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	d "github.com/dhth/mult/internal/domain"
)

//go:embed assets/help.txt
var helpText string

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	if m.msg.numFramesLeft == 0 {
		m.msg.value = ""
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			switch m.activePane {
			case outputPane:
				m.activePane = cmdRunListPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(inactivePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(activePaneColor))
			case helpPane:
				m.activePane = m.lastPane
			default:
				return m, tea.Quit
			}
		case "tab", "shift+tab":
			switch m.activePane {
			case cmdRunListPane:
				m.activePane = outputPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(activePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(inactivePaneColor))
			case outputPane:
				m.activePane = cmdRunListPane
				m.outputTitleStyle = m.outputTitleStyle.Background(lipgloss.Color(inactivePaneColor))
				m.runList.Styles.Title = m.runList.Styles.Title.Background(lipgloss.Color(activePaneColor))
			}
		case "ctrl+f":
			if m.activePane != cmdRunListPane && m.activePane != outputPane {
				break
			}
			if !m.config.Sequential {
				break
			}

			m.config.FollowResults = !m.config.FollowResults
		case "ctrl+r":
			if m.activePane != cmdRunListPane && m.activePane != outputPane {
				break
			}

			if m.numRunsFinished < m.config.NumRuns {
				m.msg = userMsg{"cannot restart while commands are being run; wait for them to finish", userMsgErr, 4}
			} else {
				cmds = append(cmds, m.clearRunList())
			}
		case "h", "left", "l", "right":
			if m.activePane != outputPane {
				break
			}

			selectedIndex := m.runList.Index()
			switch msg.String() {
			case "h", "left":
				if selectedIndex > 0 {
					m.runList.Select(selectedIndex - 1)
				}
			case "l", "right":
				if selectedIndex < len(m.runList.Items())-1 {
					m.runList.Select(selectedIndex + 1)
				}
			}

			if m.config.FollowResults {
				m.config.FollowResults = false
			}
		case "?":
			if m.activePane == helpPane {
				m.activePane = m.lastPane
				break
			}

			m.lastPane = m.activePane
			m.helpVP.GotoTop()
			m.showHelpIndicator = false
			m.activePane = helpPane
		}
	case HideHelpMsg:
		m.showHelpIndicator = false
	case CmdListClearedMsg:
		cmds = append(cmds, m.restartRuns())
	case tea.WindowSizeMsg:
		_, h1 := m.runListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.runList.SetHeight(msg.Height - h1 - 4)

		if !m.outputVPReady {
			m.outputVP = viewport.New(viewport.WithWidth(msg.Width-m.runListStyle.GetWidth()-2), viewport.WithHeight(msg.Height-8))
			m.outputVPReady = true
		} else {
			m.outputVP.SetWidth(msg.Width - m.runListStyle.GetWidth() - 2)
			m.outputVP.SetHeight(msg.Height - 8)
		}

		if !m.helpVPReady {
			m.helpVP = viewport.New(viewport.WithWidth(msg.Width-1), viewport.WithHeight(msg.Height-7))
			m.helpVP.SetContent(helpText)
			m.helpVPReady = true
		} else {
			m.helpVP.SetWidth(msg.Width - 1)
			m.helpVP.SetHeight(msg.Height - 7)
		}

	case CmdRanMsg:
		m.numRunsFinished++
		i := msg.iterationNum
		runItem, ok := m.runList.Items()[i].(cmdRunItem)
		if !ok {
			break
		}

		runItem.Output = msg.output
		runItem.Err = msg.err
		runItem.RunStatus = d.Finished
		runItem.TookMS = msg.tookMS
		cmds = append(cmds, m.runList.SetItem(i, runItem))
		if m.config.FollowResults && m.config.Sequential {
			m.runList.Select(i)
		}

		if msg.err != nil {
			errDetails := cmdErrorDetailsStyle.Render(fmt.Sprintf("---\n%s", msg.err.Error()))
			m.resultsCache[i] = fmt.Sprintf("%s\n%s", runItem.Output, errDetails)
			m.numErrors++
		} else {
			m.numSuccessfulRuns++
			m.resultsCache[i] = runItem.Output
			m.totalMS += runItem.TookMS
			m.averageMS = m.totalMS / int64(m.numSuccessfulRuns)
		}

		if m.firstFetch {
			selected, ok := m.runList.SelectedItem().(cmdRunItem)
			if ok {
				resultFromCache, ok := m.resultsCache[selected.IterationNum]
				if ok {
					m.outputVP.SetContent(resultFromCache)
					m.firstFetch = false
				}
			}
		}

		if i < m.config.NumRuns-1 && m.config.Sequential {
			if (m.config.StopOnFirstFailure && msg.err != nil) || (m.config.StopOnFirstSuccess && msg.err == nil) {
				for j := i + 1; j < m.config.NumRuns; j++ {
					nextRunItem, ok := m.runList.Items()[j].(cmdRunItem)
					if ok {
						nextRunItem.RunStatus = d.Abandoned
						cmds = append(cmds, m.runList.SetItem(j, nextRunItem))
					}
				}
				m.abandoned = true
			} else {
				if m.config.DelayMS == 0 {
					nextRunItem, ok := m.runList.Items()[i+1].(cmdRunItem)
					if ok {
						nextRunItem.RunStatus = d.Running
						cmds = append(cmds, m.runList.SetItem(i+1, nextRunItem))
						cmds = append(cmds, runCmd(m.cmd, i+1))
					}
				} else {
					nextRunItem, ok := m.runList.Items()[i+1].(cmdRunItem)
					if ok {
						nextRunItem.RunStatus = d.Waiting
						cmds = append(cmds, m.runList.SetItem(i+1, nextRunItem))
						cmds = append(cmds, runAfterDelay(time.Millisecond*time.Duration(m.config.DelayMS), i+1))
					}
				}
			}
		}
	case DelayTimeElapsedMsg:
		runItem, ok := m.runList.Items()[msg.iterationNum].(cmdRunItem)
		if !ok {
			break
		}

		runItem.RunStatus = d.Running
		cmds = append(cmds, m.runList.SetItem(msg.iterationNum, runItem))
		cmds = append(cmds, runCmd(m.cmd, msg.iterationNum))

	case CmdRunChosenMsg:
		if m.config.FollowResults {
			m.config.FollowResults = false
		}
	}

	switch m.activePane {
	case cmdRunListPane:
		m.runList, cmd = m.runList.Update(msg)
		cmds = append(cmds, cmd)
	case outputPane:
		m.outputVP, cmd = m.outputVP.Update(msg)
		cmds = append(cmds, cmd)
	case helpPane:
		m.helpVP, cmd = m.helpVP.Update(msg)
		cmds = append(cmds, cmd)
	}

	runIndex := m.runList.Index()
	if runIndex != m.lastRunIndex {
		m.lastRunIndex = runIndex
		resultFromCache, ok := m.resultsCache[runIndex]
		if ok {
			m.outputVP.SetContent(resultFromCache)
		} else {
			m.outputVP.SetContent("")
		}
	}

	if m.msg.numFramesLeft > 0 {
		m.msg.numFramesLeft--
	}

	return m, tea.Batch(cmds...)
}
