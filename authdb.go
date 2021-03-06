package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type AuthDB interface {
	Open() error
	Close() error
	Validate(string, string) (bool, error)
	Ping() error
}

type AuthDBImpl struct {
	DSN   string
	query string
	DB    *sql.DB
}

func AuthDatabase(dsn string, query string) AuthDB {
	return &AuthDBImpl{DSN: dsn, query: query}
}

func (auth *AuthDBImpl) Open() error {
	db, err := sql.Open("mysql", auth.DSN)
	if nil != err {
		return err
	}
	auth.DB = db
	return nil
}

func (auth *AuthDBImpl) Close() error {
	return auth.DB.Close()
}

func (auth *AuthDBImpl) Validate(username string, password string) (bool, error) {
	//result, err := auth.DB.Query(FIND_USER_WITH_PASSWORD, username, password)
	result, err := auth.DB.Query(auth.query, username, password)
	if nil != err {
		return false, err
	}
	defer result.Close()

	if !result.Next() {
		return false, nil
	}

	var valid int64
	if err := result.Scan(&valid); nil != err {
		return false, err
	}

	return (valid > 0), nil
}

func (auth *AuthDBImpl) Ping() error {
	return auth.DB.Ping()
}
