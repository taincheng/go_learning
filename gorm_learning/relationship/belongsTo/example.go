package belongsTo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID string
	Company   Company `gorm:"foreignKey:CompanyID;references:ID"`
}

type Company struct {
	ID   int
	Code string
	Name string
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.AutoMigrate(&User{}, &Company{})

	//users := []User{
	//	{Name: "小明", Company: Company{ID: 1, Code: "111", Name: "xxx公司"}},
	//	{Name: "小李", Company: Company{ID: 1, Code: "111", Name: "xxx公司"}},
	//}

	//db.Save(&users)
	//db.Debug().Create(&users)

	// 场景1：创建用户时更新已存在的公司信息
	//users := []User{
	//	{
	//		Name: "小明",
	//		Company: Company{
	//			ID:   1,
	//			Code: "111-new", // 更新公司代码
	//			Name: "新xxx公司",  // 更新公司名称
	//		},
	//	},
	//	{
	//		Name: "小李",
	//		Company: Company{
	//			ID:   1,
	//			Code: "111-new2", // 另一个用户试图更新同一家公司
	//			Name: "更新后的公司",
	//		},
	//	},
	//}

	// 不使用 FullSaveAssociations - 只创建用户，不更新公司信息
	//db.Debug().Create(&users)

	// 使用 FullSaveAssociations - 会更新关联的公司信息
	// 注意：这可能会导致多次更新同一记录
	//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Create(&users)

	// 场景2：批量更新用户及其关联公司
	//var existingUsers []User
	//db.Preload("Company").Find(&existingUsers)
	//
	//// 修改用户和关联公司信息
	//for i := range existingUsers {
	//	existingUsers[i].Name = existingUsers[i].Name + "(更新)"
	//	existingUsers[i].Company.Name = existingUsers[i].Company.Name + "(更新)"
	//}
	//
	//// 使用 FullSaveAssociations 确保公司信息也被更新
	//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&existingUsers)

}
