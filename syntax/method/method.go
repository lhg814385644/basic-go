package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

}

func ChangeUser() {
	//u1 := User{
	//	Name: "zhangsan",
	//	Age:  18,
	//}
}

func (u User) ChangeName(name string) {
	fmt.Printf("u adderss: %p \n", &u)
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	fmt.Printf("u adderss: %p \n", u)
	u.Age = age
}
