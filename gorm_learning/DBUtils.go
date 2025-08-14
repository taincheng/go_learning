package gorm_learning

import (
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

func InsertData(db *gorm.DB, user *User) {
	db.Create(user)
}

func InsertBatches(db *gorm.DB, users *[]User) {
	db.CreateInBatches(users, 100)
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

func TestDB() {
	connect := ConnectSqLite()
	//CreateTable(connect)
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//InsertData(connect, &user)
	users := []User{
		{Name: "Jinzhu1", Age: 10, Birthday: time.Now()},
		{Name: "Jinzhu2", Age: 29, Birthday: time.Now()},
	}
	InsertBatches(connect, &users)
	for _, user := range users {
		fmt.Println(user.ID)
	}
}
