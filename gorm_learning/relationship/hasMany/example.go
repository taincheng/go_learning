package hasMany

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	//user := User{
	//	CreditCards: []CreditCard{
	//		{Number: "123"},
	//		{Number: "124"},
	//		{Number: "125"},
	//	},
	//}

	db.AutoMigrate(&User{}, &CreditCard{})

	// 写入数据
	//db.Debug().Save(&user)

	user := User{
		Model: gorm.Model{ID: 1},
	}
	//删除指定用户时，也删除该用户的卡信息, 必须在 User 中指定 ID 主键，才能正确关联删除
	db.Debug().Select("CreditCards").Delete(&user)
}
