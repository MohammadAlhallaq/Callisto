package views

import (
	"Callisto/components"

	"fyne.io/fyne/v2/container"
)

func NewMainView() *container.DocTabs {
	HeaderTabs := container.NewDocTabs(
		container.NewTabItem("New Request", components.NewFullBody()),
	)
	HeaderTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("New Request", components.NewFullBody())
	}
	return HeaderTabs
}
