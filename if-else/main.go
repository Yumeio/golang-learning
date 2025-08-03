package main

import (
	"fmt"
)

func main() {
	if 10%2 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	if 7%2 == 0 || 7%3 == 0 {
		fmt.Println("7 is even or divisible by 3")
	} else {
		fmt.Println("7 is odd and not divisible by 3")
	}

	if num := 100; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is 10 or greater")
	}
}
