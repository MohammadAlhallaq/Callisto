package auth

import (
	"Callisto/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type AuthService struct {
	client *supabase.Client
	User   *models.User
}

func NewAuthService(client *supabase.Client) *AuthService {
	return &AuthService{client: client}
}

func (a *AuthService) SignInWithEmailPassword(email, password string) error {
	session, err := a.client.SignInWithEmailPassword(email, password)
	if err != nil {
		return err
	}

	saveSessionData(session)

	a.User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func (a *AuthService) SignUpWithEmail(user models.User) error {
	req := types.SignupRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	session, err := a.client.Auth.Signup(req)
	if err != nil {
		return err
	}

	saveSessionData(session.Session)

	a.User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func getSessionFilePath() string {
	dir, _ := os.UserConfigDir()
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

func (a *AuthService) FetchLoggedInUser() error {
	authData, err := loadSession()
	if err != nil {
		return err
	}

	res, err := a.client.Auth.RefreshToken(authData.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	session, err := a.client.Auth.WithToken(res.AccessToken).GetUser()
	a.client.EnableTokenAutoRefresh(res.Session)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	a.User = &models.User{
		ID:        session.User.ID,
		Email:     session.User.Email,
		CreatedAt: session.User.CreatedAt,
	}
	return nil
}

func (a *AuthService) Logout() error {
	authData, err := loadSession()
	if err != nil {
		return err
	}

	res, err := a.client.Auth.RefreshToken(authData.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	err = a.client.Auth.WithToken(res.AccessToken).Logout()
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
	a.User = nil

	return nil
}
