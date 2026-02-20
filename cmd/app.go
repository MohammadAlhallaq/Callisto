package cmd

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Execute() {
	a := app.New()
	w := a.NewWindow("Callisto")
	mainPage := views.NewMainView(w)
	entryPage := views.NewEntryView(w)

	if auth.User != nil {
		w.SetContent(mainPage)
		navigation.PushPage(w, entryPage, mainPage)
	} else {
		w.SetContent(entryPage)
		navigation.PushPage(w, entryPage)
	}
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
