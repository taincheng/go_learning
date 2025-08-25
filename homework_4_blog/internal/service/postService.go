package service

import (
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/repository"
)

func CreatePost(post *model.Post) error {
	return repository.CreatePost(post)
}

func SelectPostList(userID uint) (*[]map[string]interface{}, error) {
	return repository.SelectPostList(userID)
}

func SelectPostInfoByTitle(title string) (*map[string]interface{}, error) {
	return repository.SelectPostInfoByTitle(title)
}

func UpdatePost(post *model.Post) error {
	return repository.UpdatePost(post)
}

func DeletePost(post *model.Post) error {
	return repository.DeletePost(post)
}
