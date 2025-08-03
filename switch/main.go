package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Printf("Write %d as: ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	timeNow := time.Now()
	switch {
	case timeNow.Hour() < 12:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good afternoon!")
	}

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoami(true)
	whoami(1)
	whoami("hello")
}
