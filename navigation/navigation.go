package navigation

import (
	"fyne.io/fyne/v2"
)

var pageStack []fyne.CanvasObject

func PushPage(w fyne.Window, page fyne.CanvasObject) {
	pageStack = append(pageStack, page)
	w.SetContent(page)
}

func PopPage(w fyne.Window) {
	if len(pageStack) < 2 {
		return
	}
	pageStack = pageStack[:len(pageStack)-1]
	prev := pageStack[len(pageStack)-1]
	w.SetContent(prev)
}
