package services

import (
	"github.com/k-taiga/blog-app-api/models"
	"github.com/k-taiga/blog-app-api/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
