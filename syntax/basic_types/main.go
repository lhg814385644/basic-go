package main

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	var a int = 456
	var b int = 1
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
}

func ExtremeNum() {
	println(math.MinInt64)
	println("float64 最小正数", math.SmallestNonzeroFloat64)
	println("float32 最小正数", math.SmallestNonzeroFloat32)
}

func String() {
	// hello said "hello,go"
	println("hello said \"hello,go\"")
	println(`hello go
换行了`)

	println("hello" + strconv.Itoa(123))
	println(len("hello")) // TODO: len拿到的是字节长度，不是字符长度 特别注意
	println(len("hello你好"))
	println(utf8.RuneCountInString("hello你好")) // TODO:获取字符长度

	// TODO:strings包相关字符串的操作都在里面
}

func Byte() {
	var a byte = 12
	print(a)
	var str string = "hello"
	var bs []byte = []byte(str)
	println(bs)
}

func Bool() {
	var a bool = true
	var b bool = false
	println(a && b)
	println(a || b)
	println(!a)

	// TODO:注意
	// !(a&&b) 等价于 !a||!b
	// !(a||b) 等价于 !a&&!b
}
