package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"

	// "github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

// Just a generic tea.Model to demo terminal information of ssh.
type model struct {
	term     string
	width    int
	height   int
	time     time.Time
	ticks    int
	help     help.Model
	keys     keyMap
	spinner  spinner.Model
	viewport viewport.Model
}

type timeMsg time.Time

type tickMsg time.Time

func ticking() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func initModel(pty ssh.Pty) model {
	m := model{
		term:   pty.Term,
		width:  pty.Window.Width,
		height: pty.Window.Height,
		time:   time.Now(),
		ticks:  5,
		help:   help.New(),
		keys:   keys,
	}

	m.spinner = spinner.New()
	m.spinner.Spinner = spinner.Pulse
	m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(m.width, m.height-5)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.width),
	)
	if err != nil {
		panic(err)
	}

	str, err := renderer.Render(content)
	if err != nil {
		panic(err)
	}

	vp.SetContent(str)
	m.viewport = vp

	return m
}

func (m model) Init() tea.Cmd {
	time.Sleep(time.Second)
	return tea.Batch(
		ticking(),
		m.spinner.Tick,
	)
}
