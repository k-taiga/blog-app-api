package services

import (
	"github.com/k-taiga/blog-app-api/apperrors"
	"github.com/k-taiga/blog-app-api/models"
	"github.com/k-taiga/blog-app-api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
