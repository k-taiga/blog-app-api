package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// w http.ResponseWriterにHello, world!を書き込む
	// http.ResponseWriterはWrite([]byte) (int, error)のinterfaceを満たす
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int

	// queryMap["page"]があればokにはtrue && pが0より上
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		// デフォルトのpage
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// func Vars(r *http.Request) map[string]string でmapをパスパラメータのmapを取得する
	// AtoiでStringをintで取得
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
