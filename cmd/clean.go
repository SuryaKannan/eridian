package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type cleanModel struct {
}

func (m cleanModel) Init() tea.Cmd {
	return nil
}

func (m cleanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (m cleanModel) View() tea.View {
	var s strings.Builder
	s.WriteString(normalStyle.Render("Clean command!") + "\n\n")
	return tea.NewView(s.String())
}
