package server

import (
	"SoftwareGoDay4/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Def *gin.Engine
}

func NewServer() Server {
	serv := Server{Def: gin.Default()}
	routes.ApplyRoutes(serv.Def)
	return serv
}
