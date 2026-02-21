package components

import (
	"Callisto/network"
	"bytes"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewFullBody(w fyne.Window) *container.Split {

	// INITIALIZE URL INPUT
	urlEntry := NewURLEntry()
	// INITIALIZE RESPONSE WIDGET
	output := NewResponseView()
	// INITIALIZE BODY REQUEST INPUT
	bodyEntry := NewRequestBody()
	// INITIALIZE METHODS DROPDOWN
	selecty := NewDropdownMethods()
	// INITIALIZE HEADERS ENTRY
	headersEntry := NewHeadersEntry()

	tabs := container.NewAppTabs(
		container.NewTabItem("Body", bodyEntry.Container),
		container.NewTabItem("Headers", headersEntry.Container),
	)

	var sendBtn *widget.Button
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Hide()

	sendBtn = widget.NewButton("Send", func() {
		client := network.NewClient(60 * time.Second)
		headers := headersEntry.GetHeaders()
		var body *bytes.Buffer
		var contentType string
		var err error

		if bodyEntry.mode == "raw" {
			body, contentType, err = bodyEntry.GetRawData()
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			headers["Content-Type"] = contentType
		} else {
			body, contentType = bodyEntry.GetFormData()
			headers["Content-Type"] = contentType
		}

		progressBar.Show()
		sendBtn.Hide()
		go func() {
			result, err := client.Send(
				HTTPMethods[selecty.SelectedIndex()],
				urlEntry.Text,
				body,
				headers,
			)
			fyne.Do(func() {
				if err != nil {
					output.SetText(fmt.Sprintf("Error sending request: %v", err))
				} else {
					output.SetText(result)
				}
				sendBtn.Show()
				progressBar.Hide()
			})
		}()
	})
	sendBtn.Importance = widget.WarningImportance

	// URL bar: [Method ▼] [URL_______________] [Send]
	sendStack := container.NewStack(sendBtn, progressBar)
	urlBar := container.New(
		layout.NewBorderLayout(nil, nil, selecty, sendStack),
		selecty,
		sendStack,
		urlEntry,
	)

	upper := container.NewVBox(
		urlBar,
		widget.NewSeparator(),
		tabs,
	)

	fullBody := container.NewVSplit(upper, output)
	fullBody.SetOffset(0.4)

	return fullBody
}
