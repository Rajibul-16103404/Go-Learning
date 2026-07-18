package Other_Functions

import "fmt"

func printMessage() {
	fmt.Println("Hello, World!")
}

func inputMessage(message string) {
	fmt.Println("Welcome to Go Programming: ", message)
}

func main() {
	printMessage()
	inputMessage("Rajibul Islam")
}
