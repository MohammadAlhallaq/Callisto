package components

import "fyne.io/fyne/v2/widget"

func NewRequestBody() *widget.Entry {
	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter json input here...")
	return bodyEntry
}
