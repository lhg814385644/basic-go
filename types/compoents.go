package main

import "io"

func main() {
	Component()
}

type NameI interface {
	Name() string
}

// 组合
type Outer struct {
	Inner
}

type Outer1 struct {
	*Inner
}

// 组合 interface
type Outer2 struct {
	io.Closer // 接口
}

type Inner struct {
}

func (i Inner) Name() string {
	println("hello Name")
	return "hell Name"
}

func (i Inner) Hello() {
	println("hello by inner")
}

func Component() {
	var o Outer
	o.Hello()

	// TODO 注意由于Outer1里面组合的是*Inner,因此在调用*Inner内部的方法
	// 时必须先初始化*Inner,否则panic
	var o1 Outer1
	o1.Inner = &Inner{}
	o1.Hello()
}
