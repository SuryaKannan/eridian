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

type Screen int

const (
	Root Screen = iota
	New
	Use
	List
	Label
	Translate
	Edit
	Status
	Clean
)

var screenName = map[Screen]string{
	Root:      "root",
	New:       "new",
	Use:       "use",
	List:      "list",
	Label:     "label",
	Translate: "translate",
	Edit:      "edit",
	Status:    "status",
	Clean:     "clean",
}

func (s Screen) String() string {
	return screenName[s]
}

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F2A17C"))
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00604d")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F2A17C"))
	italicStyle   = normalStyle.Italic(true)
)

type menuItem struct {
	choice      Screen
	description string
}

type rootModel struct {
	activeScreen Screen
	items        []menuItem
	quote        string
	cursor       int
	spinner      spinner.Model
}

func initialModel(quoteIndex int, activeScreen Screen) rootModel {
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00604d"))
	return rootModel{
		activeScreen: activeScreen,
		items: []menuItem{
			{choice: New, description: "Create a new language dictionary"},
			{choice: Use, description: "Switch active language context"},
			{choice: List, description: "Show all languages"},
			{choice: Label, description: "Capture mic audio and assign label"},
			{choice: Translate, description: "Capture audio and return translation"},
			{choice: Edit, description: "Manage and delete entries"},
			{choice: Status, description: "Show active language, dictionary size, and last rebuild"},
			{choice: Clean, description: "Wipe all entries for the active language"},
		},
		quote:   eridianQuotes[quoteIndex],
		spinner: s,
	}
}

func (m rootModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m rootModel) invokeCmd(cmd string) {
	switch cmd {
	case "new":

	}
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			// todo
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m rootModel) View() tea.View {
	var s strings.Builder
	s.WriteString(titleStyle.Render(eridianTitle) + "\n" + italicStyle.Render(m.quote) + " " + m.spinner.View() + "\n\n")

	for i, item := range m.items {
		if m.cursor == i {
			s.WriteString(selectedStyle.Render("> "+item.choice.String()+" ("+item.description+")") + "\n")
		} else {
			s.WriteString(normalStyle.Render("  "+item.choice.String()) + "\n")
		}
	}

	s.WriteString(titleStyle.Render("\nPress q to quit.\n"))

	return tea.NewView(s.String())
}

var rootCmd = &cobra.Command{
	Use:   "eridian",
	Short: "Eridian is a language dictionary that semantically translates one language to another",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel(rand.IntN(len(eridianQuotes)), Root))
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
