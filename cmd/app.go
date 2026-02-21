package cmd

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Execute(authSvc *auth.AuthService, nav *navigation.Navigator) {
	a := app.New()
	w := a.NewWindow("Callisto")
	mainPage := views.NewMainView(w, authSvc, nav)
	entryPage := views.NewEntryView(w, authSvc, nav)

	if authSvc.User != nil {
		w.SetContent(mainPage)
		nav.PushPage(w, entryPage, mainPage)
	} else {
		w.SetContent(entryPage)
		nav.PushPage(w, entryPage)
	}
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
