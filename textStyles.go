package main

import "github.com/charmbracelet/lipgloss"

var selected = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
var highlighted = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Bold(true)

var activeGroup = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("4"))
var inactiveGroup = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("7"))
