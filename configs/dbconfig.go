package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"samgeeks/constants"
)

func DbConnect(user, password, host, database string) (*sql.DB, error) {
	//func DbConnect(user, password, host, database string) (*sql.DB, error) {

	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening database connect: %w", err)
	}
	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}
	return db, nil
}

func DbConnected() (*sql.DB, error) {
	//func DbConnect(user, password, host, database string) (*sql.DB, error) {

	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s", constants.DBUsername, constants.DBPassword, constants.DBHost, constants.DBTable)
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening database connect: %w", err)
	}
	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}
	return db, nil
}
