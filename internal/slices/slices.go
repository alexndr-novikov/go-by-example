package slices

import "fmt"

func Slices() {
	fmt.Println("Slices package output:")

	fmt.Println("make slice length 3:")
	s := make([]string, 3)
	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(len(s))

	s = append(s, "four")
	s = append(s, "five", "six")
	fmt.Println("Appended:", s)

	c := make([]string, len(s))
	arr := [4]int{1, 2, 3, 4}
	copy(c, s)
	s = append(s, "seven")
	fmt.Println("Copy", c)
	fmt.Println("Original", s)

	fmt.Println("slice 2:5", s[2:5])
	fmt.Println("slice :2", s[:2])
	fmt.Println("slice 2:", s[2:])
	fmt.Println("slice :", arr[:])

	shortSyntaxSlice := []int{1, 2, 3}
	fmt.Println(shortSyntaxSlice)

	twoDimArray := make([][]int, 3)
	fmt.Println(twoDimArray)
	for i := 0; i < 3; i++ {
		twoDimArray[i] = make([]int, i+1)
		for j := 0; j < len(twoDimArray[i]); j++ {
			twoDimArray[i][j] = i + j
		}
	}
	fmt.Println(twoDimArray)
	fmt.Println("")
}
