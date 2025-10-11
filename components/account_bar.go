package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewAccountBar(w fyne.Window) *fyne.Container {

	var accountBar *fyne.Container
	var button *widget.Button

	menuItem1 := fyne.NewMenuItem("Option 1", func() { println("Option 1 clicked") })
	menuItem2 := fyne.NewMenuItem("Option 2", func() { println("Option 2 clicked") })
	menu := fyne.NewMenu("File", menuItem1, menuItem2)

	button = widget.NewButton("Show Menu", func() {
		popUp := widget.NewPopUpMenu(menu, w.Canvas())
		popUp.ShowAtPosition(button.Position().Add(fyne.NewPos(0, button.Size().Height)))
	})
	accountBar = container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(), // push to right
		button,
	)
	return accountBar
}
