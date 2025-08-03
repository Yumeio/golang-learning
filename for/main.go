package main

import (
	"fmt"
)

func main() {
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++ // i = i + 1
	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	for k := range 3 {
		fmt.Println(k)
	}

	for {
		fmt.Println("Infinite loop")
		break // To prevent an actual infinite loop, we break out of it
	}

	for n := range 10 {
		if n%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(n) // Print only odd numbers
	}
}
