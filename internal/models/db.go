package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	sqlite_path := "./hello.db"
	db, err := sql.Open("sqlite3", sqlite_path)
	if err != nil {
		return err
	}
	DB = db
	err = CreateSchema()
	if err != nil {
		return err
	}
	return nil
}

func CreateSchema() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(500) PRIMARY KEY,
    username VARCHAR(500),
    email VARCHAR(500),
    password VARCHAR(500)
  );`)
	if err != nil {
		return err
	}
	return nil
}
