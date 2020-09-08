package main

import (
	"io/ioutil"

	"gitlab.com/tslocum/cview"
)

var logs *cview.TextView
var filePath string

func logView() cview.Primitive {

	grid := cview.NewGrid().SetColumns(20, 20, 0, 20).SetRows(1, 3, 0)

	updateButton := cview.NewButton("Update Log").SetSelectedFunc(func() {
		updateLogText()
	})
	updateButton.SetBorder(true)

	gatewayLogs := cview.NewButton("Show Gateway Logs").SetSelectedFunc(func() {
		updateFilePath("./logs/gateway_log")
	})
	gatewayLogs.SetBorder(true)

	websiteLogs := cview.NewButton("Show Website Logs").SetSelectedFunc(func() {
		updateFilePath("./logs/website_log")
	})
	websiteLogs.SetBorder(true)

	Title := cview.NewTextView().SetText("Log Viewer").SetTextAlign(cview.AlignCenter)
	logs = cview.NewTextView().SetScrollable(true)
	logs.SetBorder(true)
	updateLogText()

	grid.AddItem(Title, 0, 0, 1, 4, 0, 0, false)
	grid.AddItem(gatewayLogs, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(websiteLogs, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(updateButton, 1, 3, 1, 1, 0, 0, false)
	grid.AddItem(logs, 2, 0, 1, 4, 0, 0, false)

	return grid
}

func updateLogText() {
	f, _ := ioutil.ReadFile(filePath)
	logs.SetText(string(f))
}

func updateFilePath(newFilePath string) {
	filePath = newFilePath
	updateLogText()
}
