package components

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainHeaderTabs() *container.AppTabs {

	addTab := container.NewTabItem("+", widget.NewLabel(""))
	HeaderTabs := container.NewAppTabs(
		container.NewTabItem("Body", NewFullBody()),
		addTab,
	)

	HeaderTabs.OnSelected = func(ti *container.TabItem) {
		if ti == addTab {
			newTab := container.NewTabItem("New Tab", NewFullBody())
			items := HeaderTabs.Items
			HeaderTabs.Items = append(items[:len(items)-1], newTab, addTab)
			HeaderTabs.Select(newTab)
		}
	}

	return HeaderTabs
}
