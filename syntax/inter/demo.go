package inter

import "fmt"

// 定义List interface接口
type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int)
}

// TODO:即便是做业务开发 也应该面向接口编程

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(name string) {
	fmt.Printf("u adderss: %p \n", &u)
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	fmt.Printf("u adderss: %p \n", u)
	u.Age = age
}
