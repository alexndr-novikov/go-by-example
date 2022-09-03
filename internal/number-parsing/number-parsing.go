package number_parsing

import (
	"fmt"
	"strconv"
)

func RunSample() {
	fmt.Println("Number parsing package output:")
	i, _ := strconv.ParseInt("123", 0, 64)
	f, _ := strconv.ParseFloat("1.123", 64)
	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	ui, _ := strconv.ParseUint("112", 0, 64)
	aoi, _ := strconv.Atoi("112")
	_, err := strconv.Atoi("wat")
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(h)
	fmt.Println(ui)
	fmt.Println(aoi)
	fmt.Println(err)
	fmt.Println("")
}
