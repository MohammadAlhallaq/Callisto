package views

import (
	"Callisto/models"
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/services/validation"
	"fmt"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewSignUpForm(w fyne.Window) *fyne.Container {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	titleLabel := widget.NewLabel("SignUp")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: emailEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			email := strings.TrimSpace(emailEntry.Text)
			password := strings.TrimSpace(passwordEntry.Text)

			// Validation rules
			if email == "" {
				dialog.ShowError(fmt.Errorf("email cannot be empty"), w)
				return
			}
			if !validation.IsValidEmail(email) {
				dialog.ShowError(fmt.Errorf("invalid email format"), w)
				return
			}
			if password == "" {
				dialog.ShowError(fmt.Errorf("password cannot be empty"), w)
				return
			}
			if len(password) < 6 {
				dialog.ShowError(fmt.Errorf("password must be at least 6 characters"), w)
				return
			}
			user := models.User{Email: emailEntry.Text, Password: passwordEntry.Text}
			if err := auth.SignUpWithEmail(user); err != nil {
				dialog.ShowError(fmt.Errorf("signup failed: %v", err), w)
			} else {
				w.SetContent(NewMainView(w))
			}
		},
		OnCancel: func() {
			navigation.PopPage(w)
		},
	}

	formBox := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	return formBox
}
