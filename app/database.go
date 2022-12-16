package app

import (
	"database/sql"
	"restGo/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/go_rest_api")
	helper.PanicError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
