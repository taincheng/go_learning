package sql_practice1

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作

假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：

	编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   uint8
	Grade string
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db.AutoMigrate(&Student{})

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	db.Save(&student)

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var ageGt18 []Student
	db.Where("age > ?", 18).Find(&ageGt18)
	for _, student := range ageGt18 {
		fmt.Printf("Name: %v, Age: %v, Grade: %v\n", student.Name, student.Age, student.Grade)
	}

	//	编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "四年级")

	//	编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	students := []Student{
		{Name: "lishi", Age: 13, Grade: "三年级"},
		{Name: "wangwu", Age: 15, Grade: "五年级"},
	}
	db.Save(&students)
	db.Where("age < ?", 15).Delete(&students)
}
