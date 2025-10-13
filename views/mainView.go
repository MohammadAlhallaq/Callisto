package views

import (
	"Callisto/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewMainView(w fyne.Window) *fyne.Container {

	accountBar := components.NewAccountBar(w)

	
	HeaderTabs := container.NewDocTabs(
		container.NewTabItem("New Request", components.NewFullBody()),
	)
	HeaderTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("New Request", components.NewFullBody())
	}

	content := container.New(
		layout.NewBorderLayout(accountBar, nil, nil, nil),
		accountBar,
		HeaderTabs,
	)
	return content
}
