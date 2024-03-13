package main

func Defer() {
	defer func() {
		println("第一个 defer")
	}()

	defer func() {
		println("第二个 defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i) // 闭包写法
	}()
	i = 1
	// TODO: 确定值的原则：
	// TODO: 作为参数传入的：定义defer的时候就确定了
	// TODO: 作为闭包引入的：执行defer对应的方法时才确定
}

func DeferClosureV1() {
	i := 0
	defer func(val int) { // 参数传递
		println(val)
	}(i)
	i = 1
	// TODO: 确定值的原则：
	// TODO: 作为参数传入的：定义defer的时候就确定了
	// TODO: 作为闭包引入的：执行defer对应的方法时才确定
}

// TODO: return的a=0:这样是无法篡改的
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

// TODO: return的a=1:这样是可以篡改的
func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() *MyStruct {
	m := &MyStruct{name: "hello"}
	defer func() {
		m.name = "Jerry"
	}()
	return m
}

type MyStruct struct {
	name string
}

// TODO:自测题，先判断一下 然后在自己运行

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
