package defet

import "fmt"

func RunSample() {
	fmt.Println("Defer package output:")
	defer func() {
		fmt.Println("First line")
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	defer func() {
		fmt.Println("4")
	}()
	fmt.Println("After defer line")
	fmt.Println("")
}
