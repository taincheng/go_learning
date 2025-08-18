package base_crud

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"autoIncrement"` // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     time.Time      // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

// CreateTable 创建表
func CreateTable(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(fmt.Println("err"))
	}
}

// InsertData 插入单条数据
func InsertData(db *gorm.DB, user *User) {
	db.Create(user)
}

// InsertBatches 批量写入数据
func InsertBatches(db *gorm.DB, users *[]User) {
	db.CreateInBatches(users, 100)
}

func selectData(db *gorm.DB) {
	user := User{}
	fmt.Println("test First ====")
	result := db.First(&user)
	fmt.Printf("查询到的条数: %d\n", result.RowsAffected)

	fmt.Println("test Take ====")
	result = db.Take(&user)
	fmt.Printf("查询到的条数: %d\n", result.RowsAffected)

	fmt.Println("test Last ====")
	result = db.Last(&user)
	fmt.Printf("查询到的条数: %d\n", result.RowsAffected)

	fmt.Printf("user ID: %d\n", user.ID)
}

func updateData(db *gorm.DB) {
	var user = User{}
	db.First(&user, 1)
	fmt.Printf("name: %v, ActivatedAt: %v \n", user.Name, user.ActivatedAt)

	user.Name = "小王"
	user.ActivatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	db.Save(&user)
	fmt.Printf("name: %v, ActivatedAt: %v \n", user.Name, user.ActivatedAt)

	ctx := context.Background()
	// Update with conditions
	_, err := gorm.G[User](db).Where("id = ?", user.ID).Update(ctx, "name", "小李")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 批量更新
	// Update attributes with `struct`, 只会更新非零值
	gorm.G[User](db).Where("id = ?", 2).Updates(ctx, User{Name: "小丽", Age: 18})

	// Update attributes with `map`
	user.ID = 3
	db.Debug().Model(&user).Updates(map[string]interface{}{"Name": "小坤", "Age": 14})
	// UPDATE `users` SET `age`=14,`name`="小坤",`updated_at`="2025-08-18 11:07:56.955" WHERE `id` = 3

	// Select 更新选定的列, Omit 排除选定的列
	user.ID = 4
	db.Debug().Model(&user).Select("name", "age").Omit("age").Updates(map[string]interface{}{"name": "王哥", "age": 38})
	// UPDATE `users` SET `name`="王哥",`updated_at`="2025-08-18 11:07:56.955" WHERE `id` = 4

}

func testDeleteData(db *gorm.DB) {
	db.Where("age > ?", 200).Delete(&User{})
}

// ConnectSqLite 创建表sqlLite的表连接
func ConnectSqLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("err")
	} else {
		return db
	}
	return nil
}

func TestCreateTable() {
	connect := ConnectSqLite()
	CreateTable(connect)
}

func TestInsertTable() {
	connect := ConnectSqLite()
	// 单条写入
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	InsertData(connect, &user)

	// 批量写入
	users := []User{
		{Name: "Jinzhu1", Age: 10, Birthday: time.Now()},
		{Name: "Jinzhu2", Age: 29, Birthday: time.Now()},
	}
	InsertBatches(connect, &users)
	for _, user := range users {
		fmt.Println(user.ID)
	}
}

func TestSelectTable() {
	connect := ConnectSqLite()
	selectData(connect)
}

func TestUpdateTable() {
	connect := ConnectSqLite()
	updateData(connect)
}

func TestDeleteTable() {
	connect := ConnectSqLite()
	testDeleteData(connect)
}
