package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewRequestHistoryList() *widget.List {

	var data = []string{}

	if len(data) == 0 {
		return nil
	}

	return widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Request Item")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[id])
		},
	)

}
