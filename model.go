package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

// Just a generic tea.Model to demo terminal information of ssh.
type model struct {
	term    string
	width   int
	height  int
	time    time.Time
	ticks   int
	help    help.Model
	keys    keyMap
	spinner spinner.Model
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
	m.spinner.Spinner = spinner.Dot
	m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return m
}

func (m model) Init() tea.Cmd {
	time.Sleep(time.Second)
	return tea.Batch(
		ticking(),
		m.spinner.Tick,
	)
}
