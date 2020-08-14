package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir, file := filepath.Split("/home/zuolan/examply")
	fmt.Println(dir, file)
	dir, file = filepath.Split("/home/zuolan/examply/")
	fmt.Println(dir, file)

	path := filepath.Join("yubowen", "", "/home/yu", "bo")
	fmt.Println(path)
}
