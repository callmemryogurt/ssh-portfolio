package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func welcomeModel(m model) {}

func readModel(m model) {}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit every where in app after init screen!
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tickMsg:
		if m.ticks > 0 {
			m.ticks--
			return m, ticking()
		}
	case tea.KeyMsg:
		// This is anit bot protection to terminate connection on any key
		if m.ticks > 0 {
			return m, tea.Quit
		}
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, nil
}
