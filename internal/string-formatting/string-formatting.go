package string_formatting

import (
	"fmt"
	"os"
)

var p = fmt.Printf

type point struct {
	x, y int
}

func RunSample() {
	fmt.Println("String formatting package output:")
	str := point{1, 2}
	p("struct1: %v\n", str)
	p("struct2: %+v\n", str)
	p("struct3: %#v\n", str)
	p("type: %T\n", str)
	p("boolean: %t\n", true)
	p("int: %d\n", 123)
	p("bin: %b\n", 123)
	p("char: %c\n", 64)
	for _, s := range "hello" {
		fmt.Printf("%d %c \n", s, s)
	}
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("hex: %x\n", 1)
	fmt.Printf("hex: %x\n", 12)
	p("Float: %f\n", 11.01)
	p("Float: %e\n", 11.01)
	p("Float: %E\n", 11.01)
	p("Pointer: %p\n", &str)
	p("String: %s\n", "\"hello\"")
	p("String: %s\n", "hello")
	p("String: %q\n", "\"hello\"")
	p("String (hex): %x\n", "hello")
	p("|%6d|%6d|\n", 1, 11)
	p("|%-6d|%-6d|\n", 1, 11)
	p("|%-6.2f|%-6.2f|\n", 1.001, 11.1)

	sp := fmt.Sprintf("|%-50s|", "test")
	fmt.Println(sp)
	fmt.Fprintf(os.Stderr, "|%-10s|", "error")
	fmt.Println("")
}
