package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 float64 = 3.4
	var x *float64 = &x1
	fmt.Println("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("kind:", v.Kind())
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())

	///////////////////////////////////////////////////////////修改设置值
	fmt.Println("v: ", v.CanSet())
	//v.SetFloat(3.1415)

	v = reflect.ValueOf(&x)
	fmt.Println("v: ", v.CanSet())
	v = v.Elem()
	fmt.Println("v的Elem是： ", v)
	fmt.Println("v: ", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
