package main

/*TODO:
核心原则
*/

func main() {
	Invoke()
}

func Func0(name string) string {
	return "hello" + name
}

func Func1(a, b int, name string) (string, error) {
	return "hello" + name, nil
}

func Func2(a, b int) (result string, err error) {
	result = "he"
	return
}

func Invoke() {
	println(Func0("大名"))
}

// Recursive 递归:
func Recursive() {
	Recursive()
}
