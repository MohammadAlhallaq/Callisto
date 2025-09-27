package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ReadOnlyMultiLineEntry struct {
	*widget.Entry
}

func NewResponseView() *ReadOnlyMultiLineEntry {
	e := &ReadOnlyMultiLineEntry{
		widget.NewMultiLineEntry(),
	}
	e.Wrapping = fyne.TextWrapWord
	e.SetPlaceHolder("Responses will appear here...")
	return e
}

// Prevent typing
func (e *ReadOnlyMultiLineEntry) TypedRune(r rune)            {}
func (e *ReadOnlyMultiLineEntry) TypedKey(key *fyne.KeyEvent) {}
