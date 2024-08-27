package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/k-taiga/blog-app-api/controllers"
	"github.com/k-taiga/blog-app-api/services"
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

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppControllers(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandle).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server started")
	// ListenAndServeで起動
	// 第二引数は渡さなければデフォルトルータ(defaultServeMux)が使われる
	log.Fatal(http.ListenAndServe(":8080", r))
}
