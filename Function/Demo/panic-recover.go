package main

import (
	"fmt"
	"time"
)

func main() {
	throwPanci(genErr)
	fmt.Println("正常结束")
}

func genErr() {
	fmt.Println(time.Now(), "正常的语句")

	defer func() {
		fmt.Println(time.Now(), "defer正常的语句")
		panic("第二个错误")
	}()

	panic("第一个错误")
}

func throwPanci(f func()) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(time.Now(), "捕获到异常：", r)
		}
	}()

	f()
	return
}
