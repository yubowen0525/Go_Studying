package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	a := 20
	c := 200
	c = a
	fmt.Println("赋值操作c= ", c)

	c += a
	fmt.Println("c= ", c)

	c -= a
	fmt.Println("c= ", c)

	c *= a
	fmt.Println("c= ", c)

	c /= a
	fmt.Println("c= ", c)

	c <<= 2
	fmt.Println("c= ", c)

	c >>= 2
	fmt.Println("c= ", c)

	//与
	c &= 2
	fmt.Println("c= ", c)

	// 异或和
	c ^= 2
	fmt.Println("c= ", c)

	//或
	c |= 2
	fmt.Println("c= ", c)
}
