package _func

import "fmt"

func Func() {
	fmt.Println("Func package output:")
	fmt.Println("plus(1,2):", plus(1, 2))
	fmt.Println("plusPlus(1,2,3):", plusPlus(1, 2, 3))
	fmt.Println("plusPlusPlus(1,2,3,'Hello'):", plusPlusPlus(1, 2, 3, "Hello"))
	fmt.Println("")
}

func plus(i int, a int) int {
	return i + a
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusPlusPlus(a, b, c int, str string) string {
	sum := a + b + c
	return fmt.Sprintf("%d, %s", sum, str)
}
