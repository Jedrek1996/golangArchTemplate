package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		return nil, err
	}
	return db, nil
}
