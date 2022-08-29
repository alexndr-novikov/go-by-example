package panicpack

import "os"

func RunSample() {
	panic("A huge problem")

	// this is unreachable
	_, err := os.Create("/tmp/fl12")
	if err != nil {
		panic(err)
	}
}
