package polymorphism

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Post 模型
type Post struct {
	gorm.Model
	Title string
	Body  string
	// GORM 会自动管理 CreatedAt 和 UpdatedAt
	Comments []Comment `gorm:"polymorphic:Commentable"` // 声明这是一个多态关联，"Commentable" 是多态的类型
}

// Video 模型
type Video struct {
	gorm.Model
	Title    string
	URL      string
	Duration int
	// 同样，Video 也可以有多个评论
	Comments []Comment `gorm:"polymorphic:Commentable"`
}

// Comment 模型 (多态关联的拥有者)
type Comment struct {
	gorm.Model
	Body            string
	CommentableID   uint   // 多态外键的 ID 部分 (会映射到 commentable_id)
	CommentableType string // 多态外键的类型部分 (会映射到 commentable_type)
}

func Run() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db.AutoMigrate(&Post{}, &Video{}, &Comment{})
	//
	//// 创建一个 Post
	//post := Post{Title: "My First Post", Body: "Hello World!"}
	//db.Create(&post)
	//
	//// 创建一个 Video
	//video := Video{Title: "My First Video", URL: "https://example.com/video1", Duration: 300}
	//db.Create(&video) // 假设 video.ID = 1
	//
	//// 创建一个属于 Post 的评论
	//comment1 := Comment{
	//	Body:            "Great post!",
	//	CommentableType: "Post",  // 指定关联类型为 Post
	//	CommentableID:   post.ID, // 指定关联的 Post 的 ID
	//}
	//db.Create(&comment1)
	//
	//// 创建一个属于 Video 的评论
	//comment2 := Comment{
	//	Body:            "Awesome video!",
	//	CommentableType: "Video",  // 指定关联类型为 Video
	//	CommentableID:   video.ID, // 指定关联的 Video 的 ID
	//}
	//db.Create(&comment2)

	// 查询所有评论，并预加载它们所属的资源 (Post 或 Video)
	var comments []Comment
	db.Preload("Commentable").Find(&comments)

	// GORM 会自动根据 CommentableType 判断是加载 Post 还是 Video
	for _, comment := range comments {
		fmt.Printf("Comment: %s\n", comment.Body)
		// 使用类型断言来判断具体类型并访问其字段
		switch comment.CommentableType {
		case "Post":
			fmt.Printf("  Belongs to Post: %s\n", comment.Body)
		case "Video":
			fmt.Printf("  Belongs to Video: %s\n", comment.Body)
		default:
			fmt.Println("  Belongs to unknown type")
		}
	}
}
