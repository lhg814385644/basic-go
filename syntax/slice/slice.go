package main

import "fmt"

func main() {
	//Slice()
	//SubSlice()
	ShareSlice()
}

func Slice() {
	s1 := []int{1, 2, 3, 4} // 直接初始化4个元素的切片
	fmt.Printf("s1:%v,len:%d cap:%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 直接初始化三个元素，容量为4的切片
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 7) // 追加一个元素，没有发生扩容，因为本身的容量为4
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) // 再追加一个原始，发生扩容
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // make只传入一个参数，表示创建一个4个元素的切片
	fmt.Printf("s3:%v,len:%d cap:%d\n", s3, len(s3), cap(s3))

	// 按照下标索引
	fmt.Printf("s3[2]:%d\n", s3[2])
	// TODO:最佳实践，在初始化切片的时候要预估容量
}

func SubSlice() {
	// 数组和切片都可以通过[start:end]的形式来获取子切片
	// arr[start:end] 获取[start,end)之间的元素
	// arr[:end] 获取[0:end)之间的元素
	// arr[start:] 获取[start:len(arr))之间的元素
	// 都是左闭右开
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2))

	s3 := s1[2:] //
	fmt.Printf("s3:%v,len:%d cap:%d\n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4:%v,len:%d cap:%d\n", s4, len(s4), cap(s4))

	// TODO: 内存共享问题
	// 核心：共享数组
	/*
		TODO:
		  子切片和切片究竟会不会互相影响就抓住一点：他们是不是还共享数组？
		  什么意思？
		  就是如果他们结构没有变化，那肯定是共享的 但是结构变化了 就可能不共享了
		  什么情况下结构会发生变化？扩容了
		  所以，切片与子切片，切片作为参数传递到别的方法，结构体里面 任何情况下你要判断是否内存共享 那么就一点 有没有扩容
	*/
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1:%v,len:%d cap:%d\n", s1, len(s1), cap(s1)) // {1,2,3,4} 4,4
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2)) // {3,4} 2,2

	s2[0] = 99
	fmt.Printf("s1:%v,len:%d cap:%d\n", s1, len(s1), cap(s1)) // {1,2,99,4} 4,4  共享底层数组
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2)) // {99,4} 2,2

	// s2 cap=2, 再append一个元素 因此会出现扩容，所以他和s1不共享底层数组了
	s2 = append(s2, 199)
	fmt.Printf("s1:%v,len:%d cap:%d\n", s1, len(s1), cap(s1)) // {1,2,99,4} 4 4
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2)) // {99,4,199} 3,4 扩原来cap的两倍

	s2[1] = 1999                                              // 不会修改s1的值，不再共享底层数组
	fmt.Printf("s1:%v,len:%d cap:%d\n", s1, len(s1), cap(s1)) //
	fmt.Printf("s2:%v,len:%d cap:%d\n", s2, len(s2), cap(s2))
}
