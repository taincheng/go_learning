package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"size:30;not null;index" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"userId"`
	User    User
}

// Comment
//
//	comments 表：存储文章评论信息，包括
//	id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time
	PostID    uint
	Post      Post
	UserID    uint
	User      User
}

// User users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;index;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
}
