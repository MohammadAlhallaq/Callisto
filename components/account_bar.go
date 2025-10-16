package components

import (
	"Callisto/navigation"
	"Callisto/services/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewAccountBar(w fyne.Window, singinView, signupView *fyne.Container) *fyne.Container {

	var button *widget.Button

	if auth.User != nil {
		logout := fyne.NewMenuItem("Lotgout", func() {
			auth.Logout()
			navigation.PopPage(w)
		})

		menu := fyne.NewMenu("", logout)
		button = widget.NewButton(auth.User.Email, func() {
			popUp := widget.NewPopUpMenu(menu, w.Canvas())
			popUp.ShowAtPosition(button.Position().Add(fyne.NewPos(0, button.Size().Height)))
		})

	} else {
		signin := fyne.NewMenuItem("Signin", func() {
			w.SetContent(singinView)
		})
		signup := fyne.NewMenuItem("Signup", func() {
			w.SetContent(signupView)
		})

		menu := fyne.NewMenu("", signin, signup)
		button = widget.NewButton("Options", func() {
			popUp := widget.NewPopUpMenu(menu, w.Canvas())
			popUp.ShowAtPosition(button.Position().Add(fyne.NewPos(0, button.Size().Height)))
		})
	}

	return container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button,
	)
}
