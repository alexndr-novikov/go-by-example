package variadic

import "fmt"

func RunSample() {
	fmt.Println("Variadic package output:")

	sum(1, 10, 100, 1000)
	fmt.Println("")

	nums := []int{1, 100, 10000}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
