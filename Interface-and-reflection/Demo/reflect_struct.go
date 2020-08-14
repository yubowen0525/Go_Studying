package main

import (
	"fmt"
	"reflect"
)

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotKnownType{"Google", "Go", "Code"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	// 等于
	//typ := value.Type()
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d : %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}

	//调用第一个方法
	results := value.Method(0).Call(nil)
	fmt.Println(results)

	/////////////////////////////////////////////////
	t := T{23, "hello"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("the world")
	fmt.Println(t)
	fmt.Println(s)
}

type T struct {
	A int
	B string
}
