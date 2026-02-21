package views

import (
	"Callisto/components"
	"Callisto/navigation"
	"Callisto/services/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewMainView(w fyne.Window, authSvc *auth.AuthService, nav *navigation.Navigator) *fyne.Container {

	// Top account bar
	accountBar := components.NewAccountBar(w, authSvc, nav, NewSignInForm(w, authSvc, nav), NewSignUpForm(w, authSvc, nav))

	// Left-side request history
	var contentSplit fyne.CanvasObject
	requestHistoryList := components.NewRequestHistoryList()

	headerTabs := container.NewDocTabs(
		container.NewTabItem("New Request", components.NewFullBody(w)),
	)
	headerTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("New Request", components.NewFullBody(w))
	}

	if requestHistoryList != nil {
		contentSplit = container.NewHSplit(requestHistoryList, headerTabs)
		contentSplit.(*container.Split).SetOffset(0.25) // FYI: Split offset is set this way in Fyne
	} else {
		contentSplit = headerTabs
	}

	content := container.New(
		layout.NewBorderLayout(accountBar, nil, nil, nil),
		accountBar,
		contentSplit,
	)
	return content
}
