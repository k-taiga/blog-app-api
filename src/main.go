package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/k-taiga/blog-app-api/api"
)

var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser = os.Getenv("USER_NAME")
	dbPassword = os.Getenv("USER_PASSWORD")
	dbDatabase = os.Getenv("DB")
	dbConn = fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
}

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server started")
	// ListenAndServeで起動
	// 第二引数は渡さなければデフォルトルータ(defaultServeMux)が使われる
	log.Fatal(http.ListenAndServe(":8080", r))
}
