package main

import "fmt"

func main() {
	str := "abcd"
	for index, ch := range str {
		fmt.Println(index, ch)
	}

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, m[key])
		fmt.Println(key, value)
	}

}
