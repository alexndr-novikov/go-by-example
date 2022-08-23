package constant

import (
	"fmt"
	"math"
	"reflect"
)

const const_string string = "constant string"

func RunSample() {
	fmt.Println("Constant package output:")
	fmt.Println("Const string outside function:", const_string)
	const infunc_number = 500000000
	fmt.Println("Check infunc_nuber constant type after declaration:", reflect.TypeOf(infunc_number))
	const calculated_const = 3e20 / infunc_number
	fmt.Println("Check calculated_const constant type after declaration:", reflect.TypeOf(calculated_const))
	fmt.Println("Calculated via division const: ", calculated_const)
	fmt.Println("Calculated via division const cast to int64: ", int64(calculated_const))
	fmt.Println("Calculated sinus  via math package", math.Sin(infunc_number))
	fmt.Println("")
}
