package util

import (
	"homework_4_blog/internal/model"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// 程序执行时，初始化数据库连接
	dsn := "blog.db"
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	err1 := db.AutoMigrate(&model.Post{}, &model.Comment{}, &model.User{})
	if err1 != nil {
		panic("迁移表结构失败")
	}

	// 配置数据库池
	configureConnectPool()
}

func configureConnectPool() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to configure database pool")
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
