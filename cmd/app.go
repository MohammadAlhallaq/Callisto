package cmd

import (
	"Callisto/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Excute() {
	a := app.New()
	w := a.NewWindow("Callisto")

	welcomeLabel := widget.NewLabel("Welcome to Callisto")
	welcomeLabel.Alignment = fyne.TextAlignCenter

	legoImage := canvas.NewImageFromFile("./assets/callisto-logo.png")
	legoImage.FillMode = canvas.ImageFillContain
	legoImage.SetMinSize(fyne.NewSize(150, 150))

	// Buttons
	loginBtn := widget.NewButton("Login", func() {
	})

	signupBtn := widget.NewButton("Sign Up", func() {
	})

	guestBtn := widget.NewButton("Continue as Guest", func() {
		mainHeaderTabs := components.NewMainHeaderTabs()
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


	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
