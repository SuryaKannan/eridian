package cmd

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/SuryaKannan/eridian/internal/config"
	"github.com/SuryaKannan/eridian/internal/store"
)

type cleanModel struct {
	screenName      string
	currentLanguage string
	errorMsg        string
}

func initCleanModel() cleanModel {
	cfg := config.FetchConfig()
	return cleanModel{
		screenName:      ScreenName[Clean],
		currentLanguage: cfg.ActiveLanguage,
	}
}

func (m cleanModel) Init() tea.Cmd {
	return nil
}

func (m cleanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "esc":
			return m, returnToRoot
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.currentLanguage != "" {
				err := store.NewStore(m.currentLanguage).DeleteLanguageDB()
				if err != nil {
					m.errorMsg = err.Error()
					m.currentLanguage = ""
					return m, nil
				}
				return m, returnToRoot
			}
		}
	}
	return m, nil
}

func (m cleanModel) View() tea.View {
	var s strings.Builder

	s.WriteString(selectedStyle.Render("(home/"+m.screenName+")") + "\n\n")

	if m.currentLanguage != "" {

		s.WriteString(normalStyle.Render("Are you sure you would like to delete: "))
		s.WriteString(titleStyle.Render(m.currentLanguage))
		s.WriteString(normalStyle.Render(" (press enter to continue)" + "\n\n"))

	} else if m.errorMsg != "" {
		s.WriteString(normalStyle.Render("Failed to delete: " + m.errorMsg))
		s.WriteString(normalStyle.Render(", please exit to menu\n\n"))

	} else {

		s.WriteString(normalStyle.Render("No active language set!" + "\n\n"))
	}

	s.WriteString(titleStyle.Render("\nPress ESC to return home.\n"))

	view := tea.NewView(s.String())
	view.AltScreen = true
	return view
}
