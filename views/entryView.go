package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewEntryView(w fyne.Window) *fyne.Container {

	welcomeLabel := widget.NewLabel("Welcome to Callisto")
	welcomeLabel.Alignment = fyne.TextAlignCenter

	legoImage := canvas.NewImageFromFile("./assets/callisto-logo.png")
	legoImage.FillMode = canvas.ImageFillContain
	legoImage.SetMinSize(fyne.NewSize(400, 400))

	// Buttons
	loginBtn := widget.NewButton("Login", func() {
		prevPage := NewEntryView(w)
		signInForm := NewSignInForm(w, prevPage)
		w.SetContent(signInForm)
	})

	signupBtn := widget.NewButton("Sign Up", func() {
		prevPage := NewEntryView(w)
		signInForm := NewSignUpForm(w, prevPage)
		w.SetContent(signInForm)
	})

	guestBtn := widget.NewButton("Continue as Guest", func() {
		mainHeaderTabs := NewMainView()
		w.SetContent(mainHeaderTabs)
	})

	buttons := container.NewVBox(
		loginBtn,
		signupBtn,
		guestBtn,
	)

	content := container.New(layout.NewCenterLayout(),
		container.NewVBox(
			legoImage,
			welcomeLabel,
			layout.NewSpacer(),
			buttons,
		),
	)
	return content

}
