package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Just a generic tea.Model to demo terminal information of ssh.
type model struct {
	term   string
	width  int
	height int
	time   time.Time
	ticks  int
	help   help.Model
	keys   keyMap
}

type timeMsg time.Time

type tickMsg time.Time

func ticking() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	time.Sleep(time.Second)
	return tea.Batch(
		ticking(),
	)
}
