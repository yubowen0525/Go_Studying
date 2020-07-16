package main

import (
	"fmt"
	_ "fmt"
	"os"
	_ "os"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func GetClass() (stuNum int, className, headTeacher string) {
	return 49, "一班", "张山"
}

func main() {
	stuNum, _, _ := GetClass()
	fmt.Println(stuNum)
	fmt.Println(HOME, USER, GOROOT)
}
