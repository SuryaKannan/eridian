package cmd

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

var eridianTitle = `           
###################################        
###################################       

	░█▀▀░█▀▄░▀█▀░█▀▄░▀█▀░█▀█░█▀█
	░█▀▀░█▀▄░░█░░█░█░░█░░█▀█░█░█
	░▀▀▀░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀░▀


####################################
###################################                                                                                                  
`

var eridianDescription = `"fist my bump" - Rocky`

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F2A17C"))
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00604d")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F2A17C"))
)

type model struct {
	choices      []string
	descriptions []string
	cursor       int
	selected     map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"new", "use", "list", "label", "translate", "edit", "status", "clean"},
		descriptions: []string{
			"Create a new language dictionary",
			"Switch active language context",
			"Show all languages",
			"Capture mic audio and assign label",
			"Capture audio and return translation",
			"Manage and delete entries",
			"Show active language, dictionary size, and last rebuild",
			"Wipe all entries for the active language",
		},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := titleStyle.Render(eridianTitle) + "\n" + normalStyle.Render(lipgloss.NewStyle().Italic(true).Render(eridianDescription)) + "\n\n"

	for i, choice := range m.choices {
		if m.cursor == i {
			s += lipgloss.JoinHorizontal(lipgloss.Top, selectedStyle.Render("> "+choice+" ("), selectedStyle.Render(m.descriptions[i]+")")) + "\n"
		} else {
			s += normalStyle.Render("  "+choice) + "\n"
		}
	}

	s += titleStyle.Render("\nPress q to quit.\n")

	return tea.NewView(s)
}

var rootCmd = &cobra.Command{
	Use:   "eridian",
	Short: "Eridian is a language dictionary that semantically translates one language to another",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
