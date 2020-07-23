package main

import (
	. "./test" //导入的包仅对Go语言有效
	"fmt"
	mytest "test"
)

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(mytest.Even(1))
}
