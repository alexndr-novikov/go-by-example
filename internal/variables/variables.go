package variables

import "fmt"

func Variables() {
	var stringVar = "initial"
	fmt.Println(stringVar)

	var int1, int2 = 1, 2
	fmt.Println(int1, int2)

	var booleanFlag = true
	fmt.Println(booleanFlag)

	var emptyInt int
	var emptyString string
	var emptyBool bool

	fmt.Println(emptyInt, emptyString, emptyBool)

	initVarViaValueNoVar := "sugar"
	fmt.Println(initVarViaValueNoVar)

}
