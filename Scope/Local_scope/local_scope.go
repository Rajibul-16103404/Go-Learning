package main

import "fmt"

var (
	a = 20 //Global scope variable in main package
	b = 30 //Global scope variable in main package
)

func main() {
	x := 18 //Global scope variable in main function

	if x >= 18 {
		p := 10 // local scope variable in main function
		fmt.Println("I am Mature boy")
		fmt.Println("I have ", p, "girlfriends!")
	}
}
