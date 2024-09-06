package controllers_test

import (
	"testing"

	"github.com/k-taiga/blog-app-api/controllers"
	"github.com/k-taiga/blog-app-api/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleControllers(ser)

	m.Run()
}
