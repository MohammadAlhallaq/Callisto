package main

import (
	"Callisto/cmd"
	"Callisto/services/auth"
	"Callisto/supabase"
)

func main() {
	supabase.Init()
	auth.FetchLoggedInUser()
	cmd.Excute()
}
