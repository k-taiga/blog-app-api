package services

import (
	"database/sql"
	"errors"

	"github.com/k-taiga/blog-app-api/apperrors"
	"github.com/k-taiga/blog-app-api/models"
	"github.com/k-taiga/blog-app-api/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to select data")
		return []models.Article{}, err
	}

	if len(articleList) == 0 {
		err := apperrors.NotAvailableData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NotAvailableData.Wrap(err, "failed to select data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to select data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to select data")
		return models.Article{}, err
	}

	// commentListのスライスの中身を展開してarticle.CommentListに追加
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}

		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice num")
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
