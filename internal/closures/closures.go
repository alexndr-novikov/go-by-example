package closures

import "fmt"

func Closures() {
	fmt.Println("Closures package output:")
	seq := initSequence()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	seq = initSequence()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println("")
}

func initSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
