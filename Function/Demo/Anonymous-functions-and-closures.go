package main

import "fmt"

func main() {
	//////////匿名函数的调用方法
	func(num int) int {
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		fmt.Println(sum)
		return sum
	}(10)

	/////////闭包
	fmt.Println(Add()(3))
	fmt.Println(Add2(2)(3))

	//////////闭包2
	j := 5
	// a等于是返回的内层函数
	a := func() func() {
		i := 10
		return func() {
			fmt.Println(i, j)
		}
	}()

	//上面的写法，保证了i的安全性，只有在函数内部才能进行修改
	a()
	j = 10
	a()

	/////////////////闭包3
	//2层函数包，解一层包，等于到了内层包
	var f = adder() //相当于 fmt.Println(adder()(1))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	/////////////闭包3
	//闭包会保存局部变量的值
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

}

// Add 无参函数，返回值是一个匿名函数,匿名函数返回一个int类型的值
func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// Add2 有参函数，返回值是一个匿名函数,匿名函数返回一个int类型的值
func Add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func adder() func(int) int {
	//闭包函数保存并积累其中的变量的值
	var x int
	return func(d int) int {
		fmt.Println(x, d)
		x += d
		return x
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
