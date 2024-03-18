package array

import "fmt"

func Array() {
	// 直接初始化三个元素的数组
	a1 := []int{1, 2, 3}
	fmt.Printf("a1: %v,len:%d,cap:%d \n", a1, len(a1), cap(a1))

	// 少了部分就是默认值 等价于
	a2 := [3]int{1, 2}
	fmt.Printf("a2: %v,len:%d,cap:%d \n", a2, len(a2), cap(a2))

	// 虽然没有显示初始化，但是实际上内存已经分配好，等价于 0 0 0
	a3 := [3]int{}
	fmt.Printf("a3: %v,len:%d,cap:%d \n", a3, len(a3), cap(a3))

	// 数组不支持append 操作
}
