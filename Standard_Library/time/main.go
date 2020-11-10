package main

import (
	"fmt"
	"time"
)

type T1 struct {
	Time time.Time `json:"time"`
}

func getZeroTime() (t time.Time) {
	return t
}

func main() {
	t1, _ := time.Parse(time.RFC3339, "")

	if t1 != getZeroTime() {
		fmt.Println("不相等")
	}

}
