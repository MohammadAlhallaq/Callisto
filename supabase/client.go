package supabase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func NewClient() (*supabase.Client, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SERVICE_ROLE")

	return supabase.NewClient(url, key, nil)
}
