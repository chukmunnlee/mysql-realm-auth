package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CONTENT_TYPE      = "Content-Type"
	AUTHORIZATION     = "Authorization"
	BEARER_WITH_SPACE = "Bearer "
	FORM_URL_ENCODED  = "application/x-www-form-urlencoded"
	JSON              = "application/json"
	HTML              = "text/html"
	FIELD_USERNAME    = "username"
	FIELD_PASSWORD    = "password"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("The current time is %s", time.Now().Format(time.RFC850)),
	})
}

func healthz(authDB AuthDB) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authDB.Ping(); nil != err {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("OK: %s", time.Now().Format(time.RFC850)),
		})
	}
}

func authenticate(authDB AuthDB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user User

		contentType := strings.ToLower(c.Request.Header.Get(CONTENT_TYPE))

		if strings.HasPrefix(contentType, JSON) {
			user = User{}
			if err := c.BindJSON(&user); nil != err {
				c.JSON(http.StatusNotAcceptable, gin.H{
					"message": fmt.Sprintf("Cannot read credentials: %s", err.Error()),
				})
				return
			}
		} else if strings.HasPrefix(contentType, FORM_URL_ENCODED) {
			user = User{
				Username: c.PostForm(FIELD_USERNAME),
				Password: c.PostForm(FIELD_PASSWORD),
			}
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": fmt.Sprintf("Authenticate only supports %s and %s", JSON, FORM_URL_ENCODED),
			})
			return
		}

		log.Printf("username: %s, password: %s\n", user.Username, user.Password)

		//result, err := authDB.Validate("abc", "xyz")

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
