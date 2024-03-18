package main

func Switch(status int) string {
	switch status {
	case 0:
		return "初始化"
	case 1:
		return "运行中"
	default:
		return "其他"
	}
}

/*
TODO: switch 的值必须是可比较的，实践中 不能用于Switch的值，编译器会报错
*/
