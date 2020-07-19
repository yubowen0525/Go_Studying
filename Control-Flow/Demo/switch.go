package main

import "fmt"

var x interface{} //空接口

func main() {
	grade := "B"
	marks := 75

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Println("你的成绩为", grade)

	switch {
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 60:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("成绩优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("一般")
	case grade == "D":
		fmt.Println("不及格")
	}

	fmt.Println("你的成绩为", grade)

	/////////////////////////////初始化语句
	x = 1
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("未知")

	}

}
