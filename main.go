package main

import (
	// "bytes"
	// "fmt"
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "io"
	// "net/http"
	// "time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go HTTP Client")

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL (e.g. https://httpbin.org/post)")

	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder(`Enter JSON body here...`)

	// output := widget.NewMultiLineEntry()
	// output.SetPlaceHolder("Response will appear here...")

	// sendBtn := widget.NewButton("Send POST Request", func() {
	// 	client := &http.Client{Timeout: 10 * time.Second}

	// 	// Create request with body
	// 	req, err := http.NewRequest("POST", urlEntry.Text, bytes.NewBuffer([]byte(bodyEntry.Text)))
	// 	if err != nil {
	// 		output.SetText(fmt.Sprintf("Error creating request: %v", err))
	// 		return
	// 	}

	// 	// Set header (default JSON)
	// 	req.Header.Set("Content-Type", "application/json")

	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		output.SetText(fmt.Sprintf("Error sending request: %v", err))
	// 		return
	// 	}
	// 	defer resp.Body.Close()

	// 	body, _ := io.ReadAll(resp.Body)
	// 	output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, string(body)))
	// })

	w.SetContent(container.NewVBox(
		urlEntry,
		bodyEntry,
	))

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
