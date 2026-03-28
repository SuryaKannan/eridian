package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type statusModel struct {
}

func (m statusModel) Init() tea.Cmd {
	return nil
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "q":
			return m, returnToRoot
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m statusModel) View() tea.View {
	var s strings.Builder
	s.WriteString(normalStyle.Render("Status command!") + "\n\n")
	return tea.NewView(s.String())
}
