package views

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewSignInForm() *fyne.Container {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: emailEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			log.Println("Email:", emailEntry.Text)
			log.Println("Password:", passwordEntry.Text)
		},
		OnCancel: func() {
			log.Println("Login canceled")
		},
	}

	formBox := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(widget.NewLabel("SignIn")),
		form,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	return formBox
}
