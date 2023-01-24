package app

import (
	"database/sql"
	"fmt"
	"time"
)

func NewDB() *sql.DB {
	host := "127.0.0.1"
	port := "5432"
	user := "admin"
	password := "admin"
	dbname := "belajar_golang_restful_api"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
