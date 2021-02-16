package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "world",
	})
}

func query(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func parameter(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "world",
	})
}

func body(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func header(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "world",
	})
}

func cookie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "world",
	})
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", parameter)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)

}
