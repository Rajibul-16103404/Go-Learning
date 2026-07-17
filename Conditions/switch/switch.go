package main

import "fmt"

func main() {
	a := 2

	switch {
	case a > 20:
		fmt.Println("a is greater than 20")
	case a < 20:
		fmt.Println("a is less than 20")
	default:
		fmt.Println("a is equal to 20")
	}
}