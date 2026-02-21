package navigation

import (
	"fyne.io/fyne/v2"
)

type Navigator struct {
	pageStack []fyne.CanvasObject
}

func NewNavigator() *Navigator {
	return &Navigator{}
}

func (n *Navigator) PushPage(w fyne.Window, pages ...fyne.CanvasObject) {
	for _, page := range pages {
		n.pageStack = append(n.pageStack, page)
		w.SetContent(page)
	}
}

func (n *Navigator) PopPage(w fyne.Window) {
	if len(n.pageStack) < 2 {
		return
	}
	n.pageStack = n.pageStack[:len(n.pageStack)-1]
	prev := n.pageStack[len(n.pageStack)-1]
	w.SetContent(prev)
}
