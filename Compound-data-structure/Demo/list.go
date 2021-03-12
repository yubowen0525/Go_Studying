package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println("len(a)", l.Len())
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	l.MoveAfter(e4, e1)
	// 遍历链表并打印它包含的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
