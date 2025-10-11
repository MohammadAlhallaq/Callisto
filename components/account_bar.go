package components

import (
	"Callisto/services/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewAccountBar() *fyne.Container {

	var accountBar *fyne.Container
	userBtn := widget.NewButton(auth.User.Email, func() {})

	accountBar = container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(), // push to right
		userBtn,
	)
	return accountBar
}
