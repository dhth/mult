package ui

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
		case tea.KeyPressMsg:
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
