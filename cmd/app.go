package cmd

import (
	"Callisto/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Excute() {
	a := app.New()
	w := a.NewWindow("Callisto")
	content := components.NewHelloPage(w)
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
