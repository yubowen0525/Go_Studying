package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "233"
	fmt.Printf("orig 当前是%T类型，且操作系统是%d位。\n", orig, strconv.IntSize)

	num, err := strconv.Atoi(orig)
	fmt.Printf("orig 当前是%T类型，且数值是%d位。\n", num, num)
	fmt.Println(err)

	num += 5
	newS := strconv.Itoa(num)
	fmt.Printf("%s\n", newS)

}
