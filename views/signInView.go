package views

import (
	"Callisto/navigation"
	"Callisto/services/auth"
	"fmt"

	"fyne.io/fyne/v2"
)

func NewSignInForm(w fyne.Window, authSvc *auth.AuthService, nav *navigation.Navigator) *fyne.Container {
	return newAuthForm(w, authSvc, nav, "Sign In", func(email, password string) error {
		if err := authSvc.SignInWithEmailPassword(email, password); err != nil {
			return fmt.Errorf("login failed: %v", err)
		}
		return nil
	})
}
