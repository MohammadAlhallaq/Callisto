package components

import "fyne.io/fyne/v2/widget"

var HTTPMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

func NewDropdownMethods() *widget.Select {
	selecty := widget.NewSelect(HTTPMethods, func(s string) {
	})
	selecty.Selected = "GET"
	return selecty
}
