package sqlx_practice1

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type employees struct {
	ID         uint
	Name       string
	Department string
	Salary     uint
}

func Run() {
	db, err := sqlx.Connect("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println("连接数据失败")
		return
	}

	var sliceEmploy []employees
	err = db.Select(&sliceEmploy, "select * from employees where department = ?", "技术部")
	if err != nil {
		fmt.Printf("查询失败：%v\n", err)
	}
	for _, e := range sliceEmploy {
		fmt.Printf("技术部: name: %v, department: %v, salary: %v\n", e.Name, e.Department, e.Salary)
	}

	var maxSalaryEmploy employees
	err = db.Get(&maxSalaryEmploy, "select * from employees where salary = (select max(salary) as salary from employees)")
	if err != nil {
		fmt.Printf("查询最大Salary员工失败失败：%v\n", err)
	} else {
		fmt.Printf("最大Salary员工: name: %v, department: %v, salary: %v\n", maxSalaryEmploy.Name, maxSalaryEmploy.Department, maxSalaryEmploy.Salary)
	}
}
