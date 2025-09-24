package components

import (
	"Callisto/network"
	"fmt"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewFullBody() *container.Split {

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
		container.NewTabItem("Body", bodyEntry),
		container.NewTabItem("Headers", headersEntry.Container),
	)

	var sendBtn *widget.Button
	sendBtn = widget.NewButton("Send Request", func() {
		sendBtn.Hide()
		client := network.NewClient(10 * time.Second)
		headers := headersEntry.GetHeaders()
		// Create request with body
		result, err := client.Send(
			HTTPMethods[selecty.SelectedIndex()],
			urlEntry.Text,
			bodyEntry.Text,
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

	fullBody := container.NewVSplit(upper, output)
	fullBody.SetOffset(0.1)

	return fullBody
}
