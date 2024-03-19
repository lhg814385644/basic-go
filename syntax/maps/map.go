package maps

import "fmt"

func InitMap() {
	// 直接初始化
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	val, ok := m1["key1"]
	if ok {
		fmt.Println(val)
	}
	// 预估容量
	// m2 := make(map[string]string, 2)

	// 遍历M1(注意MAP是无序的)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 删除m1中key2
	delete(m1, "key2")
}
