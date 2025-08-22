package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func newCmdItemDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(lipgloss.Color(cmdRunListColor)).
		BorderStyle(lipgloss.OuterHalfBlockBorder()).
		BorderLeftForeground(lipgloss.Color(cmdRunListColor))
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		//revive:disable:unnecessary-stmt
		switch msgType := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msgType,
				list.DefaultKeyMap().CursorUp,
				list.DefaultKeyMap().CursorDown,
				list.DefaultKeyMap().GoToStart,
				list.DefaultKeyMap().GoToEnd,
				list.DefaultKeyMap().NextPage,
				list.DefaultKeyMap().PrevPage):
				runItem, ok := m.SelectedItem().(cmdRunItem)
				if ok {
					return chooseRunEntry(runItem.IterationNum)
				}
			}
		}
		//revive:enable:unnecessary-stmt
		return nil
	}

	return d
}
