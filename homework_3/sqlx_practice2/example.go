package sqlx_practice2

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type books struct {
	Id     uint
	Title  string
	Author string
	Price  float64
}

func Run() {
	db, err := sqlx.Connect("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println("连接数据失败")
		return
	}
	var sliceBooks []books
	err = db.Select(&sliceBooks, "select * from books where price > ?", 50)
	if err != nil {
		fmt.Printf("查询错误：%v\n", err)
	}
	for _, book := range sliceBooks {
		fmt.Println(book)
	}
}
