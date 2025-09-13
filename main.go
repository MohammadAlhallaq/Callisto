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
	"fyne.io/fyne/v2/widget"
)

type Respones struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Callisto")

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL (e.g. https://httpbin.org/post)")

	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter json input here...")

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Respones will appear here...")
	output.Disable()
	content := container.NewStack(output)

	sendBtn := widget.NewButton("Send POST Request", func() {
		client := &http.Client{Timeout: 10 * time.Second}

		// Create request with body
		req, err := http.NewRequest("POST", urlEntry.Text, bytes.NewBuffer([]byte(bodyEntry.Text)))
		if err != nil {
			output.SetText(fmt.Sprintf("Error creating request: %v", err))
			return
		}

		// Set header (default JSON)
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
			output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, string(body)))
			return
		}
		output.SetText(fmt.Sprintf("Status: %s\n\n%s", resp.Status, wrapText(prettyJSON.String(), 120)))
	})

	content = container.NewBorder(
		container.NewVBox(urlEntry, bodyEntry, sendBtn), // top
		nil,    // bottom
		nil,    // left
		nil,    // right
		output, // center (will expand)
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
