package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for true {
		select {
		case ch <- 0:
		case ch <- 1:
		}

		i := <-ch
		fmt.Println("接收到值", i)
	}
}
