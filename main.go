package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	r.GET("/", index)

	api := r.Group("/api", cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	api.GET("/healthz", healthz(authDB))
	api.POST("/authenticate", authenticate(authDB))

	log.Println(fmt.Sprintf("Starting server on port %s", opts.Port))

	if err := r.Run(fmt.Sprintf(":%s", opts.Port)); nil != err {
		log.Panicf("Cannot start server: %s\n", err.Error())
	}
}
