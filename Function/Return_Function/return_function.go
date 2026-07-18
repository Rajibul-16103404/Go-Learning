package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	a := 75
	b := 35

	sum := add(a, b)
	fmt.Println("Sum:", sum)
}
