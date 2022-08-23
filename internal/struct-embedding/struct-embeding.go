package struct_embedding

import "fmt"

type base struct {
	num int
}

type addon struct {
	feature string
}

func (add addon) printFeature() {
	fmt.Println(add.feature)
}

type container struct {
	base
	addon
	name string
}

func RunSample() {
	fmt.Println("Struct embedding package output:")
	c := container{
		base:  base{1},
		addon: addon{"Custom thing"},
		name:  "Smth",
	}
	fmt.Println(c)
	fmt.Println(c.feature)
	fmt.Println(c.addon.feature)
	c.printFeature()
	c.addon.printFeature()
}
