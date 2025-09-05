package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	Metadata FrontmatterMetaData
}

func (i Item) Title() string       { return i.Metadata.Title }
func (i Item) Description() string { return i.Metadata.Subject + " " + i.Metadata.Description }
func (i Item) FilterValue() string { return i.Metadata.Subject + " " + i.Metadata.Title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			i := m.list.SelectedItem().(Item)
			labsLocation := GetLabsLocation()
			preferredEditor := viper.GetString("preferred_editor")
			labLocation := fmt.Sprintf("%s/%s/%s-%d-%s", labsLocation, i.Metadata.Subject, i.Metadata.Title, i.Metadata.UniYearAndSemester, i.Metadata.Date)
			args := []string{labLocation}
			cmd := exec.Command(preferredEditor, args...)
			cmd.Run()

			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func RUNTUI(items []list.Item, title string) {

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = title

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
