package main

import (
	"os/exec"

	"github.com/jtheiss19/Go-Open-Api/communication"

	"github.com/jtheiss19/Go-Open-Api/gateway"
	"gitlab.com/tslocum/cview"
)

func sidebarContent() cview.Primitive {
	grid := cview.NewGrid().SetColumns(0).SetRows(5, 5, 0, 5).SetBorders(true)

	button := cview.NewButton("Create Website Instance")
	button.SetSelectedFunc(func() {
		exec.Command("go", "run", "../website/.").Start()
	})
	button.SetBorder(true)

	button2 := cview.NewButton("Start Gateway").SetSelectedFunc(func() {
		if gateway.Status() {
			go gateway.Stop()
		} else {
			go gateway.Start("8081")
		}
	})
	button2.SetBorder(true)

	button3 := cview.NewButton("Exit & Close Open Apps").SetSelectedFunc(func() {
		gateway.Stop()
		communication.SendReq()
		app.Stop()
	})
	button3.SetBorder(true)

	grid.AddItem(button, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(button2, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(button3, 3, 0, 1, 1, 0, 0, false)

	return grid
}
