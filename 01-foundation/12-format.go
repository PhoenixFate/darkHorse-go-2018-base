package main

import (
	"fmt"
)

func main() {

	//%T
	a := 10
	b := 3.24
	c := 'c'
	d := "sfs"
	fmt.Printf("a:%T; b:%T; c:%T; d:%T\n", a, b, c, d)
	//%d %f %s %c
	fmt.Printf("a=%d; b=%f; c=%c; d=%s\n", a, b, c, d)

	//%v 自动匹配格式
	fmt.Printf("a=%v; b=%v; c=%v; d=%v\n", a, b, c, d)

	fmt.Println("12")
}
