package repository

import (
	"fmt"
	"homework_4_blog/internal/model"
	"homework_4_blog/pkg/util"
)

func CreatePost(post *model.Post) error {
	err := util.GetDB().Create(&post).Error
	if err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	} else {
		return nil
	}
}

func SelectPostList(userID uint) (*[]map[string]interface{}, error) {
	var posts []map[string]interface{}
	err := util.GetDB().
		Model(&model.Post{}).
		Where("user_id = ?", userID).
		Select("title", "content").
		Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("查询文章列表失败: %w", err)
	}
	return &posts, nil
}

func SelectPostInfoByTitle(title string) (*map[string]interface{}, error) {
	var post map[string]interface{}
	err := util.GetDB().
		Model(&model.Post{}).
		Where("title = ?", title).
		First(&post).Error
	if err != nil {
		return nil, fmt.Errorf("查询文章失败: %w", err)
	}
	return &post, nil
}

func UpdatePost(post *model.Post) error {
	err := util.GetDB().
		Model(&model.Post{}).
		Where("id = ?", post.ID).
		Update("title", post.Title).
		Update("content", post.Content).
		Error
	if err != nil {
		return fmt.Errorf("更新文章失败: %w", err)
	} else {
		return nil
	}
}

func DeletePost(post *model.Post) error {
	err := util.GetDB().Delete(&model.Post{}, post.ID).Error
	if err != nil {
		return fmt.Errorf("删除文章失败: %w", err)
	}
	return nil
}
