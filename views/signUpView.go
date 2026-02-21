package views

import (
	"Callisto/models"
	"Callisto/navigation"
	"Callisto/services/auth"
	"fmt"

	"fyne.io/fyne/v2"
)

func NewSignUpForm(w fyne.Window, authSvc *auth.AuthService, nav *navigation.Navigator) *fyne.Container {
	return newAuthForm(w, authSvc, nav, "Sign Up", func(email, password string) error {
		user := models.User{Email: email, Password: password}
		if err := authSvc.SignUpWithEmail(user); err != nil {
			return fmt.Errorf("signup failed: %v", err)
		}
		return nil
	})
}
