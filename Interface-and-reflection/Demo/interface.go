package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func showArea(area Shaper) {
	fmt.Println(area.Area())
}

func main() {
	//////////////////////////////////////////////////////////
	sq1 := new(Square)
	sq1.side = 5
	areaIntf := Shaper(sq1)
	//areaIntf := sq1
	fmt.Println(areaIntf.Area())

	////////////////////////////////////////////////////////////
	fmt.Println()
	r := Rectangle{5, 3}

	q := &Square{3}

	//shapers := []Shaper{Shaper(r),Shaper(q)}
	shapers := []Shaper{r, q}

	for n, _ := range shapers {
		fmt.Println(shapers[n])
		fmt.Println(shapers[n].Area())
	}
	/////////////////////////////////////////////////////////////
	fmt.Println()
	var a Shaper = Rectangle{5, 3}
	fmt.Println(a)
	showArea(a)

	a = &Square{3}
	fmt.Println(q)
	showArea(a)

	////////////////////////////////////////////////////////////////
	fmt.Println()
	var areaIntff Shaper
	areaIntff = &Circle{2}
	switch t := areaIntff.(type) {
	case *Circle:
		fmt.Println(t, t.Area())
	case *Square:
		fmt.Println(t, t.Area())
	case *Rectangle:
		fmt.Println(t, t.Area())
	}

	/////////////////////////////////////////////////////////////////
	fmt.Println()
	clasifier(1, "string", true, map[string]int{"stirng": 1})
}

func clasifier(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case int:
			fmt.Printf("参数 #%d 类型是 int\n", i)
		case bool:
			fmt.Printf("参数 #%d 类型是 bool\n", i)
		case string:
			fmt.Printf("参数 #%d 类型是 string\n", i)
		default:
			fmt.Printf("参数 #%d 类型未知\n", i)
		}
	}
}
