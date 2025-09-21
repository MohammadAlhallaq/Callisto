package components

import (
	"fyne.io/fyne/v2/container"
)

func NewMainHeaderTabs() *container.DocTabs {
	HeaderTabs := container.NewDocTabs(
		container.NewTabItem("New Request", NewFullBody()),
	)
	HeaderTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("New Request", NewFullBody())
	}
	return HeaderTabs
}
