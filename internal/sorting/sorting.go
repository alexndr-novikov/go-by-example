package sorting

import (
	"fmt"
	"sort"
)

func RunSample() {
	fmt.Println("Sorting package output:")

	strings := []string{"za", "zz", "a", "c", "z"}
	sort.Strings(strings)
	fmt.Println(strings)

	ints := []int{1, 5, 8, 1, 2}
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints))

	fmt.Println("")
}
