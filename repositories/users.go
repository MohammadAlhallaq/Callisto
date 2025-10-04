package repositories

import (
	"Callisto/models"
	"Callisto/supabase"
	"encoding/json"
	"fmt"
)

func AddUser(req models.User) ([]models.User, error) {
	res, _, err := supabase.Client.
		From("users").
		Insert(req, false, "", "representation", "").
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	var users []models.User
	if err := json.Unmarshal(res, &users); err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return users, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var users []models.User

	_, err := supabase.Client.
		From("users").
		Select("id,email,password,created_at", "*", false).
		Eq("email", email).
		Limit(1, "").
		ExecuteTo(&users)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	if len(users) == 0 {
		return models.User{}, fmt.Errorf("user not found")
	}
	return users[0], nil
}
