package main

import "fmt"

func main() {
	////////////创建
	dict := make(map[string]int)
	dict1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	fmt.Println(dict)
	fmt.Println(dict1)
	///////////////查找与遍历
	///方式1
	colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}
	///方式2
	value = colors["Red"]
	if value != "" {
		fmt.Println(value)
	}

	////////////////元素删除
	delete(colors, "Red")
	fmt.Println(colors)
}
