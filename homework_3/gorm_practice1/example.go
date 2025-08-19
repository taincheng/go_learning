package gorm_practice1

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

type User struct {
	gorm.Model
	UserName   string
	PostNumber uint
	Posts      []Post
}

type Post struct {
	gorm.Model
	Title         string
	CommentStatus string
	UserID        uint
	Comments      []Comment

	User User `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Context string
	PostID  uint

	Post Post `gorm:"foreignKey:PostID"`
}

// BeforeCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_number", gorm.Expr("post_number + ?", 1)).Error
	if err != nil {
		fmt.Printf("更新用户文章数量失败: %v\n", err)
		return err
	} else {
		return nil
	}
}

// AfterDelete 评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		fmt.Printf("计算文章评论条数出错: %v\n", err)
		return err
	}
	if count == 0 {
		err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
		if err != nil {
			fmt.Printf("更新无状态出错: %v\n", err)
			return err
		}
	}
	return nil
}

func Run() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接失败: %v", err)
		return
	}
	err = db.AutoMigrate(&Post{}, &User{}, &Comment{})
	if err != nil {
		fmt.Printf("创建表失败: %v", err)
	}

	users := []User{
		{
			UserName: "小李",
			Posts: []Post{
				{
					Title: "小李发的文章 A",
					Comments: []Comment{
						{Context: "小李的文章 A 写的不错"},
						{Context: "小李的文章 A 写的还可以，有点进步"},
					},
				},
			},
		},
		{
			UserName: "小王",
			Posts: []Post{
				{
					Title: "小王发的文章 A",
					Comments: []Comment{
						{Context: "小王的文章 A 写的不错"},
					},
				},
				{
					Title: "小王发的文章 B",
					Comments: []Comment{
						{Context: "小王的文章 B 确实好"},
						{Context: "小王的文章 B 还行哈，比我差点"},
						{Context: "小王的文章 B 一般吧，没小李的强"},
					},
				},
			},
		},
	}
	err = db.Save(&users).Error
	if err != nil {
		fmt.Printf("写入数据失败：%v\n", err)
		return
	}

	// 查询某个用户发布的所有文章及其对应的评论信息
	var postInfo []Post
	err = db.
		Preload("Comments").
		Where("user_id = ?", 1).
		Find(&postInfo).
		Error
	if err != nil {
		fmt.Println(err)
	} else {
		for _, posts := range postInfo {
			for _, comments := range posts.Comments {
				fmt.Printf("Title: %v, Comments: %v\n", posts.Title, comments.Context)
			}
		}
	}

	// 查询评论数量最多的文章信息
	var maxCommentPostInfo Post
	err = db.
		Joins("left join comments on posts.id = comments.post_id").
		Select("posts.id, posts.title, count(comments.id) as count").
		Group("posts.id").
		Order("count desc").
		First(&maxCommentPostInfo).
		Error
	if err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("评论最多的文章ID: %v, 标题: %v\n", maxCommentPostInfo.ID, maxCommentPostInfo.Title)
	}

	// 测试删除评论
	deleteComment := Comment{
		PostID: 2,
	}
	err = db.Debug().Where("post_id = ?", deleteComment.PostID).Delete(&deleteComment).Error
	if err != nil {
		fmt.Println(err)
	}
}
