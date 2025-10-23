package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type BodyEntry struct {
	Container     *fyne.Container
	rowsContainer *fyne.Container
	rows          []struct {
		Key   *widget.Entry
		Value *widget.Entry
	}
	rawEntry *widget.Entry
	radio    *widget.RadioGroup
	mode     string
}

func NewRequestBody() *BodyEntry {
	b := &BodyEntry{
		rowsContainer: container.NewVBox(),
		rawEntry:      widget.NewMultiLineEntry(),
		mode:          "raw",
	}

	b.rawEntry.SetPlaceHolder("Enter JSON input here...")

	b.radio = widget.NewRadioGroup([]string{"form-data", "raw"}, nil)
	b.radio.Horizontal = true

	b.Container = container.NewVBox(
		b.radio,
		b.rawEntry,
	)

	b.radio.OnChanged = func(selected string) {
		b.switchMode(selected)
	}

	b.radio.Horizontal = true
	b.radio.SetSelected("raw")

	return b
}

func (b *BodyEntry) switchMode(mode string) {
	b.mode = mode
	b.Container.RemoveAll()
	b.Container.Add(b.radio)

	addRow := func() {
		key := widget.NewEntry()
		key.SetPlaceHolder("Key")

		value := widget.NewEntry()
		value.SetPlaceHolder("Value")

		fields := container.New(
			layout.NewGridLayout(2),
			key,
			value,
		)
		row := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
			fields,
		)

		b.rowsContainer.Add(row)
	}

	if mode == "raw" {
		b.rawEntry.SetPlaceHolder("Enter JSON input here...")
		b.Container.Add(b.rawEntry)
	} else {
		b.Container.Add(b.rowsContainer)
		b.Container.Add(widget.NewButton("+", addRow))
	}
}

func (b *BodyEntry) GetHeaders() map[string]string {
	headers := map[string]string{}

	for _, r := range b.rows {
		if r.Key.Text != "" {
			headers[r.Key.Text] = r.Value.Text
		}
	}
	return headers
}
