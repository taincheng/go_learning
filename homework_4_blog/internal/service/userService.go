package service

import (
	"errors"
	"fmt"
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/repository"
	"homework_4_blog/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	// 加密密码，进行存储
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(user)
}

// SelectUser 查询用户
func SelectUser(user *model.User) (*string, error) {
	storeUser, err := repository.SelectUserByUserName(user)
	if err != nil {
		return nil, fmt.Errorf("用户不存在: %w", err)
	}
	token, err := util.GenerateToken(storeUser.ID, storeUser.Username)
	if err != nil {
		return nil, fmt.Errorf("token 生成失败: %w", err)
	}
	return &token, nil
}
