package main

import (
	"fmt"
)

func main() {

	fmt.Println("LOOP1 start")

LOOP1:
	for true {
		x := 1
		switch {
		case x > 0:
			fmt.Println("A")
			//跳出switch，但是还是在for中
			//break
			break LOOP1
		case x == 1:
			fmt.Println("B")

		default:
			fmt.Println("C")

		}
	}

	fmt.Println("LOOP2 start")

LOOP2:
	for x := 0; x <= 5; x++ {
		switch {
		case x > 0:
			fmt.Println("A")
			//跳出switch，但是还是在for中
			//break
			continue LOOP2
		case x == 1:
			fmt.Println("B")

		default:
			fmt.Println("C")
		}

		fmt.Println("X=", x)
	}

}
