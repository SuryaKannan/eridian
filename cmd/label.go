package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

func initLabelModel() labelModel {
	return labelModel{
		screenName: ScreenName[Label],
	}
}

type labelModel struct {
	screenName string
}

func (m labelModel) Init() tea.Cmd {
	return nil
}

func (m labelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "esc":
			return m, returnToRoot
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m labelModel) View() tea.View {
	var s strings.Builder

	s.WriteString(selectedStyle.Render("(home/"+m.screenName+")") + "\n\n")

	s.WriteString(titleStyle.Render("\nPress ESC to return home.\n"))

	view := tea.NewView(s.String())
	view.AltScreen = true
	return view
}
