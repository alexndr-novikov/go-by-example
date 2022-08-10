package runes

import (
	"fmt"
	"unicode/utf8"
)

func Runes() {
	fmt.Println("Runes package output:")
	const greeting = "вітання"
	fmt.Println(greeting, "len:", len(greeting))

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Println("\r\nrunes:", utf8.RuneCountInString(greeting))

	for idx, runeValue := range greeting {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	for i, w := 0, 0; i < len(greeting); i += w {
		runeValue, width := utf8.DecodeRuneInString(greeting[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
	fmt.Println("")
}

func examineRune(r rune) {
	if r == 'i' {
		fmt.Println("Found i")
	}
}
