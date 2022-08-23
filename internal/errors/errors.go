package errors

import (
	"errors"
	"fmt"
)

func errorTester(i int) (int, error) {
	if i == 42 {
		return 0, errors.New("42 is forbidden")
	}
	return i + 3, nil
}

type customErrStruct struct {
	arg  int
	prob string
}

func (e *customErrStruct) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}

func customErrorTester(i int) (int, error) {
	if i == 42 {
		return 0, &customErrStruct{-1, "42 is forbidden"}
	}
	return i + 3, nil
}

func RunSample() {
	fmt.Println("Errors package output:")
	fmt.Println(errorTester(1))
	fmt.Println(errorTester(42))
	fmt.Println(customErrorTester(1))
	fmt.Println(customErrorTester(42))

	_, e := customErrorTester(42)
	if ae, ok := e.(*customErrStruct); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	fmt.Println("")
}
