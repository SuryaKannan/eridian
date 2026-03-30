package cmd

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/SuryaKannan/eridian/internal/config"
	"github.com/SuryaKannan/eridian/internal/store"
)

func initNewModel() newModel {
	ti := textinput.New()
	ti.Placeholder = "new language"
	ti.Focus()
	return newModel{
		screenName: ScreenName[New],
		textInput:  ti,
		result:     "",
	}
}

type newModel struct {
	screenName string
	result     string
	textInput  textinput.Model
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
		case "enter":
			if m.result != "" {
				m.result = ""
				m.textInput.Reset()
				return m, nil
			}
			s := store.NewStore(m.textInput.Value())
			if err := s.CreateLanguageDB(); err != nil {
				m.result = err.Error()
				return m, nil
			}
			config.SetActiveLanguage(m.textInput.Value())
			return m, returnToRoot
		}

		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m newModel) View() tea.View {
	var s strings.Builder

	s.WriteString(selectedStyle.Render("(home/" + m.screenName + ")" + "\n\n"))

	if m.result == "" {
		s.WriteString(normalStyle.Render("\nEnter the name of your new language!\n") + "\n\n")

		s.WriteString(normalStyle.Render("> "+m.textInput.Value()) + "\n\n")

	} else {
		s.WriteString(normalStyle.Render(m.result) + "\n\n")
	}

	s.WriteString(titleStyle.Render("\nPress ESC to return home.\n"))

	return tea.NewView(s.String())
}
