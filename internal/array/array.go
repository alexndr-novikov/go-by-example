package array

import "fmt"

func RunSample() {
	fmt.Println("Array package output:")

	fmt.Println("Zero value array example:")
	a := [5]int{}
	fmt.Println("Emp:", a)

	a[4] = 100
	fmt.Println("Updated:", a)
	fmt.Println("Last element:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4}
	fmt.Println("New arr:", b)

	twoDim := [2][3]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("Tow dim array:", twoDim)

	fmt.Println("")
}
