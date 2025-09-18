package cmd

import (
	"Callisto/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Excute() {
	a := app.New()
	w := a.NewWindow("Callisto")

	mainHeaderTabs := components.NewMainHeaderTabs()

	w.SetContent(mainHeaderTabs)
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
