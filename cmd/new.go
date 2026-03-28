package cmd

import (
	tea "charm.land/bubbletea/v2"
)

type newModel struct {
}

func (m newModel) Init() tea.Cmd {
	return nil
}

func returnToRoot() tea.Msg {
	return backToMenuMsg{}
}

func (m newModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (m newModel) View() tea.View {

	return tea.NewView("Hi")
}
