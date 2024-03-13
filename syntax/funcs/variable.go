package main

// 不定参数：

func YourName(name string, alias ...string) {
	if len(alias) > 0 {
		println(alias[0])
	}
}

func VariableInvoke() {
	YourName("邓明")
	YourName("邓明", "daMing", "大明明")
}
