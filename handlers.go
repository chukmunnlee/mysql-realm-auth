package main

import (
	"fmt"
	"net/http"
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
)

func index(c *gin.Context) {
	msg := fmt.Sprintf("<h1>The current time is %s</h1>", time.Now().Format(time.RFC850))
	c.Data(http.StatusOK, HTML, []byte(msg))
}

func healthz(authDB AuthDB) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authDB.Ping(); nil != err {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("OK: %s", time.Now().Format(time.RFC850)),
		})
	}
}

func authenticate(authDB AuthDB) func(*gin.Context) {
	return func(c *gin.Context) {
	}
}
