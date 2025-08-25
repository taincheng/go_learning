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
