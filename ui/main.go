package main

import (
	"gitlab.com/tslocum/cview"
)

var app = cview.NewApplication()
var mainWindow *cview.Pages = cview.NewPages()

func main() {
	newPrimitive := func(text string) cview.Primitive {
		return cview.NewTextView().
			SetTextAlign(cview.AlignCenter).
			SetText(text)
	}
	menu := NewMenu()
	sideBar := sidebarContent()

	grid := cview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Go Open API Gateway"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Placeholder Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(mainWindow, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(mainWindow, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
