package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func ConnectDB() (*sql.DB, error) {
	dbUser = os.Getenv("USER_NAME")
	dbPassword = os.Getenv("USER_PASSWORD")
	dbDatabase = os.Getenv("DB")
	dbConn = fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	log.Printf("DB User: %s, DB Password: %s, DB Name: %s\n, DB Conn: %s\n", dbUser, dbPassword, dbDatabase, dbConn)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
