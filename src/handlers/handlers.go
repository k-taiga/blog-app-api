package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/k-taiga/blog-app-api/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// w http.ResponseWriterにHello, world!を書き込む
	// http.ResponseWriterはWrite([]byte) (int, error)のinterfaceを満たす
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandle(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// ストリームから直接Decodeする(メモリに一度いれる必要がない)
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle

	// ストリームに直接Encodeした結果をw(ResponseWriter)で返す
	json.NewEncoder(w).Encode(article)
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

	log.Println(page)

	// []models.Articleの型をもつスライスを定義し初期化
	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// func Vars(r *http.Request) map[string]string でmapをパスパラメータのmapを取得する
	// AtoiでStringをintで取得
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	log.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	json.NewEncoder(w).Encode(comment)
}
