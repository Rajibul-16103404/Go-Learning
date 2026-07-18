package if_else

import "fmt"

func main() {

	//Single condition
	a := 2

	if a > 20 {
		fmt.Println("a is greater than 20")
	} else if a < 20 {
		fmt.Println("a is less than 20")
	} else {
		fmt.Println("a is equal to 20")
	}

	//Double condition
	age := 25
	sex := "emale"

	if age >= 18 && sex == "male" {
		fmt.Println("You are a male and you are eligible to marry")
	} else if age >= 18 && sex != "female" {
		fmt.Println("You are a ? and you are eligible to marry")
	} else {
		fmt.Println("You are not eligible to marry")
	}

	//Boolean condition
	isMarried := false
	if !isMarried {
		fmt.Println("You are not married")
	}
}
