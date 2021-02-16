package server

import (
	"SoftwareGoDay2/routes"

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
