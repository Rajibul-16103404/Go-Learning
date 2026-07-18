package Return_Function

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	a := 70
	b := 30

	sum := add(a, b)
	fmt.Println("Sum:", sum)

	p, q := getNumbers(a, b)

	fmt.Println(p, q)
}
