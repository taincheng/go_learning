package repository

import (
	"fmt"
	"homework_4_blog/internal/model"
	"homework_4_blog/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser 往数据库中创建用户数据
func CreateUser(user *model.User) error {
	err := util.GetDB().Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// SelectUserByUserName 根据用户名查询用户数据
func SelectUserByUserName(user *model.User) (*model.User, error) {
	var storeUser model.User
	err := util.GetDB().Where("username = ?", user.Username).First(&storeUser).Error
	if err != nil {
		return nil, fmt.Errorf("用户不存在: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("密码错误: %w", err)
	}
	return &storeUser, nil
}
