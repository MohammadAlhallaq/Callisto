package repositories

import (
	"Callisto/models"
	"Callisto/supabase"
	"fmt"
)

func AddUser(req models.UserInsert) error {
	_, _, err := supabase.Client.
		From("users").
		Insert(req, false, "", "*", "").
		Execute()

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}
