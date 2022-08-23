package recursion

import "fmt"

func RunSample() {
	fmt.Println("recursion package output:")
	fmt.Println(factorial(7))

	var closureFib func(n int) int

	closureFib = func(n int) int {
		if n < 2 {
			return n
		}
		return closureFib(n-1) + closureFib(n-2)
	}

	fmt.Println(closureFib(7))
	fmt.Println("")
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
