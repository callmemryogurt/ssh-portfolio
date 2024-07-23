package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// func welcomeModel(m model) {}

// func readModel(m model) {}

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
		log.Info("Msg Window resize")
		m.height = msg.Height
		m.width = msg.Width
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
		m.updateSize()
		_, y := m.viewport.Update(msg)
		return m, y

	case tickMsg:
		log.Info("Got tick")
		if m.ticks > 0 {
			m.ticks--
			return m, ticking()
		}
	case tea.KeyMsg:
		log.Info("Got key")
		// This is anit bot protection to terminate connection on any key
		if m.ticks > 0 {
			return m, tea.Quit
		}
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			m.updateSize()
			log.Info("Got key to change help")
		default:
			log.Info("Got key default")
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		log.Info("Got key default big loop")
		if m.ticks > 0 {
			log.Info("Got default to update spinner")
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}
