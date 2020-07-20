package main

import "fmt"

func pipe(ff func() int) int {
	return ff()
}

//FormatFunc 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

func isEven(v int) bool {
	if v%2 == 0 {
		return true
	}
	return false
}

type boolFunc func(int) bool

/**
f为通用接口
*/
func filter(slice []int, f boolFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	//测试1
	s1 := pipe(func() int {
		return 100
	})

	s2 := format(func(s string, x, y int) string {
		//s是"format"，x,y是需要写入的内容
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20)

	fmt.Println(s1)
	fmt.Println(s2)

	//测试2
	ans := filter([]int{1, 2, 3, 4}, isOdd)
	fmt.Println(ans)

	ans = filter([]int{1, 2, 3, 4}, isEven)
	fmt.Println(ans)

}
