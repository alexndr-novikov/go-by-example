package exit

import (
	"fmt"
	"os"
)

func RunSample() {
	defer fmt.Println("exit package output:")
	os.Exit(3)
	fmt.Println("")
}
