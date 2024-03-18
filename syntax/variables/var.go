package main

func main() {
	var a int = 123
	println(a)
}

func IfElse(age int) string {
	if age >= 18 {
		return "成年了"
	} else {
		return "18岁以下"
	}
}

func IfNewVariable(start, end int) string {
	if distance := start - end; distance > 100 {
		println(distance)
		return "距离太远"
	} else {
		println(distance)
		return "距离较近"
	}

	// 编译错误（超出if范围）
	// println(distance)
}
