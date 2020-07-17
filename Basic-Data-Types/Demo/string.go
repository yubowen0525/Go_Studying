package main

import (
	"fmt"
	_ "fmt"
)

var value1 float64

func main() {

	s := "ab"
	fmt.Println("len=", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	/**
	编译错误，字符串的值不能修改
	*/

	//s := "你好，"
	t := s
	s += "世界。"
	fmt.Println(s)
	//	s[0] = 'L'
	fmt.Println(t)

	str1 := `苟利国家生死以\n`

	str2 := "今天天气\n真好"

	fmt.Println(str1)
	fmt.Println(str2)

	s1 := "我是中国人"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d = %v\n", i, s1[i])
	}
	fmt.Println()
	for _, v := range s1 {
		fmt.Printf("%c", v)
	}
	fmt.Print("\n")

}
