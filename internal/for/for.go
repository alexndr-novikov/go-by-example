package _for

import "fmt"

func RunSample() {
	fmt.Println("For package output:")
	fmt.Println("Only condition loop from 1 to 5:")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++

	}
	fmt.Println("Init;Condition;Loop for:")
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Infinite loop break && continue example:")

	counter := 0
	for {
		if counter < 10 {
			counter++
			fmt.Println("Increment counter and continue:", counter)
			continue
		} else {
			fmt.Println("Counter increment finished, break")
			break
		}
	}

	fmt.Println("")
}
