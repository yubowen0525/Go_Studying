package main

import (
	"fmt"
	"strings"
)

func foo(slice []int) []int {
	//...
	for i, _ := range slice {
		slice[i] = i
	}

	slice = append(slice, 1, 2, 3, 4, 5)
	return slice
}

func main() {
	////////////////////////////切片的使用
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(slice), slice)
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(cap(newSlice), newSlice)
	fmt.Println(cap(slice), slice)
	////////////////////////////////////////////
	fmt.Println()
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println(cap(slice), slice)
	newSlice = slice[1:3]
	slice = append(slice, 60)
	fmt.Println(cap(newSlice), newSlice)
	fmt.Println(cap(slice), slice)
	///////////////////////////////////////////
	fmt.Println()
	slice = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(slice), slice)
	newSlice2 := slice[0:1:3]
	fmt.Println(cap(newSlice2), newSlice2)
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(cap(slice), slice)

	////////////////////切片的传递
	//fmt.Println()
	//slice3 := make([]int, 1e6)
	//slice4 := foo(slice3)
	//fmt.Println(cap(slice4))
	/////////////////////////////

	/////////////////////////////////////
	var s string
	slices := []string{"10", "20", "10", "20", "10"}
	for _, value := range slices {
		if !strings.Contains(s, fmt.Sprintf("[%s %s]", value, value)) {
			s += fmt.Sprintf("[%s %s]", value, value)
		}
	}
	fmt.Println("s=", s)
	///////////////////////////////////
	fmt.Println()
	var n = 100
	var sum float64

	for i := 1; i <= n; i++ {
		x1 := 1 / float64(10*i-5)
		x2 := 1 / float64(10*i)
		sum += x1
		sum -= x2
	}

	fmt.Printf("%.2f", sum)
}
