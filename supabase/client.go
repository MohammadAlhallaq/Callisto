package supabase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	var err error
	Client, err = supabase.NewClient(url, key, nil)
	if err != nil {
		fmt.Println("Supabase init error:", err)
	}
}
