package auth

import (
	"Callisto/supabase"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/supabase-community/gotrue-go/types"
)

func SignInWithEmailPassword(email, password string) (types.Session, error) {
	session, err := supabase.Client.SignInWithEmailPassword(email, password)
	if err != nil {
		return types.Session{}, err
	}
	return session, nil
}

func SignUpWithEmail(email, password string) (*types.SignupResponse, error) {
	req := types.SignupRequest{
		Email:    email,
		Password: password,
	}
	session, err := supabase.Client.Auth.Signup(req)
	if err != nil {
		return &types.SignupResponse{}, err
	}
	return session, nil
}

func getSessionFilePath() string {
	dir, _ := os.UserConfigDir() // cross-platform config dir
	return filepath.Join(dir, "callisto-session.json")
}

func SaveSessionData(session types.Session) error {
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

func LoadSession() (*authData, error) {

	file, err := os.Open(getSessionFilePath())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data authData
	err = json.NewDecoder(file).Decode(&data)
	return &data, err
}
