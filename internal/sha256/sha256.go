package sha256

import (
	sh "crypto/sha256"
	"fmt"
)

func RunSample() {
	fmt.Println("sha256 package output:")
	s := "Lorem ipsum"
	h := sh.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Println("")
}
