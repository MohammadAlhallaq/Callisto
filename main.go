package main

import (
	"Callisto/components"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var methods = []string{"POST", "GET", "PATCH", "PUT", "DELETE"}

func main() {
	a := app.New()
	w := a.NewWindow("Callisto")

	// INITIALIZE URL INPUT
	urlEntry := components.NewURLEntry()

	// INITIALIZE HEADERS INPUT
	rowsContainer := container.NewVBox()
	addRow := func() {
		var row *fyne.Container

		keyEntry := widget.NewEntry()
		keyEntry.SetPlaceHolder("Key")

		valueEntry := widget.NewEntry()
		valueEntry.SetPlaceHolder("Value")

		removeBtn := widget.NewButton("x", func() {
			rowsContainer.Remove(row)
			rowsContainer.Refresh()
		})
		
		row = container.New(layout.NewGridLayout(3), keyEntry, valueEntry, removeBtn)

		rowsContainer.Add(row)
		rowsContainer.Refresh()
	}

	addBtn := widget.NewButton("+", func() {
		addRow()
	})

	headersEntry := container.NewVBox(
		rowsContainer,
		addBtn,
	)

	// INITIALIZE BODY REQUEST INPUT
	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter json input here...")

	// INITIALIZE RESPONSE WIDGET
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Respones will appear here...")
	output.Disable()
	output.Wrapping = fyne.TextWrapWord

	selecty := widget.NewSelect(methods, func(s string) {
	})
	selecty.Selected = "GET"

	tabs := container.NewAppTabs(
		container.NewTabItem("Body", bodyEntry),
		container.NewTabItem("Headers", headersEntry),
	)

	sendBtn := widget.NewButton("Send Request", func() {
		selecty.SelectedIndex()
		client := &http.Client{Timeout: 10 * time.Second}
		method := methods[selecty.SelectedIndex()]

		// Create request with body
		req, err := http.NewRequest(method, urlEntry.Text, bytes.NewBuffer([]byte(bodyEntry.Text)))
		if err != nil {
			output.SetText(fmt.Sprintf("Error creating request: %v", err))
			return
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			output.SetText(fmt.Sprintf("Error sending request: %v", err))
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
			// If response is not valid JSON, just show it as is
			output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, prettyJSON.String()))
			return
		}
		output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, prettyJSON.String()))
	})

	hbox := container.New(
		layout.NewBorderLayout(nil, nil, selecty, nil),
		urlEntry,
		selecty,
	)

	upper := container.NewVBox(
		hbox,
		widget.NewLabel(""),
		tabs,
		container.NewPadded(sendBtn),
	)

	split := container.NewVSplit(upper, output)
	split.SetOffset(0.1)

	w.SetContent(split)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
