package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/SuryaKannan/eridian/internal/config"
)

type listModel struct {
	screenName     string
	languages      []string
	activeLanguage string
	cursor         int
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func initListModel() listModel {
	cfg := config.FetchConfig()
	return listModel{
		screenName:     ScreenName[List],
		activeLanguage: cfg.ActiveLanguage,
		languages:      cfg.Languages,
	}
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.languages)-1 {
				m.cursor++
			}
		case "esc":
			return m, returnToRoot
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if len(m.languages) == 0 {
				return m, nil
			}
			config.SetActiveLanguage(m.languages[m.cursor])
			return m, returnToRoot
		}

	}
	return m, nil
}

func (m listModel) View() tea.View {
	var s strings.Builder

	s.WriteString(selectedStyle.Render("(home/"+m.screenName+")") + "\n\n")

	if len(m.languages) == 0 {
		s.WriteString(titleStyle.Render("No languages yet! Create one first." + "\n\n"))
	} else {
		s.WriteString(titleStyle.Render("Current language:") + " " + selectedStyle.Render(m.activeLanguage) + "\n\n")
		for i, language := range m.languages {
			if m.cursor == i {
				s.WriteString(selectedStyle.Render("> "+language) + normalStyle.Render(" — hit enter to switch") + "\n")
			} else {
				s.WriteString(normalStyle.Render("  "+language) + "\n")
			}
		}
	}

	s.WriteString(titleStyle.Render("\nPress ESC to return home.\n"))

	view := tea.NewView(s.String())
	view.AltScreen = true
	return view
}
