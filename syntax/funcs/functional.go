package main

func MyFunc3() string {
	println("My FUNC3")
	return "My FUNC3"
}

// 函数式编程
func Func4() {
	myFunc3 := MyFunc3 // 方法赋值给变量
	myFunc3()
}

// 函数式编程:局部方法（可以在方法内部声明一个局部方法，他的作用域就在本方法内）

func Func5() {
	fn := func(name string) string {
		return "hello" + name
	}
	str := fn("abc")
	println(str)
}

func Func6() func(name string) string {
	return func(name string) string {
		return name
	}
}

func FuncInvoke() {
	fn := Func6
	fn()
}

// 闭包：方法+他绑定的上下文
