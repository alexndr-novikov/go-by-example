package conditions

import "fmt"

func Conditions() {
	if 8%2 == 0 {
		fmt.Println("8 is even, if-else condition example")
	} else {
		fmt.Println("8 is odd, if-else condition example")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2, if without else example")
	}

	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
