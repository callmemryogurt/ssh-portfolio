package main

import (
	"fmt"
	"time"
)

func botView(m model) string {
	s := "This is init screen\n"
	s += "I set this up to prevent SSH bots to do something harmful\n"
	s += fmt.Sprintf("Please wait till timer go down: %v\n", m.ticks)
	return s
}

func welcomeView(m model) string {
	s := "Your term is %s\n"
	s += fmt.Sprintf("Ticker %v\n", m.ticks)
	s += "Your window size is x: %d y: %d\n"
	s += "Time: " + m.time.Format(time.RFC1123) + "\n\n"
	// s += "Press 'q' to quit\n"
	return fmt.Sprintf(s, m.term, m.width, m.height)
}

func (m model) View() string {
	helpView := m.help.View(m.keys)
	if m.ticks > 0 {
		k := botView(m)
		return k
	}
	return welcomeView(m) + helpView
}
