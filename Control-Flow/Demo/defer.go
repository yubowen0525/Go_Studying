package main

import "fmt"

var i int

func print() {
	fmt.Println(i)
}

func print_push(i int) {
	fmt.Println(i)
}

func main() {
	for i = 0; i < 5; i++ {
		//只会在i=5的时候执行5次
		defer print()
	}

	for i = 0; i < 5; i++ {
		//只会在i=5的时候执行5次
		defer print_push(i)
	}

}
