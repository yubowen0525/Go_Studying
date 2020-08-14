package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始 main")
	go longWait()
	go shortWait()
	time.Sleep(10 * 1e9) // ns 级别
	fmt.Println("结束 main")
}

func longWait() {
	fmt.Println("开始 longWait")
	time.Sleep(5 * 1e9) // ns 级别
	fmt.Println("结束 longWait")
}

func shortWait() {
	fmt.Println("开始 shortWait")
	time.Sleep(5 * 1e9) // ns 级别
	fmt.Println("结束 shortWait")
}
