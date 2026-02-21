package views

import (
	"Callisto/navigation"
	"Callisto/services/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewEntryView(w fyne.Window, authSvc *auth.AuthService, nav *navigation.Navigator) *fyne.Container {

	welcomeLabel := widget.NewLabel("Welcome to Callisto")
	welcomeLabel.Alignment = fyne.TextAlignCenter

	legoImage := canvas.NewImageFromFile("./assets/callisto-logo.png")
	legoImage.FillMode = canvas.ImageFillContain
	legoImage.SetMinSize(fyne.NewSize(400, 400))

	loginBtn := widget.NewButton("Login", func() {
		signInPage := NewSignInForm(w, authSvc, nav)
		nav.PushPage(w, signInPage)
		w.SetContent(signInPage)
	})
	loginBtn.Importance = widget.HighImportance

	signupBtn := widget.NewButton("Sign Up", func() {
		signUpPage := NewSignUpForm(w, authSvc, nav)
		nav.PushPage(w, signUpPage)
		w.SetContent(signUpPage)
	})
	signupBtn.Importance = widget.MediumImportance

	guestBtn := widget.NewButton("Continue as Guest", func() {
		mainHeaderTabs := NewMainView(w, authSvc, nav)
		nav.PushPage(w, mainHeaderTabs)
		w.SetContent(mainHeaderTabs)
	})

	buttons := container.NewVBox(
		loginBtn,
		signupBtn,
		guestBtn,
	)

	content := container.New(layout.NewCenterLayout(),
		container.NewPadded(
			container.NewVBox(
				legoImage,
				welcomeLabel,
				widget.NewSeparator(),
				buttons,
			),
		),
	)
	return content
}
