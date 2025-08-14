package main

import "fmt"

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  uint8
}

type Employee struct {
	Person
	EmployeeID uint32
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %v, Age: %d, EmployeeID: %d", e.Name, e.Age, e.EmployeeID)
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) Area() {
	area := r.width * r.length
	fmt.Printf("Rectangle area is : %v \n", area)
}

func (r Rectangle) Perimeter() {
	perimeter := (r.length + r.width) * 2
	fmt.Printf("Rectangle perimeter is : %v \n", perimeter)
}

type Circle struct {
	radius float32
}

func (c Circle) Area() {
	area := 3.14 * c.radius * c.radius
	fmt.Printf("Circle area is : %v \n", area)
}

func (c Circle) Perimeter() {
	perimeter := 2 * 3.14 * c.radius
	fmt.Printf("Circle perimeter is : %v \n", perimeter)
}

func main() {
	rectangle := Rectangle{2.3, 3.4}
	circle := Circle{2}

	rectangle.Area()
	rectangle.Perimeter()

	circle.Area()
	circle.Perimeter()

	fmt.Println("==============")

	// Go会自动将employee的地址传给方法（等价于(&employee).PrintInfo()），不会报错
	employee := Employee{
		Person: Person{
			Name: "小明",
			Age:  18,
		},
		EmployeeID: 123456,
	}

	employee.PrintInfo()
}
