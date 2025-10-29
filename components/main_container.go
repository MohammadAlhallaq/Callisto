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
	sendBtn = widget.NewButton("Send Request", func() {
		sendBtn.Hide()
		client := network.NewClient(10 * time.Second)
		headers := headersEntry.GetHeaders()
		var body *bytes.Buffer
		var contentType string
		var err error

		if bodyEntry.mode == "raw" {
			body, contentType, err = bodyEntry.GetRawData()
			if err != nil {
				dialog.ShowError(err, w)
			}
			headers["Content-Type"] = contentType
		} else {
			body, contentType = bodyEntry.GetFormData()
			headers["Content-Type"] = contentType
		}

		result, err := client.Send(
			HTTPMethods[selecty.SelectedIndex()],
			urlEntry.Text,
			body,
			headers,
		)

		sendBtn.Show()

		if err != nil {
			output.SetText(fmt.Sprintf("Error sending request: %v", err))
			return
		}
		output.SetText(result)
	})
	sendBtn.Importance = widget.WarningImportance

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

	content := container.NewBorder(nil, nil, nil, nil, output)

	fullBody := container.NewVSplit(upper, content)
	fullBody.SetOffset(0.5)

	return fullBody
}
