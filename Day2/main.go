package main

import (
	"SoftwareGoDay2/server"
)

func main() {
	r := server.NewServer()
	r.Def.Run()
}
