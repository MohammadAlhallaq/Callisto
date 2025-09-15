package components

import "fyne.io/fyne/v2/widget"

// NewURLEntry creates a reusable URL input field with placeholder text.
func NewURLEntry() *widget.Entry {
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL (e.g. https://httpbin.org/post)")
	return urlEntry
}
