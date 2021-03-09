package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Type %T\n", myFunc)
	coolFunc(myFunc)
}

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

// coolFunc takes in a function
// as a parameter
func coolFunc(a func()) {
	// it then immediately calls that functino
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}
