package auth

import (
	"Callisto/models"
	"Callisto/supabase"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/supabase-community/gotrue-go/types"
)

var User *models.User

func SignInWithEmailPassword(email, password string) error {
	session, err := supabase.Client.SignInWithEmailPassword(email, password)
	if err != nil {
		return err
	}

	saveSessionData(session)

	User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func SignUpWithEmail(user models.User) error {
	req := types.SignupRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	session, err := supabase.Client.Auth.Signup(req)
	if err != nil {
		return err
	}

	saveSessionData(session.Session)

	User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func getSessionFilePath() string {
	dir, _ := os.UserConfigDir() // cross-platform config dir
	return filepath.Join(dir, "callisto-session.json")
}

func saveSessionData(session types.Session) error {
	data := authData{
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
	}

	file, err := os.Create(getSessionFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

func loadSession() (*authData, error) {
	file, err := os.Open(getSessionFilePath())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data authData
	err = json.NewDecoder(file).Decode(&data)
	return &data, err
}

func FetchLoggedInUser() error {
	authData, err := loadSession()
	if err != nil {
		return err
	}

	res, err := supabase.Client.Auth.RefreshToken(authData.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	session, err := supabase.Client.Auth.WithToken(res.AccessToken).GetUser()
	supabase.Client.EnableTokenAutoRefresh(res.Session)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func Logout() error {
	authData, err := loadSession()
	if err != nil {
		return err
	}

	res, err := supabase.Client.Auth.RefreshToken(authData.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	err = supabase.Client.Auth.WithToken(res.AccessToken).Logout()
	if err != nil {
		return fmt.Errorf("failed to logout user: %w", err)
	}

	path := getSessionFilePath()
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
