package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the Application")
}

func inputUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Read user input and store it in the 'name' variable
	return name
}

func getTwoNumbers() (int, int) {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1) // Read user input and store it in the 'num1' variable

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2) // Read user input and store it in the 'num2' variable

	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayResult(name string, num1 int, num2 int, sum int) {
	fmt.Println("Hello,", name)
	fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using the application!")
	fmt.Println("Goodbye!")
}

func main() {
	printWelcomeMessage()

	name := inputUserName()

	num1, num2 := getTwoNumbers()

	sum := add(num1, num2)

	displayResult(name, num1, num2, sum)

	printGoodbyeMessage()
}

// func main() {

// 	// Welcome Message
// 	fmt.Println("Welcome to the Application")

// 	// Get user name as input
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name) // Read user input and store it in the 'name' variable

// 	// input 2 numbers from user
// 	var num1, num2 int
// 	fmt.Print("Enter first number: ")
// 	fmt.Scanln(&num1)
// 	fmt.Print("Enter second number: ")
// 	fmt.Scanln(&num2)

// 	sum := num1 + num2

// 	fmt.Println("Hello, ", name)
// 	fmt.Println("The sum of", num1, "and", num2, "is:", sum)
// 	fmt.Println("Thank you for using the application!")
// 	fmt.Println("Goodbye!")

// }
