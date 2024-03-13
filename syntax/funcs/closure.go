package main

// 闭包
func Closure(name string) func() string {
	//  name 变量
	//  方法本身
	return func() string {
		return "hello" + name
	}
}

func ClosureInvoke() {
	fn := Closure("daMing")
	println(fn())
}
