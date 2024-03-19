package main

import "fmt"

func main() {
	//LoopBug()
	Cal()
}

func forArray() {
	arr := [3]string{"11", "22", "33"}
	for i, s := range arr {
		println(i, s)
	}
	for i := range arr {
		println(i, arr[i])
	}
}

func forMap() {
	m := map[string]int{"a": 1, "b": 2}
	for s, i := range m {
		println(s, i)
	}
}

// LoopBug 特别注意
func LoopBug() {
	users := []User{{Name: "Tom"}, {Name: "Jerry"}}

	m := make(map[string]*User, 2)
	/*
		TODO: 千万不要对迭代参数取地址
		在内存里面，迭代参数都是放在同一个地方的，你循环开始就确定了
		，所以你一旦取地址，那么你拿到的就是这个地址
		所以，右边的map中的键值对的值，最终都是同一个，也就是最后一个。
		TODO:但是在 Go1.22之后，这个问题被修复了
	*/
	for _, user := range users {
		m[user.Name] = &user
	}

	for k, v := range m {
		fmt.Printf("name:%s,user:%v\n", k, v)
	}
}

type User struct {
	Name string
}

func LoopBreak() {
	i := 0
	for {
		if i >= 100 {
			break
		}
		i++
	}
}
