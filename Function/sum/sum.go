package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 20

	sum := sum(a, b)
	fmt.Println(sum)
}
