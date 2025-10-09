package views

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/services/validation"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewSignInForm(w fyne.Window) *fyne.Container {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	titleLabel := widget.NewLabel("Sign In")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Label to show validation errors
	errorLabel := widget.NewLabel("")
	errorLabel.Wrapping = fyne.TextWrapWord
	errorLabel.Alignment = fyne.TextAlignCenter
	errorLabel.TextStyle = fyne.TextStyle{Bold: true}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: emailEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			errorLabel.SetText("")

			email := strings.TrimSpace(emailEntry.Text)
			password := strings.TrimSpace(passwordEntry.Text)

			// Validation rules
			if email == "" {
				errorLabel.SetText("Email cannot be empty")
				return
			}
			if !validation.IsValidEmail(email) {
				errorLabel.SetText("Invalid email format")
				return
			}
			if password == "" {
				errorLabel.SetText("Password cannot be empty")
				return
			}
			if len(password) < 6 {
				errorLabel.SetText("Password must be at least 6 characters")
				return
			}

			if session, err := auth.SignInWithEmailPassword(email, password); err != nil {
				errorLabel.SetText("Login failed: " + err.Error())
			} else {
				auth.SaveSessionData(session)
				log.Println("User found successfully")
			}
		},
		OnCancel: func() {
			navigation.PopPage(w)
		},
	}

	formBox := container.NewVBox(
		layout.NewSpacer(),
		titleLabel,
		form,
		errorLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	return formBox
}
