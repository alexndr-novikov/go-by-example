package envs

import (
	"fmt"
	"os"
)

func RunSample() {
	fmt.Println("Envs package output:")
	os.Setenv("FOO", "foo")
	fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("BAR"))
	for k, v := range os.Environ() {
		fmt.Println(k, v)
	}
	fmt.Println("")
}
