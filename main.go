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
		fmt.Printf("failed to get user: %v\n", err)
	}
	cmd.Execute()
}
