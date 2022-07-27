package maps

import "fmt"

func Maps() {
	fmt.Println("Maps package output:")
	mapVar := make(map[string]int)
	mapVar["one"] = 1
	mapVar["two"] = 2
	fmt.Println(mapVar)

	mapVal := mapVar["one"]
	fmt.Println("Got map value:", mapVal)
	fmt.Println("Len:", len(mapVar))
	delete(mapVar, "one")
	fmt.Println(mapVar)

	smth, prs := mapVar["one"]

	fmt.Println("Presence:", smth, prs)

	shortSyntaxMap := map[int]string{1: "One", 2: "Two"}
	fmt.Println(shortSyntaxMap)
	fmt.Println("")
}
