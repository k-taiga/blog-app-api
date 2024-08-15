package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-taiga/blog-app-api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandle).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodPatch)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server started")
	// ListenAndServeで起動
	// 第二引数は渡さなければデフォルトルータ(defaultServeMux)が使われる
	log.Fatal(http.ListenAndServe(":8080", r))
}
