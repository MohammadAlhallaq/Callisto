package main

import (
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

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL (e.g. https://httpbin.org/post)")

	selecty := widget.NewSelect(methods, func(s string) {
	})
	selecty.Selected = "GET"

	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter json input here...")

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Respones will appear here...")
	output.Disable()
	content := container.NewStack(output)

	sendBtn := widget.NewButton("Send POST Request", func() {
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
			output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, wrapText(prettyJSON.String(), 120)))
			return
		}
		output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, wrapText(prettyJSON.String(), 120)))
	})

	hbox := container.New(
		layout.NewBorderLayout(nil, nil, selecty, nil),
		urlEntry,
		selecty,
	)

	content = container.NewBorder(
		container.NewVBox(hbox, bodyEntry, sendBtn),
		nil,
		nil,
		nil,
		output,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

func wrapText(s string, maxLine int) string {
	var result string
	for i := 0; i < len(s); i += maxLine {
		end := i + maxLine
		if end > len(s) {
			end = len(s)
		}
		result += s[i:end] + "\n"
	}
	return result
}
