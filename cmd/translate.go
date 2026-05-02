package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

func initTranslateModel() translateModel {
	return translateModel{
		screenName: ScreenName[Translate],
	}
}

type translateModel struct {
	screenName string
}

func (m translateModel) Init() tea.Cmd {
	return nil
}

func (m translateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (m translateModel) View() tea.View {
	var s strings.Builder

	s.WriteString(selectedStyle.Render("(home/"+m.screenName+")") + "\n\n")

	s.WriteString(titleStyle.Render("\nPress ESC to return home.\n"))

	view := tea.NewView(s.String())
	view.AltScreen = true
	return view
}
