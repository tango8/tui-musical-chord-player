package main

import (
	//tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	pressedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
)
