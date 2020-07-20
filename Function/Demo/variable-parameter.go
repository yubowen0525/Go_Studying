package main

import "fmt"

func main() {
	// 手动填写参数
	age := ageMinorMax("min", 1, 3, 2, 0)
	fmt.Printf("年龄最小的是%d岁。", age)
	age = ageMinorMax("max", 1, 3, 2, 0)
	fmt.Printf("年龄最大的是%d岁。", age)

	ageArr := []int{7, 9, 3, 5, 1}
	//可变参数的传入需要带...
	f1(ageArr...)
}

func ageMinorMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}

	if m == "max" {
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	} else if m == "min" {
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	} else {
		e := -1
		return e
	}
}

func f1(arr ...int) {
	//可变参数的传入需要带...
	f2(arr...)
	fmt.Println("")
	//数组的传入不需要带...
	f3(arr)
}

func f2(arr ...int) {
	for _, char := range arr {
		fmt.Printf("%d\t", char)
	}
}

func f3(arr []int) {
	for _, char := range arr {
		fmt.Printf("%d\t", char)
	}
}
