package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人，下次某个时间继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	time.Sleep(5 * 1e9)

}
