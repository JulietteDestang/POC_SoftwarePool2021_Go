package main

import (
	"SoftwareGoDay4/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := server.NewServer()
	r.Def.Run()
}
