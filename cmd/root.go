package cmd

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"charm.land/bubbles/v2/spinner"
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

var eridianQuotes = []string{
	`"fist my bump"`,
	`"amaze amaze amaze"`,
	`"usually you not stupid. Why stupid, question?"`,
	`"hold screws better."`,
	`"you poked it with a stick?"`,
}

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F2A17C"))
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00604d")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F2A17C"))
	italicStyle   = normalStyle.Italic(true)
)

type menuItem struct {
	choice, description string
}

type model struct {
	items   []menuItem
	quote   string
	cursor  int
	spinner spinner.Model
}

func initialModel(quoteIndex int) model {
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00604d"))
	return model{
		items: []menuItem{
			{choice: "new", description: "Create a new language dictionary"},
			{choice: "use", description: "Switch active language context"},
			{choice: "list", description: "Show all languages"},
			{choice: "label", description: "Capture mic audio and assign label"},
			{choice: "translate", description: "Capture audio and return translation"},
			{choice: "edit", description: "Manage and delete entries"},
			{choice: "status", description: "Show active language, dictionary size, and last rebuild"},
			{choice: "clean", description: "Wipe all entries for the active language"},
		},
		quote:   eridianQuotes[quoteIndex],
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
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
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter", "space":

		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString(titleStyle.Render(eridianTitle) + "\n" + italicStyle.Render(m.quote) + " " + m.spinner.View() + "\n\n")

	for i, item := range m.items {
		if m.cursor == i {
			s.WriteString(selectedStyle.Render("> "+item.choice+" ("+item.description+")") + "\n")
		} else {
			s.WriteString(normalStyle.Render("  "+item.choice) + "\n")
		}
	}

	s.WriteString(titleStyle.Render("\nPress q to quit.\n"))

	return tea.NewView(s.String())
}

var rootCmd = &cobra.Command{
	Use:   "eridian",
	Short: "Eridian is a language dictionary that semantically translates one language to another",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel(rand.IntN(len(eridianQuotes))))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running Eridian! Check setup: %v", err)
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
