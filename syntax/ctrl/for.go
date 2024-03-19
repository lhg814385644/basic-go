package main

import "fmt"

func Loop1() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	for i := 0; i < 10; {
		println(i)
		i++
	}
}

func Loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func Cal() {
	for i := 10; i < 100; i++ {
		for j := 10; j < 100; j++ {
			if i+j == 80 && i-j == 20 {
				fmt.Printf("i=%d,j=%d\n", i, j)
			}
		}
	}
}
