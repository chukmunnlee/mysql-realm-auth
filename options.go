package main

import (
	"flag"
	"log"
)

const (
	DEFAULT_PORT              = "5000"
	DEFAULT_DSN               = "fred:fred@tcp(localhost:3306)/auth"
	DEFAULT_AUTH_TOKEN_HEADER = "X-Auth-Token"
	DEFAULT_ISSUER            = "mysql-realm-auth"
	DEFAULT_AUDIENCE          = "application"
	DEFAULT_QUERY             = "select count(*) as valid from users where username = ? and password = sha1(?)"
)

type Options struct {
	Port        string
	CORS        bool
	Logger      bool
	DSN         string
	TokenHeader string
	SignKey     string
	Issuer      string
	Audience    string
	Query       string
}

func ParseOptions() *Options {

	var port string
	var cors bool
	var logger bool
	var dsn string
	var tokenHeader string
	var signKey string
	var issuer string
	var audience string
	var query string

	flag.StringVar(&port, "port", DEFAULT_PORT, "port number")
	flag.StringVar(&dsn, "dsn", DEFAULT_DSN, "connection string for MySQL")
	flag.StringVar(&tokenHeader, "token", DEFAULT_AUTH_TOKEN_HEADER, "token HTTP header")
	flag.StringVar(&issuer, "issuer", DEFAULT_ISSUER, "token issuer")
	flag.StringVar(&audience, "audience", DEFAULT_AUDIENCE, "token audience")
	flag.StringVar(&signKey, "signKey", "", "token signing key")
	flag.StringVar(&query, "query", DEFAULT_QUERY, "query to validate user")
	flag.BoolVar(&cors, "cors", true, "enable cors")
	flag.BoolVar(&logger, "log", true, "enable logging")

	flag.Parse()

	return &Options{
		Port:        port,
		CORS:        cors,
		Logger:      logger,
		DSN:         dsn,
		TokenHeader: tokenHeader,
		Audience:    audience,
		Issuer:      issuer,
		SignKey:     signKey,
		Query:       query,
	}
}

// business rules goes here
func (o *Options) Validate() {

	if 0 == len(o.SignKey) {
		log.Fatalln("Set token signing key")
	}

	if DEFAULT_DSN == o.DSN {
		log.Println("WARNING: Using default DSN")
	}

	log.Printf("Query: %s\n", o.Query)
}
