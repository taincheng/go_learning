package service

import (
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/repository"
)

func CreatePost(post *model.Post) error {
	return repository.CreatePost(post)
}
