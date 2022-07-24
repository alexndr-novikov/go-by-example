package values

import "fmt"

func Values() {
	fmt.Println("Values package output:")
	fmt.Println("hello" + " go")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
	fmt.Println("")
}
