package main

import (
	"Callisto/cmd"
	"Callisto/navigation"
	"Callisto/services/auth"
	"Callisto/supabase"
	"fmt"
	"log"
)

func main() {
	client, err := supabase.NewClient()
	if err != nil {
		log.Fatal("failed to init supabase:", err)
	}

	authSvc := auth.NewAuthService(client)
	nav := navigation.NewNavigator()

	if err := authSvc.FetchLoggedInUser(); err != nil {
		fmt.Printf("failed to get user: %v\n", err)
	}

	cmd.Execute(authSvc, nav)
}
