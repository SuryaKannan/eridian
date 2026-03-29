package cmd

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func initNewModel() newModel {
	ti := textinput.New()
	ti.Placeholder = "new language"
	ti.Focus()
	return newModel{ti}
}

type newModel struct {
	textInput textinput.Model
}

func (m newModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m newModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "esc":
			return m, returnToRoot
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m newModel) View() tea.View {
	var s strings.Builder

	s.WriteString(titleStyle.Render("\nEnter name of your new language!\n") + "\n\n")

	s.WriteString(normalStyle.Render("> "+m.textInput.Value()) + "\n\n")

	s.WriteString(titleStyle.Render("\nPress ESC to return to main screen.\n"))

	return tea.NewView(s.String())
}
