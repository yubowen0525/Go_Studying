package main

import "fmt"

func main() {
	a := 10
	if a > 10 {
		fmt.Printf("a 大于 10 ")
	} else if a < 10 {
		fmt.Printf("a 小于 10")
	} else {
		fmt.Printf("a 等于 10")
	}

	if b := 10; b == 10 {
		fmt.Printf("b 等于 10")
	}
}
