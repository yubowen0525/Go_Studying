package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//r := ring.New(3)
	//n := r.Len()
	//fmt.Println(r.Len())
	//for i := 0; i < n; i++ {
	//	r.Value = i
	//	r = r.Next()
	//}
	//
	//r.Do(func(p interface{}) {
	//	fmt.Println(p.(int))
	//})

	//////////////

	// Create a new ring of size 6
	r2 := ring.New(6)

	// Get the length of the ring
	n2 := r2.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n2; i++ {
		r2.Value = i
		r2 = r2.Next()
	}
	fmt.Println("r2.now:", r2.Value) // 0
	// Unlink three elements from r, starting from r2.Next()
	r2.Unlink(3) // 1 , 2, 3

	// Iterate through the remaining ring and print its contents
	r2.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
