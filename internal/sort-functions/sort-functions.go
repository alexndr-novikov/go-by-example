package sort_functions

import (
	"fmt"
	"sort"
)

type Person struct {
	age uint
}

type byAge []Person

func (persons byAge) Len() int {
	return len(persons)
}

func (persons byAge) Less(i, j int) bool {
	return persons[i].age < persons[j].age
}

func (persons byAge) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func RunSample() {
	fmt.Println("sort functions package output:")
	persons := []Person{{age: 81}, {age: 19}, {age: 22}, {age: 1}}
	sort.Sort(byAge(persons))
	fmt.Println(persons)
	fmt.Println("")
}
