package main

import (
	"Callisto/cmd"
	"Callisto/services/auth"
	"Callisto/supabase"
	"fmt"
)

func main() {
	supabase.Init()
	err := auth.FetchLoggedInUser()
	if err != nil {
		fmt.Println("failed to get user: %w", err)
	}
	cmd.Excute()
}
