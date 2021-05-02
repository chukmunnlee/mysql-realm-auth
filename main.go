package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

func main() {

	//startedOn := time.Now()

	opts := ParseOptions()
	opts.Validate()

	// Connect to database
	authDB := AuthDatabase(opts.DSN)
	if err := authDB.Open(); nil != err {
		log.Panicf("Error. Cannot connect to database. \t%s\n", err.Error())
	}
	defer authDB.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		msg := fmt.Sprintf("<h1>The current time is %s</h1>", time.Now().Format(time.RFC850))
		c.Data(http.StatusOK, HTML, []byte(msg))
	})

	api := r.Group("/api", cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	api.GET("/healthz", func(c *gin.Context) {
		if err := authDB.Ping(); nil != err {
		}
	})

	log.Println(fmt.Sprintf("Starting server on port %s", opts.Port))

	if err := r.Run(fmt.Sprintf(":%s", opts.Port)); nil != err {
		log.Panicf("Cannot start server: %s\n", err.Error())
	}
}
