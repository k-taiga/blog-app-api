package services

import (
	"github.com/k-taiga/blog-app-api/models"
	"github.com/k-taiga/blog-app-api/repositories"
)

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := ConnectDB()
	if err != nil {
		return []models.Article{}, err
	}

	defer db.Close()

	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func GetArticleService(articleID int) (models.Article, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// commentListのスライスの中身を展開してarticle.CommentListに追加
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
