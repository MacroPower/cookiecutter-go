package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	sb := strings.Builder{}
	Hello(&sb)
	fmt.Println(sb.String())
}

func Hello(r io.Writer) {
	_, err := r.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
