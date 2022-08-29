package recovery

import "fmt"

func crash() {
	panic("out of memory")
}

func RunSample() {
	fmt.Println("Recovery package output:")
	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Println("Recovered from problem: ", recovered)
		}
	}()
	crash()
	fmt.Println("Last line")
	fmt.Println("")
}
