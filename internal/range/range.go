package _range

import "fmt"

func RunSample() {
	fmt.Println("Range package output:")
	exampleArray := [5]int{1, 2, 3, 4, 5}

	exampleSlice := make([]int, 3)
	exampleSlice = append(exampleSlice, 1)
	exampleSlice = append(exampleSlice, 2)
	exampleSlice = append(exampleSlice, 3)

	exampleMap := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("Iterate array:")
	for key, value := range exampleArray {
		fmt.Printf("%d -> %d\n", key, value)
	}

	fmt.Println("Iterate slice:")
	for key, value := range exampleSlice {
		fmt.Printf("%d -> %d\n", key, value)
	}

	fmt.Println("Iterate map:")
	for key, value := range exampleMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	longString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	fmt.Println("Original string:", longString)
	fmt.Println("Iterate via range:")
	for key, value := range longString {
		fmt.Printf("key: %d , rune: %d, val: %#U\n", key, value, value)
	}

	fmt.Println("")
}
