package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type HeadersEntry struct {
	Container     *fyne.Container
	rowsContainer *fyne.Container
	rows          []struct {
		Key   *widget.SelectEntry
		Value *widget.Entry
	}
}

func NewHeadersEntry() *HeadersEntry {
	h := &HeadersEntry{
		rowsContainer: container.NewVBox(),
	}

	defaultHeaders := []map[string]string{
		{"Content-Type": "application/json"},
		{"Accept": "application/json"},
	}

	options := []string{"Authorization"}
	addRow := func() {
		var row *fyne.Container

		keyEntry := widget.NewSelectEntry(options)
		keyEntry.SetPlaceHolder("Key")

		valueEntry := widget.NewEntry()
		valueEntry.SetPlaceHolder("Value")

		removeBtn := widget.NewButton("x", func() {
			h.rowsContainer.Remove(row)

			for i, r := range h.rows {
				if r.Key == keyEntry && r.Value == valueEntry {
					h.rows = append(h.rows[:i], h.rows[i+1:]...)
					break
				}
			}

			h.rowsContainer.Refresh()
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

		h.rowsContainer.Add(row)
		h.rowsContainer.Refresh()

		h.rows = append(h.rows, struct {
			Key   *widget.SelectEntry
			Value *widget.Entry
		}{keyEntry, valueEntry})
	}

	addBtn := widget.NewButton("+", func() {
		addRow()
	})

	for _, dh := range defaultHeaders {
		for k, v := range dh {
			key := widget.NewSelectEntry(options)
			key.SetText(k)
			key.Disable()

			val := widget.NewEntry()
			val.SetText(v)
			val.Disable()

			row := container.New(
				layout.NewGridLayout(2),
				key,
				val,
			)

			h.rowsContainer.Add(row)
			h.rowsContainer.Refresh()

			// Store references to the widgets for later use
			fHeader := struct {
				Key   *widget.SelectEntry
				Value *widget.Entry
			}{
				Key:   key,
				Value: val,
			}
			h.rows = append(h.rows, fHeader)
		}
	}

	h.Container = container.NewVBox(
		h.rowsContainer,
		addBtn,
	)
	return h
}

func (h *HeadersEntry) GetHeaders() map[string]string {
	headers := map[string]string{}

	for _, r := range h.rows {
		if r.Key.Text != "" {
			headers[r.Key.Text] = r.Value.Text
		}
	}
	return headers
}
