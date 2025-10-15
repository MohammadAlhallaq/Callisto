package cmd

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Excute() {
	a := app.New()
	w := a.NewWindow("Callisto")
	var content *fyne.Container

	if auth.User != nil {
		content = views.NewMainView(w)
	} else {
		content = views.NewEntryView(w)
	}

	navigation.PushPage(w, content)
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
