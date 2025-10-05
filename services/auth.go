package auth

import (
	"Callisto/supabase"

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
