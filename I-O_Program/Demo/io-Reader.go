package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 注意reader可能是file,os.stdin,字符串等
func ReadForm(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func main() {
	data, err := ReadForm(os.Stdin, 11)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	fmt.Printf("%s\n", data)

	data, err = ReadForm(strings.NewReader("from string"), 11)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	fmt.Println(data)

}
