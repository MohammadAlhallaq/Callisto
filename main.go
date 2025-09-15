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

func main() {
	a := app.New()
	w := a.NewWindow("Callisto")

	// INITIALIZE URL INPUT
	urlEntry := components.NewURLEntry()
	// INITIALIZE RESPONSE WIDGET
	output := components.NewResponseView()
	// INITIALIZE BODY REQUEST INPUT
	bodyEntry := components.NewRequestBody()
	// INITIALIZE METHODS dROPDOWN
	selecty := components.NewDropdownMethods()

	// INITIALIZE HEADERS INPUT

	tabs := container.NewAppTabs(
		container.NewTabItem("Body", bodyEntry),
		// container.NewTabItem("Headers", headersEntry),
	)

	sendBtn := widget.NewButton("Send Request", func() {
		selecty.SelectedIndex()
		client := &http.Client{Timeout: 10 * time.Second}
		method := components.HTTPMethods[selecty.SelectedIndex()]

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
