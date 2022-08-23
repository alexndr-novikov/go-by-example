package func_multiple_return

import "fmt"

func RunSample() {
	fmt.Println("FuncMultipleReturnValues package output:")

	a := 1
	b := 2
	sum, newA, _ := sumStoreOriginal(a, b)
	fmt.Println("Sum:", sum)
	fmt.Println("New A:", newA)
	fmt.Println("Equal?:", a == newA)
	fmt.Println("")
	fmt.Println("")
}

func sumStoreOriginal(a, b int) (int, int, int) {
	return a + b, a, b
}
