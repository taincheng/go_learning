package repository

import (
	"fmt"
	"homework_4_blog/internal/model"
	"homework_4_blog/pkg/util"
)

func CreateComment(comment *model.Comment) error {
	err := util.GetDB().Create(&comment).Error
	if err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	} else {
		return nil
	}
}

func GetCommentList(postID uint) (*[]map[string]interface{}, error) {
	var comments []map[string]interface{}
	err := util.GetDB().Model(&model.Comment{}).
		Where("post_id = ?", postID).
		Select("content", "user_id", "post_id", "id", "created_at").
		Find(&comments).
		Error
	if err != nil {
		return nil, fmt.Errorf("查询评论列表失败: %w", err)
	}
	return &comments, nil
}
