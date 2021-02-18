package routes

import (
	"SoftwareGoDay4/Users"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Queries struct {
	Key   string
	Value []string
}

func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

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

func all(c *gin.Context) {
	name := c.Request.URL.Query()
	i := 0
	queries := make([]Queries, len(name))
	for key, value := range name {
		queries[i].Key = key
		queries[i].Value = value
		i++
	}
	c.JSON(http.StatusOK, queries)
}

func signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {
		for i := range Users.UserDB {
			if email == Users.UserDB[i].Email {
				c.String(http.StatusBadRequest, "Already used Connard")
				return
			}
		}
		Users.UserDB = append(Users.UserDB, Users.User{Email: email, Password: password})
		c.String(http.StatusOK, "User successfully created !")
	}
}

func signin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	for i := range Users.UserDB {
		if (email == Users.UserDB[i].Email) && (password == Users.UserDB[i].Password) {
			email = Encrypt(email, Users.APISECRET)
			c.SetCookie("cookie", email, 24, "/", "localhost", true, false)
			c.String(http.StatusOK, "User successfully logged in !")
			return
		}
	}
	c.String(http.StatusOK, "you SUCK!")
}

func session(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	for i := range Users.UserDB {
		if (email == Users.UserDB[i].Email) && (password == Users.UserDB[i].Password) {
			_, err := c.Cookie(Users.UserDB[i].Email)
			if err == nil {
				c.String(http.StatusOK, "has a cookie")
				return
			} else {
				c.String(http.StatusOK, "no cookie")
				return
			}
		}
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-all-my-queries", all)
	r.GET("/health", health)
	r.POST("/signup-session", signup)
	r.GET("/me-session", session)
	r.POST("/signin-session", signin)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", parameter)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)

}
