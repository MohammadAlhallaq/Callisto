package main

import (
	"Callisto/cmd"
	"Callisto/supabase"
)

func main() {
	
	supabase.Init()
	cmd.Excute()
}
