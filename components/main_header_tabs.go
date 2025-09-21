package components

import (
	"fyne.io/fyne/v2/container"
)

func NewMainHeaderTabs() *container.DocTabs {

	HeaderTabs := container.NewDocTabs(
		container.NewTabItem("New Request", NewFullBody()),
	)

	HeaderTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("New Tab", NewFullBody())
	}
	// HeaderTabs.OnSelected = func(ti *container.TabItem) {
	// 	if ti == addTab {
	// 		newTab := container.NewTabItem("New Request", NewFullBody())
	// 		items := HeaderTabs.Items
	// 		HeaderTabs.Items = append(items[:len(items)-1], newTab, addTab)
	// 		HeaderTabs.Select(newTab)
	// 	}
	// }

	return HeaderTabs
}
