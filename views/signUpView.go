package views

import (
	"Callisto/models"
	"Callisto/navigation"
	"Callisto/repositories"
	"Callisto/services/auth"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewSignUpForm(w fyne.Window) *fyne.Container {
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
			hashedPassword, _ := auth.HashPassword(passwordEntry.Text)
			user := models.User{Email: emailEntry.Text, Password: hashedPassword}
			if err := repositories.AddUser(user); err != nil {
				log.Println("Could not create user:", err)
			} else {
				log.Println("User created successfully")
			}
		},
		OnCancel: func() {
			navigation.PopPage(w)
		},
	}

	formBox := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(widget.NewLabel("SignUp")),
		form,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	return formBox
}
