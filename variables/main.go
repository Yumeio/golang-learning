package main

import (
	"fmt"
)

func main() {

	var a = "Hello, World!"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d float64 = 3.14
	fmt.Println(d)

	var e bool = true
	fmt.Println(e)

	var f rune = 'A'
	fmt.Println(f)

	var g complex128 = 1 + 2i
	fmt.Println(g)

	var h string = "Go is great!"
	fmt.Println(h)

	v := "Variable without type"
	fmt.Println(v)
}
