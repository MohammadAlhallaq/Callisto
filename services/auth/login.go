package auth

import (
	"Callisto/models"
	"Callisto/repositories"
)

func Login(email, password string) (models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	// Verify password
	err = CheckPassword(user.Password, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
