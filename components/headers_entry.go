package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewHeadersEntry() *fyne.Container {

	rowsContainer := container.NewVBox()
	addRow := func() {
		var row *fyne.Container

		keyEntry := widget.NewEntry()
		keyEntry.SetPlaceHolder("Key")

		valueEntry := widget.NewEntry()
		valueEntry.SetPlaceHolder("Value")

		removeBtn := widget.NewButton("x", func() {
			rowsContainer.Remove(row)
			rowsContainer.Refresh()
		})

		fields := container.New(
			layout.NewGridLayout(2),
			keyEntry,
			valueEntry,
		)

		row = container.New(layout.NewBorderLayout(nil, nil, nil, removeBtn),
			fields,
			removeBtn,
		)
		rowsContainer.Add(row)
		rowsContainer.Refresh()
	}

	addBtn := widget.NewButton("+", func() {
		addRow()
	})

	headersEntry := container.NewVBox(
		rowsContainer,
		addBtn,
	)
	return headersEntry
}
