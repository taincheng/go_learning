package hasOne

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 有一张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.AutoMigrate(&User{}, &CreditCard{})

	//user := User{
	//	CreditCard: CreditCard{
	//		Number: "123",
	//	},
	//}
	//
	//db.Save(&user)

	//关联更新
	//user := User{
	//	Model: gorm.Model{ID: 1},
	//	CreditCard: CreditCard{
	//		//Model:  gorm.Model{ID: 1},
	//		Number: "555",
	//	},
	//}
	//db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	// 预加载
	user := User{}
	err := db.Debug().Preload("CreditCard").Find(&user, 1).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(user.ID, user.CreditCard.ID, user.CreditCard.Number)

}
