package main

import "fmt"

/////////////////这一组函数用于追踪代码
func trace(s string) string {
	fmt.Println("开始执行", s)
	return s
}

func un(s string) {
	fmt.Println("结束执行", s)
}

func main() {
	///////////示例1
	fmt.Println(a())
	fmt.Println(b())

	///////////示例2
	fmt.Println(c())

}

func a() int {
	var i int

	defer func() {
		i++
		fmt.Println("defer2", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1", i)
	}()

	return i

}

func b() (i int) {

	defer func() {
		i++
		fmt.Println("defer2", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1", i)
	}()

	return i

}

func c() (i int, p *int) {
	defer un(trace("b"))
	i++
	// 注意带参，执行i拷贝给形参，然后压入栈
	defer func(i int) {
		fmt.Println("defer3", i, &i)
	}(i)

	//不是匿名函数不会继承值
	defer fmt.Println("defer2", i, &i)

	// 不带参，会在defer出栈时，继承函数内的i，打印出当时的i值
	defer func() {
		fmt.Println("defer1", i, &i)
	}()

	i++

	func() {
		defer un(trace("print1"))
		fmt.Println("print1:", i, &i)
	}()

	fmt.Println("print2:", i, &i)

	return i, &i
}
