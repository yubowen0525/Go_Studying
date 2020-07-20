package main

import (
	"fmt"
	"os"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return fmt.Sprintf("<%s> <%s> <%s>", e.Op, e.Path, e.Err.Error())
}

func Stat(name string) (fi os.FileInfo, err error) {
	_, err = os.Stat(name)
	if err != nil {
		return nil, &PathError{
			Op:   "stat",
			Path: name,
			Err:  err,
		}
	}

	return os.Stat(name)
}

func main() {
	//第一种方式
	fmt.Println(Stat("a.txt"))

	//第二种强制转换
	_, err := os.Stat("a.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			// 获取PathError类型变量e中的其他信息并处理
			fmt.Println(e.Op, e.Path, e.Err)
		}
	}
}
