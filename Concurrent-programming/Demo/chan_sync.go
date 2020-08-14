package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(5 * 1e9)
}

func sendData(ch chan string) {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	close(ch)
}

func getData(ch chan string) {
	//for true {
	//	input,open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Println(input)
	//
	//}

	// 自动检测通道的关闭
	for input := range ch {
		fmt.Println(input)
	}
}
