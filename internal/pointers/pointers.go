package pointers

import "fmt"

func RunSample() {
	fmt.Println("Pointers package output:")
	val := 1
	fmt.Println("initial:", val)
	zeroval(val)
	fmt.Println("zeroval:", val)
	zeroptr(&val)
	fmt.Println("zeroptr:", val)
	fmt.Println("pointer:", &val)
}

func zeroval(intval int) {
	intval = 0
}

func zeroptr(ptrval *int) {
	*ptrval = 0
}
