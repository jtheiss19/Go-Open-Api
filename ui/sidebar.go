package main

import (
	"github.com/jtheiss19/Go-Open-Api/gateway"
	"github.com/jtheiss19/Go-Open-Api/website"
	"gitlab.com/tslocum/cview"
)

func sidebarContent() cview.Primitive {
	grid := cview.NewGrid().SetColumns(0).SetRows(5, 5, 5, 0).SetBorders(true)

	button := cview.NewButton("Start Website").SetSelectedFunc(func() {
		go website.Start("8080")
	})
	button.SetBorder(true)

	button2 := cview.NewButton("Start Gateway").SetSelectedFunc(func() {
		go gateway.Start("8081")
	})
	button2.SetBorder(true)

	button3 := cview.NewButton("Exit").SetSelectedFunc(func() {
		app.Stop()
	})
	button3.SetBorder(true)

	grid.AddItem(button, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(button2, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(button3, 2, 0, 1, 1, 0, 0, false)

	return grid
}
