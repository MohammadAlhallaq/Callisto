package components

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewAccountBar(w fyne.Window) *fyne.Container {
	if auth.User != nil {
		var button *widget.Button
		logout := fyne.NewMenuItem("lotgout", func() {
			auth.Logout()
			navigation.PopPage(w)
		})

		menu := fyne.NewMenu("", logout)
		button = widget.NewButton(auth.User.Email, func() {
			popUp := widget.NewPopUpMenu(menu, w.Canvas())
			popUp.ShowAtPosition(button.Position().Add(fyne.NewPos(0, button.Size().Height)))
		})

		return container.New(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			button,
		)
	}
	return container.New(layout.NewHBoxLayout())
}
