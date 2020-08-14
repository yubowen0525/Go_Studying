package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	p := &Person{"张三", 28, 0}
	fmt.Println(p)
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("这是")
	buffer.WriteString(this.Name + ",")

	if this.Sex == 0 {
		buffer.WriteString("他")
	} else {
		buffer.WriteString("她")
	}

	buffer.WriteString("今年")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString("岁。")
	return buffer.String()
}
