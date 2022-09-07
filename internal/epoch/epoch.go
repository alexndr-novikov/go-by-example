package epoch

import (
	"fmt"
	"time"
)

func RunSample() {
	fmt.Println("Epoch package output:")
	now := time.Now()
	fmt.Println("Unix: ", now.Unix())
	fmt.Println("UnixMicro: ", now.UnixMicro())
	fmt.Println("UnixMilli: ", now.UnixMilli())
	fmt.Println("UnixNano: ", now.UnixNano())

	fmt.Println("Time from Unix: ", time.Unix(now.Unix(), 0))
	fmt.Println("Time from Unix: ", time.Unix(0, now.UnixNano()))

	fmt.Println("")
}
