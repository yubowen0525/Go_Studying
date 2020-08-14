package main

import (
	"fmt"
	"math"
)

type shaper interface {
	area() float32
}

type square struct {
	side float32
}

func (sq *square) area() float32 {
	return sq.side * sq.side
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	var areaIntf shaper
	sq1 := new(square)
	sq1.side = 5
	sq2 := circle{2}
	areaIntf = sq1

	if t, ok := areaIntf.(*square); ok {
		fmt.Println(t)
	}

	areaIntf = sq2
	if u, ok := areaIntf.(circle); ok {
		fmt.Println(u)
	}

}
