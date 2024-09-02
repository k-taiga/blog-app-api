package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/k-taiga/blog-app-api/apperrors"
	"github.com/k-taiga/blog-app-api/controllers/services"
	"github.com/k-taiga/blog-app-api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentControllers(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	// modelのjsonとあっていなければBadRequest
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
