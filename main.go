package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Setup signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-interrupt
		log.Printf("Caught %v. Exiting", sig)
		authDB.Close()
		os.Exit(0)
	}()

	log.Println(fmt.Sprintf("Starting server on port %s", opts.Port))

	if err := r.Run(fmt.Sprintf(":%s", opts.Port)); nil != err {
		log.Panicf("Cannot start server: %s\n", err.Error())
	}
}
