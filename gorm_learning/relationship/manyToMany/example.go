package manyToMany

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.AutoMigrate(&User{}, &Language{})

	//users := []User{
	//	{
	//		Languages: []Language{
	//			{Name: "en1"},
	//			{Name: "en2"},
	//		},
	//	},
	//	{
	//		Languages: []Language{
	//			{Name: "en3"},
	//			{Name: "en4"},
	//		},
	//	},
	//}
	//
	//db.Debug().Save(&users)

	user := User{Model: gorm.Model{ID: 1}}
	var languages []Language
	db.Debug().Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		fmt.Printf("Language.ID: %v, language.Name: %v\n", language.ID, language.Name)
	}
}
