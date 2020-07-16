package main

import (
	"fmt"
	_ "fmt"
)

const (
	a = iota // a == 0
	b        // b == 1 隐式使用iota关键字
	c
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             // g == 4
	h       = "h"              // h == "h" ， 单独赋值，iota依旧递增为5
	i       = "i"              // i == "h"，默认使用上面的赋值，iota递增为6
	j       = iota
)

const z = iota //此时z 为 0

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, z)
}
