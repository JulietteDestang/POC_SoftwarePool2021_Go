package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	env := os.Getenv("HELLO_MESSAGE")
	if env == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", env)
	}
}

func query(c *gin.Context) {
	lastname := c.Query("message")
	if lastname == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", lastname)
	}
}

func parameter(c *gin.Context) {
	name := c.Param("message")
	c.String(http.StatusOK, "%s", name)
}

func body(c *gin.Context) {
	var top string
	temp := c.ShouldBind(&top)
	if temp != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", top)
	}
}

func header(c *gin.Context) {
	name := c.GetHeader("message")
	if name == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", name)
	}
}

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func cookie(c *gin.Context) {
	cooki, err := c.Cookie("message")
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", cooki)
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/health", health)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", parameter)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)

}
