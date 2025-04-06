package progress

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Task struct {
	Label  string
	Action func(context.Context) error
}

type tuiModel struct {
	actions  list.Model
	viewport viewport.Model
	help     help.Model
	loaded   bool
	quitting bool
}

func (m *tuiModel) initActions(width, height int) {
	m.actions = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.actions.Title = "Actions"
	m.actions.SetItems([]list.Item{
		ActionType{"Create User", "Ensure a user exists"},
		ActionType{"Group User", "Ensure user belongs to a group"},
		ActionType{"SSH Key", "Ensure has ssh key"},
		ActionType{"Install Helix", "Install and configure helix editor"},
		ActionType{"Install Docker", "Install and configure Docker"},
		ActionType{"Add Hostname", "Ensure hostname is set"},
		ActionType{"Reboot", "Reboot the server"},
		ActionType{"System Update", "Update all packages"},
		ActionType{"Install NGINX", "Install and configure NGINX"},
		ActionType{"Install Node.js", "Install Node.js runtime"},
		ActionType{"Start Service", "Ensure a given service is started"},
		ActionType{"Stop Service", "Ensure a given service is stopped"},
		ActionType{"Enable Service", "Enable service to start on boot"},
		ActionType{"Disable Service", "Disable service from starting on boot"},
		ActionType{"Upload File", "Upload a file to target servers"},
		ActionType{"Download File", "Download file from remote server"},
		ActionType{"Run Shell Script", "Execute a shell script remotely"},
		ActionType{"Install Go", "Install Go language environment"},
		ActionType{"Install Rust", "Install Rust tooling via rustup"},
		ActionType{"Install Tailscale", "Install and configure Tailscale VPN"},
		ActionType{"Set Timezone", "Ensure timezone is configured correctly"},
		ActionType{"Add Cron Job", "Create or update a cron job"},
	})
	m.viewport = viewport.New(width, height)
	m.viewport.SetContent("Run an action to see output here...")
}

type ActionType struct {
	title       string
	description string
}

func (a ActionType) Title() string {
	return a.title
}

func (a ActionType) Description() string {
	return a.description
}

// implement the list.Item interface
func (a ActionType) FilterValue() string {
	return a.title
}

func NewTUIProgram() {
	m := tuiModel{
		actions:  list.Model{},
		viewport: viewport.Model{},
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}

type taskDoneMsg struct {
	err error
}

type tickMsg time.Time

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
			m.initActions(msg.Width, msg.Height)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			cmds = append(cmds, m.actions.ToggleSpinner())
			m.viewport.SetContent("Running...")
		}
	}
	var actionsCmd, viewportCmd tea.Cmd
	m.actions, actionsCmd = m.actions.Update(msg)
	m.viewport, viewportCmd = m.viewport.Update(msg)
	cmds = append(cmds, actionsCmd, viewportCmd)
	return m, tea.Batch(cmds...)
}

func (m tuiModel) View() string {
	if m.quitting {
		return ""
	}

	if !m.loaded {
		return "loading..."
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.actions.View(),
		m.viewport.View(),
	)
	// return lipgloss.JoinVertical(lipgloss.Left, board, m.help.View(keys))
}
