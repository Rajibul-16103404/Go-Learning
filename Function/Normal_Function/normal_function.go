package Normal_Function

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}

func main() {
	a := 100
	b := 200
	add(a, b)
	add(50, 70)
}
