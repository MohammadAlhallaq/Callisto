package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewResponseView() *widget.Entry {
	// INITIALIZE RESPONSE WIDGET
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Respones will appear here...")
	output.Disable()
	output.Wrapping = fyne.TextWrapWord
	return output
}
