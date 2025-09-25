package repositories

import (
	"Callisto/models"
	"Callisto/supabase"
	"fmt"
)

func AddRequest(req models.User) error {
	_, _, err := supabase.Client.From("users").Insert(req, false, "", "*", "").Execute()
	if err != nil {
		return fmt.Errorf("insert error: %s", err.Error())
	}
	return nil
}
