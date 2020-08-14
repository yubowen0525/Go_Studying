package main

import (
	"fmt"
	"time"
)

func count_1(ch chan int) {
	time.Sleep(5 * 1e9)
	ch <- 1
}

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int)
	go count_1(ch)
	go func() {
		time.Sleep(3 * 1e9)
		timeout <- true
	}()

	select {
	case count := <-ch:
		fmt.Println(count)
	case <-timeout:
		fmt.Printf("计时器：超时3s")
	}
}
