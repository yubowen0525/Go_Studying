package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}

func Count_ch(ch chan int, i int) {
	fmt.Printf("协程%d创建成功\n", i)
	// 10个协程阻塞于此
	ch <- i
	fmt.Println("Counting")
}

func main() {
	//lock := &sync.Mutex{}
	//for i := 0; i < 10; i++ {
	//	go Count(lock)
	//}
	//for true {
	//	lock.Lock()
	//	if count >= 10 {
	//		break
	//	}
	//	lock.Unlock()
	//	//if c >= 10{
	//	//	break
	//	//}
	//}

	//////////////////////////////////////////////////////////////////
	// 初始化10个chan 数组，但每个数组并没有初始化
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		// 赋值？为什么还要赋值
		chs[i] = make(chan int)
		go Count_ch(chs[i], i)
	}

	fmt.Println("主线程休息5s")
	time.Sleep(5 * 1e9)

	for _, ch := range chs {
		//<- ch
		//但读的时候，按照chs数组进行顺序读取，达到同步的目的
		value := <-ch
		fmt.Println(value)
	}

	fmt.Println("主线程休息1s")
	time.Sleep(1 * 1e9)

}
