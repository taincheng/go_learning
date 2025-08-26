package service

import (
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/repository"
)

func CreateComment(comment *model.Comment) error {
	return repository.CreateComment(comment)
}

func GetCommentList(postID uint) (*[]map[string]interface{}, error) {
	return repository.GetCommentList(postID)
}
