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
		ActionType{"Create User", "Ensure a user exists", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Created user 'deploy' with UID 1001"}
		}},
		ActionType{"Group User", "Ensure user belongs to a group", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Added 'deploy' to group 'sudo'"}
		}},
		ActionType{"SSH Key", "Ensure has ssh key", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Installed SSH key for user 'deploy' at ~/.ssh/authorized_keys"}
		}},
		ActionType{"Install Helix", "Install and configure helix editor", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Helix v24.03 installed to /usr/local/bin/hx"}
		}},
		ActionType{"Install Docker", "Install and configure Docker", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Docker version 25.0.2 installed and running"}
		}},
		ActionType{"Add Hostname", "Ensure hostname is set", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Hostname set to 'web-01.local'"}
		}},
		ActionType{"Reboot", "Reboot the server", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Reboot scheduled in 1 minute via 'shutdown -r'"}
		}},
		ActionType{"System Update", "Update all packages", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ All packages updated. 18 upgraded, 0 newly installed"}
		}},
		ActionType{"Install NGINX", "Install and configure NGINX", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ NGINX installed and running on port 80"}
		}},
		ActionType{"Install Node.js", "Install Node.js runtime", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Node.js v20.11.1 installed at /usr/local/bin/node"}
		}},
		ActionType{"Start Service", "Ensure a given service is started", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Service 'nginx' is now active (running)"}
		}},
		ActionType{"Stop Service", "Ensure a given service is stopped", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Service 'nginx' stopped successfully"}
		}},
		ActionType{"Enable Service", "Enable service to start on boot", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Service 'nginx' enabled via systemctl"}
		}},
		ActionType{"Disable Service", "Disable service from starting on boot", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Service 'nginx' disabled via systemctl"}
		}},
		ActionType{"Upload File", "Upload a file to target servers", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Uploaded ./config.yaml to /etc/myapp/config.yaml"}
		}},
		ActionType{"Download File", "Download file from remote server", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Downloaded /var/log/syslog to ./logs/syslog"}
		}},
		ActionType{"Run Shell Script", "Execute a shell script remotely", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Executed script: backup.sh (exit code 0)"}
		}},
		ActionType{"Install Go", "Install Go language environment", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Go 1.22.1 installed and GOROOT set to /usr/local/go"}
		}},
		ActionType{"Install Rust", "Install Rust tooling via rustup", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Rust 1.77.0 installed via rustup"}
		}},
		ActionType{"Install Tailscale", "Install and configure Tailscale VPN", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Tailscale installed. Connected as web-01@tailscale.net"}
		}},
		ActionType{"Set Timezone", "Ensure timezone is configured correctly", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Timezone set to UTC"}
		}},
		ActionType{"Add Cron Job", "Create or update a cron job", func() tea.Msg {
			time.Sleep(2 * time.Second)
			return ActionDoneMsg{Output: "✔ Cron job added: @daily /usr/local/bin/backup.sh"}
		}},
	})
	m.viewport = viewport.New(width, height)
	m.viewport.SetContent("Run an action to see output here...")
}

type ActionType struct {
	title       string
	description string
	executeFn   tea.Cmd
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
	case ActionDoneMsg:
		m.viewport.SetContent(msg.Output)
		m.actions.StopSpinner()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			cmds = append(cmds, m.actions.StartSpinner())
			actionType := m.actions.SelectedItem().(ActionType)
			cmds = append(cmds, actionType.executeFn)
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

type ActionDoneMsg struct {
	Output string
	Err    error
}
