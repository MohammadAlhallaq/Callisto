package views

import (
	"Callisto/components"

	"fyne.io/fyne/v2/container"
)

func NewMainView() *container.DocTabs {
	return components.NewMainHeaderTabs()
}
