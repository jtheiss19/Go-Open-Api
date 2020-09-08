package main

import (
	"gitlab.com/tslocum/cview"
)

func init() {
	mainWindow.AddPage("logView", logView(), true, false)
	mainWindow.AddPage("settingsView", settingsView(), true, false)
}

func NewMenu() cview.Primitive {
	grid := cview.NewGrid().SetColumns(0).SetRows(5, 5, 0, 5).SetBorders(true)

	logViewButton := cview.NewButton("Log Viewer").SetSelectedFunc(func() {
		mainWindow.SwitchToPage("logView")
	})
	logViewButton.SetBorder(true)

	settingsButton := cview.NewButton("Settings").SetSelectedFunc(func() {
		mainWindow.SwitchToPage("settingsView")
	})
	settingsButton.SetBorder(true)

	grid.AddItem(logViewButton, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(settingsButton, 1, 0, 1, 1, 0, 0, false)

	return grid
}
