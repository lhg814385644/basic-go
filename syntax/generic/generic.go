package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
}

type Number interface {
	int | uint
}

func Sum[T Number](data ...T) T {
	var res T
	for _, v := range data {
		res += v
	}
	return res
}
