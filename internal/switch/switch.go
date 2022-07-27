package _switch

import (
	"fmt"
	"time"
)

func Switch() {
	fmt.Println("Switch package output:")
	i := 1
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("One")
	}

	fmt.Println("Should I deploy today?")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Nope")
	default:
		fmt.Println("Why not?")
	}

	fmt.Println("Emulate if-else. Coffee?")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Yep")
	default:
		fmt.Println("Nope")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's int")
		case bool:
			fmt.Println("It's boolean")
		default:
			fmt.Println("No Idea what is", t)
		}
	}
	whatAmI(true)
	whatAmI(nil)
	whatAmI(1)
	fmt.Println("")
}
