package main

import (
	"SoftwareGoDay2/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := server.NewServer()
	r.Def.Run()
}
