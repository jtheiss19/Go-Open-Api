package main

import (
	"gitlab.com/tslocum/cview"
)

func settingsView() cview.Primitive {

	grid := cview.NewGrid().SetColumns(20, 20, 0, 20).SetRows(1, 1, 1, 1, 1, 0)

	settings1 := cview.NewInputField().SetLabel("Setting One")
	settings2 := cview.NewInputField().SetLabel("Setting Two")
	settings3 := cview.NewInputField().SetLabel("Setting Three")
	settings4 := cview.NewInputField().SetLabel("Setting Four")
	settings5 := cview.NewInputField().SetLabel("Setting Five")

	grid.AddItem(settings1, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(settings2, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(settings3, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(settings4, 3, 0, 1, 1, 0, 0, false)
	grid.AddItem(settings5, 4, 0, 1, 1, 0, 0, false)

	return grid
}
